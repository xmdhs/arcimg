package arcimg

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/xmdhs/arcimg/cache"
	"github.com/xmdhs/arcimg/cache/ram"
	"golang.org/x/sync/singleflight"
)

var caches cache.Cache = ram.Newcache()

var sl = singleflight.Group{}

func Img(w http.ResponseWriter, req *http.Request) {
	p := httprouter.ParamsFromContext(req.Context())
	uid := p.ByName("uid")

	b, err := caches.Get(uid)
	if err != nil {
		temp, err, _ := sl.Do(uid, func() (interface{}, error) {
			var err error
			for i := 0; i < 3; i++ {
				b, err = getJson(uid)
				if err != nil {
					log.Println(err)
					continue
				}
				a, err := json2(b)
				if err != nil {
					e := &ErrApiStatus{}
					if errors.As(err, &e) {
						break
					}
					log.Println(err)
					continue
				}
				by := bytes.Buffer{}
				err = createimg(&by, &a)
				if err != nil {
					log.Println(err)
					continue
				}
				return by.Bytes(), nil
			}
			return nil, err
		})
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		b = temp.([]byte)
		caches.Set(uid, b, time.Now().Add(20*time.Minute))
	}
	w.Header().Set("Cache-Control", "max-age=60")
	w.Header().Set("content-type", "image/svg+xml")
	w.Header().Set("server", "xmdhs")
	w.Write(b)
}
