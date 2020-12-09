// +build json

package main

import (
	"encoding/json"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile(`songname.json`)
	if err != nil {
		panic(err)
	}
	m := make(map[string]string)
	err = json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}
	b, err = json.MarshalIndent(m, "", "    ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(`songname.json`, b, 0777)
	if err != nil {
		panic(err)
	}
}
