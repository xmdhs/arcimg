package arcimg

import (
	"io"
	"log"
	"os"
)

var (
	logger *log.Logger
)

func logw() io.ReadWriteCloser {
	f, err := os.OpenFile("arcimg.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	fileAndStdoutWriter := io.MultiWriter(f, os.Stdout)
	logger = log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime)
	return f
}
