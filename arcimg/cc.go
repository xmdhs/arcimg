package arcimg

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

var (
	oo sync.Once
	o  sync.Once
)

func Anticc(f http.HandlerFunc) http.HandlerFunc {
	oo.Do(func() {
		go remove()
	})
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Real-Ip")
		log.Println(ip + " | " + r.Header.Get("Referer"))
		i := ma.Get(ip)
		if i > 20 {
			ma.Store(ip, 30)
			return
		}
		ma.Store(ip, i+1)
		f(w, r)
	}
}

func Log(f http.HandlerFunc) http.HandlerFunc {
	o.Do(func() {
		logger = logw()
	})
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.Referer())
		var host string
		if err == nil {
			host = u.Hostname()
		}
		if r.URL.String() != "/img.png" || !strings.HasSuffix(host, "mcbbs.net") {
			ip := r.Header.Get("X-Real-Ip")
			b := buffer.Get()
			bb := b.(*bytes.Buffer)
			bb.Reset()
			for key, v := range r.Header {
				bb.WriteString(key + ": ")
				for _, v := range v {
					bb.WriteString(v)
				}
				bb.WriteString(" ")
			}
			logger.Println(ip + " | " + bb.String() + " | " + r.URL.String())
			buffer.Put(bb)
		}
		f(w, r)
	}
}

var buffer sync.Pool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(nil)
	},
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, v := range middlewares {
		f = v(f)
	}
	return f
}
