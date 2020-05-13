package arcimg

import (
	"bytes"
	"log"
	"net/http"
	"sync"
	"time"
)

var atime int64
var mu sync.Mutex
var ajson *string
var b bytes.Buffer
var ma sync.Map

func Img(w http.ResponseWriter, req *http.Request) {
	ip := req.Header.Get("X-Forwarded-For")
	log.Println(ip)
	i, bb := ma.LoadOrStore(ip, 0)
	if bb {
		ii, _ := i.(int)
		if ii > 20 {
			ma.Store(ip, 30)
		} else {
			ma.Store(ip, ii+1)
		}
	}
	i, bb = ma.Load(ip)
	ii, _ := i.(int)
	if ii > 5 {
		return
	}
	mu.Lock()
	if time.Now().Unix()-atime > 600000 {
		atime = time.Now().Unix()
		ajson = Getjson()
	}
	mu.Unlock()
	w.Header().Add("Cache-Control", "max-age=60")
	w.Header().Add("content-type", "image/png")
	w.Header().Add("server", "xmdhs")
	info := Json2(ajson)
	mu.Lock()
	if b.Len() == 0 {
		Createimg(&b, &info)

	}
	if time.Now().Unix()-atime > 30000 {
		atime = time.Now().Unix()
		Createimg(&b, &info)
	}
	mu.Unlock()
	w.Write(b.Bytes())
}
