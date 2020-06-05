package arcimg

import (
	"io"
	"log"
	"os"
)

var (
	loggers = make(chan string, 1000)
	logger  *log.Logger
)

func Logw() {
	f, err := os.OpenFile("arcimg.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fileAndStdoutWriter := io.MultiWriter(f, os.Stdout)
	logger = log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime)
	var t string
	for {
		t = <-loggers
		logger.Printf(t)
	}
}
