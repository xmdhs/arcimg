// +build conver

package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

func main() {
	b, err := ioutil.ReadFile("songlist")
	if err != nil {
		panic(err)
	}
	var data = make(map[string]songdata)
	var a arc
	err = json.Unmarshal(b, &a)
	if err != nil {
		panic(err)
	}
	for _, v := range a.Songs {
		d := []string{}
		for _, v := range v.Difficulties {
			rating := strconv.Itoa(v.Rating)
			if v.RatingPlus {
				rating += "+"
			}
			d = append(d, rating)
		}
		s := songdata{
			Name:         v.TitleLocalized.En,
			Difficulties: d,
		}
		data[v.ID] = s
	}
	b, err = json.MarshalIndent(data, "", "    ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("songname.json", b, 0777)
	if err != nil {
		panic(err)
	}
}

type arc struct {
	Songs []arcSong `json:"songs"`
}

type arcSong struct {
	Artist          string                `json:"artist"`
	AudioPreview    float64               `json:"audioPreview"`
	AudioPreviewEnd float64               `json:"audioPreviewEnd"`
	Bg              string                `json:"bg"`
	Bpm             string                `json:"bpm"`
	BpmBase         float64               `json:"bpm_base"`
	Date            float64               `json:"date"`
	Difficulties    []arcSongDifficulty   `json:"difficulties"`
	ID              string                `json:"id"`
	Purchase        string                `json:"purchase"`
	Set             string                `json:"set"`
	Side            float64               `json:"side"`
	TitleLocalized  arcSongTitleLocalized `json:"title_localized"`
	Version         string                `json:"version"`
}

type arcSongDifficulty struct {
	ChartDesigner  string `json:"chartDesigner"`
	JacketDesigner string `json:"jacketDesigner"`
	Rating         int    `json:"rating"`
	RatingClass    int    `json:"ratingClass"`
	RatingPlus     bool   `json:"ratingPlus"`
}

type arcSongTitleLocalized struct {
	En string `json:"en"`
}

type songdata struct {
	Name         string
	Difficulties []string
}
