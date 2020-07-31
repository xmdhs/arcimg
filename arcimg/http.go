package arcimg

import (
	"bytes"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var (
	atime int64
	btime int64
	mu    sync.Mutex
	ajson atomic.Value
	ma    sync.Map
	at    atomic.Value
)

func init() {
	get()
	if ajson.Load().(string) == "" {
		log.Fatalln("Can not get json")
	}
}

func Img(w http.ResponseWriter, req *http.Request) {
	if time.Now().Unix()-atomic.LoadInt64(&atime) > 600 {
		atomic.StoreInt64(&atime, time.Now().Unix())
		go get()
	}
	if time.Now().Unix()-atomic.LoadInt64(&btime) > 30 {
		atomic.StoreInt64(&btime, time.Now().Unix())
		info, err := Json2(ajson.Load().(string))
		if err != nil {
			log.Println(err)
		} else if info.Value != nil {
			bb := buffer.Get()
			c := bb.(*bytes.Buffer)
			c.Reset()
			createimg(c, &info)
			at.Store(c)
		}
	}
	w.Header().Set("Cache-Control", "max-age=60")
	w.Header().Set("server", "xmdhs")
	w.Write(at.Load().(*bytes.Buffer).Bytes())
}
