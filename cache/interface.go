package cache

import "time"

type Cache interface {
	Get(string) ([]byte, error)
	Set(string, []byte, time.Time) error
	Del(string) error
}
