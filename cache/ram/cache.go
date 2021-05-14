package ram

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/xmdhs/arcimg/cache"
)

type Cache struct {
	m      sync.Map
	cancel func()
}

type date struct {
	Time time.Time
	Date []byte
}

var _ cache.Cache = &Cache{}

func Newcache() *Cache {
	c := &Cache{}
	cxt := context.Background()
	cxt, cancel := context.WithCancel(cxt)
	c.cancel = cancel
	c.delete(cxt)
	return c
}

func (c *Cache) delete(cxt context.Context) {
	go func() {
		t := time.NewTicker(10 * time.Minute)
		defer t.Stop()
		for {
			at := time.Now()
			c.m.Range(func(key, value interface{}) bool {
				d := value.(date)
				if d.Time.Before(at) {
					c.m.Delete(key)
				}
				return true
			})

			select {
			case <-cxt.Done():
				return
			case <-t.C:
			}
		}
	}()
}

func (c *Cache) Close() {
	c.cancel()
}

var ErrNotFind = errors.New("ErrNotFind")

func (c *Cache) Del(key string) error {
	c.m.Delete(key)
	return nil
}

func (c *Cache) Get(key string) ([]byte, error) {
	t, ok := c.m.Load(key)
	if !ok {
		return nil, ErrNotFind
	}
	d, ok := t.(date)
	if !ok {
		return nil, ErrNotFind
	}
	return d.Date, nil
}

func (c *Cache) Set(key string, adate []byte, t time.Time) error {
	d := date{
		Date: adate,
		Time: t,
	}
	c.m.Store(key, d)
	return nil
}
