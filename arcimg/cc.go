package arcimg

import (
	"bytes"
	"log"
	"sync"
)

var (
	oo sync.Once
	o  sync.Once
)

func Anticc(m *Http) {
	oo.Do(func() {
		go Remove()
	})
	ip := m.req.Header.Get("X-Forwarded-For")
	log.Println(ip + " | " + m.req.Header.Get("Referer"))
	i, bb := ma.LoadOrStore(ip, 0)
	if bb {
		ii, _ := i.(int)
		ma.Store(ip, ii+1)
	}
	i, bb = ma.Load(ip)
	ii, _ := i.(int)
	if ii > 5 {
		ma.Store(ip, 30)
		return
	}
	m.Next()
}

func Log(m *Http) {
	o.Do(func() {
		go Logw()
	})
	if m.req.URL.String() != "/img.png" {
		ip := m.req.Header.Get("X-Forwarded-For")
		b := buffer.Get()
		bb := b.(*bytes.Buffer)
		bb.Reset()
		for key, v := range m.req.Header {
			bb.WriteString(key + ": ")
			for _, v := range v {
				bb.WriteString(v)
			}
			bb.WriteString(" ")
		}
		loggers <- ip + " | " + bb.String() + " | " + m.req.URL.String()
		buffer.Put(bb)
	}
	m.Next()
}

var buffer sync.Pool

func init() {
	buffer = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(nil)
		},
	}
}
