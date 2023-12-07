package wget

import (
	"errors"
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getFilename(request *http.Request, resp *http.Response) (string, error) {
	filename := filepath.Base(request.URL.Path)

	//original link didnt represent the file type. Try using the response url (after redirects)
	if !strings.Contains(filename, ".") {
		filename = filepath.Base(resp.Request.URL.Path)
	}

	if !strings.Contains(filename, ".") {
		ct := resp.Header.Get("Content-Type")
		ext := "htm"
		mediaType, _, err := mime.ParseMediaType(ct)
		if err != nil {
			return "", err
		}

		slash := strings.Index(mediaType, "/")
		if slash != -1 {
			_, sub := mediaType[:slash], mediaType[slash+1:]
			if sub != "" {
				ext = sub
			}
		}

		filename = filename + "." + ext
	}

	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return filename, nil
		}

		return "", err
	}

	num := 1
	//just stop after 100
	for num < 100 {
		filenameNew := fmt.Sprintf("%s.%s", filename, strconv.Itoa(num))
		_, err = os.Stat(filenameNew)
		if err != nil {
			if os.IsNotExist(err) {
				return filenameNew, nil
			}

			return "", err
		}
		num += 1
	}

	return filename, errors.New("stopping after trying 100 filename variants")
}
