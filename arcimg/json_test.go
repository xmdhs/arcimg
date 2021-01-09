package arcimg

import (
	"fmt"
	"testing"
)

func TestJson2(t *testing.T) {
	a, err := json2(ajson.Load().([]byte))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}
