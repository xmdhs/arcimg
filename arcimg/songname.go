package arcimg

import (
	"encoding/json"
	"io/ioutil"
)

type m []struct {
	ID   string `json:"id"`
	Song string `json:"song"`
}

var songmap = make(map[string]songdata)

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
	return name.Name
}

func getdifficutie(id string, i int) string {
	r, ok := songmap[id]
	if !ok {
		return ""
	}
	return r.Difficulties[i]

}

type songdata struct {
	Name         string
	Difficulties []string
}
