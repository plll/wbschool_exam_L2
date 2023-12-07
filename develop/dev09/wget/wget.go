package wget

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type signal struct {
	id       int
	progress float32
}

type WGet struct {
	taskCount  int
	uploadList []string

	wg     *sync.WaitGroup
	ticker *time.Ticker
}

func NewWGet(uploadList []string) *WGet {
	return &WGet{
		uploadList: uploadList,
		taskCount:  len(uploadList),
		wg:         new(sync.WaitGroup),
		ticker:     time.NewTicker(time.Second),
	}
}

func (u *WGet) Exec(ctx context.Context) error {
	receiveChan := make(chan signal, len(u.uploadList))
	uploadTask := func(id int, url string, taskCount *int, wg *sync.WaitGroup) {
		defer func() {
			wg.Done()
			*taskCount--
		}()

		if err := uploadFile(ctx, url, id, receiveChan); err != nil {
			log.Fatalln(err.Error())
		}
	}

	u.wg.Add(len(u.uploadList))
	for i, _ := range u.uploadList {
		url := u.uploadList[i]

		go uploadTask(i, url, &u.taskCount, u.wg)
	}

	monitoringTask := func(wg *sync.WaitGroup) {
		defer u.ticker.Stop()
		defer wg.Done()

		progressState := make([]float32, len(u.uploadList))
		for {
			select {
			case r := <-receiveChan:
				progressState[r.id] = r.progress
			case <-u.ticker.C:
				if u.taskCount == 0 {
					return
				}

				resultLog := ""
				for i, _ := range progressState {
					resultLog += fmt.Sprintf("%.2f%% ", progressState[i])
				}

				fmt.Printf("%s\n", resultLog) // for console
			}
		}
	}

	u.wg.Add(1)
	go monitoringTask(u.wg)

	u.wg.Wait()
	close(receiveChan)

	return nil
}

func uploadFile(ctx context.Context, url string, id int, receiveChan chan<- signal) error {
	if !strings.Contains(url, ":") {
		url = fmt.Sprintf("http://%s", url)
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	request = request.WithContext(ctx)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("invalid response: status: %s", response.Status))
	}

	fileName, err := getFilename(request, response)
	if err != nil {
		return errors.New(fmt.Sprintf("bad file name"))
	}

	openFlags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	outFile, err := os.OpenFile(fileName, openFlags, os.FileMode(0666))
	if err != nil {
		return err
	}

	defer func() {
		if errClose := outFile.Close(); errClose != nil {
			log.Fatalln(errClose.Error())
		}
	}()

	read := int64(0)
	p := float32(0)
	length := response.ContentLength

	reader := bufio.NewReader(response.Body)
	writer := bufio.NewWriter(outFile)
	for {

		buffer := make([]byte, 4096)
		cBytes, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		read = read + int64(cBytes)

		if read <= 0 {
			break
		}

		p = float32(read*100) / float32(length)

		receiveChan <- signal{id: id, progress: p}

		if _, err = writer.Write(buffer[0:cBytes]); err != nil {
			break
		}

	}

	return nil
}
