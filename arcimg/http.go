package arcimg

import (
	"bytes"
	"log"
	"net/http"
	"sync"
	"time"
)

var atime int64
var btime int64
var mu sync.Mutex
var ajson string
var b *bytes.Buffer
var ma sync.Map

func init() {
	get()
	if ajson == "" {
		log.Fatalln("Can not get json")
	}
}

func Img(w http.ResponseWriter, req *http.Request) {
	ip := req.Header.Get("X-Forwarded-For")
	log.Println(ip + " | " + req.Header.Get("Referer"))
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
	if time.Now().Unix()-atime > 600 {
		atime = time.Now().Unix()
		go get()
	}
	if time.Now().Unix()-btime > 30 {
		btime = time.Now().Unix()
		info, err := Json2(ajson)
		if err != nil {
			return
		}
		abyte := []byte{}
		c := bytes.NewBuffer(abyte)
		createimg(c, &info)
		b = c
	}
	mu.Unlock()
	w.Header().Set("Cache-Control", "max-age=60")
	w.Header().Set("server", "xmdhs")
	w.Write(b.Bytes())
}
