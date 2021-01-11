package arcimg

import "time"

func remove() {
	for {
		time.Sleep(6000 * time.Millisecond)
		ma.Range(func(k, v interface{}) bool {
			vv := v.(int)
			if vv <= 0 {
				ma.Delete(k)
			} else {
				ma.Store(k, vv-1)
			}
			return true
		})
	}
}
