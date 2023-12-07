package main

import (
	"context"
	"fmt"
	"github.com/plll/wbschool_exam_L2/develop/dev09/wget"
	"log"
	"os"
	"time"
)

const UploadTimeout = time.Minute * 10

func main() {
	fmt.Println(os.Args[1:])
	uploader := wget.NewWGet(os.Args[1:])

	ctx, cancel := context.WithTimeout(context.Background(), UploadTimeout)
	defer cancel()

	if err := uploader.Exec(ctx); err != nil {
		log.Fatalln(err.Error())
	}
}
