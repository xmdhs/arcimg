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
	if ajson.Load().([]byte) == nil {
		log.Fatalln("Can not get json")
	}
}

func Img(w http.ResponseWriter, req *http.Request) {
	aoldtime := atomic.LoadInt64(&atime)
	if time.Now().Unix()-aoldtime > 600 && atomic.CompareAndSwapInt64(&atime, aoldtime, time.Now().Unix()) {
		go get()
	}
	boldtime := atomic.LoadInt64(&btime)
	if time.Now().Unix()-boldtime > 30 && atomic.CompareAndSwapInt64(&btime, boldtime, time.Now().Unix()) {
		info, err := Json2(ajson.Load().([]byte))
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
