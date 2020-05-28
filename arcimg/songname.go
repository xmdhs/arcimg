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
	songlist := m{}
	err = json.Unmarshal(f, &songlist)
	if err != nil {
		panic(err)
	}
	for i := range songlist {
		songmap[songlist[i].ID] = songlist[i].Song
	}
}

func getsongname(id string) string {
	name, ok := songmap[id]
	if !ok {
		return id
	}
	return name
}
