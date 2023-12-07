package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

const UploadTimeout = time.Minute * 10

func main() {
	fmt.Println(os.Args[1:])
	uploader := NewWGet(os.Args[1:])

	ctx, cancel := context.WithTimeout(context.Background(), UploadTimeout)
	defer cancel()

	if err := uploader.Exec(ctx); err != nil {
		log.Fatalln(err.Error())
	}
}
