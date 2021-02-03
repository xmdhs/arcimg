package arcimg

import (
	"sync"
	"time"
)

var ma cc

func remove() {
	for {
		time.Sleep(3 * time.Minute)
		ma.Range(func(k, v interface{}) bool {
			vv := v.(c)
			if time.Now().Unix()-vv.time > 600 {
				ma.Delete(k)
			}
			return true
		})
	}
}

type c struct {
	i    int
	time int64
}

type cc struct {
	sync.Map
}

func (cc *cc) Get(key string) int {
	v, ok := cc.Load(key)
	if !ok {
		return 0
	}
	c := v.(c)
	return c.i
}

func (cc *cc) Store(key string, value int) {
	c := c{
		i:    value,
		time: time.Now().Unix(),
	}
	cc.Map.Store(key, c)
}
