package arcimg

import (
	"encoding/json"
	"io/ioutil"
)

type m []struct {
	ID   string `json:"id"`
	Song string `json:"song"`
}

var songmap map[string]string = make(map[string]string)

func init() {
	f, err := ioutil.ReadFile("./songname.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &songmap)
	if err != nil {
		panic(err)
	}
}

func getsongname(id string) string {
	name, ok := songmap[id]
	if !ok {
		return id
	}
	return name
}
