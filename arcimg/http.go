package arcimg

import (
	"bytes"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	atime int64
	btime int64
	mu    sync.Mutex
	ajson string
	b     *bytes.Buffer = bytes.NewBuffer(nil)
	ma    sync.Map
	o     sync.Once
)

func init() {
	get()
	if ajson == "" {
		log.Fatalln("Can not get json")
	}
}

func Img(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	if time.Now().Unix()-atime > 600 {
		atime = time.Now().Unix()
		go get()
	}
	if time.Now().Unix()-btime > 30 {
		btime = time.Now().Unix()
		info, err := Json2(ajson)
		if err != nil {
			log.Println(err)
		} else if info.Value != nil {
			bb := buffer.Get()
			c := bb.(*bytes.Buffer)
			c.Reset()
			createimg(c, &info)
			b = c
		}
	}
	mu.Unlock()
	w.Header().Set("Cache-Control", "max-age=60")
	w.Header().Set("server", "xmdhs")
	w.Write(b.Bytes())
}
