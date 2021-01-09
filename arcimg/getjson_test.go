package arcimg

import (
	"fmt"
	"testing"
)

func TestGetjson(t *testing.T) {
	a := getjson(5)
	if len(a) <= 0 {
		t.Error("err")
	}
}

func TestGet(t *testing.T) {
	a := ajson.Load().([]byte)
	fmt.Println(string(a))
}
