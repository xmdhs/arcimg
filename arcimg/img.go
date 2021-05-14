package arcimg

import (
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"text/template"
)

//go:embed arc.svg
var b string

var (
	t *template.Template
)

func init() {
	var err error
	t, err = template.New("arc").Parse(b)
	if err != nil {
		panic(err)
	}
}

type arc struct {
	Sone     string
	Score    string
	Status   string
	Time     string
	P        string
	F        string
	L        string
	Rating   string
	PTT      string
	Name     string
	JOinTime string
	Size     string
	X        string
}

func createimg(w io.Writer, info *arcinfo) error {
	songname := getsongname(info.Content.RecentScore[0].SongID)
	size := "19"
	X := "68"
	switch {
	case len(songname) > 35:
		size = "8"
		X = "15"
	case len(songname) > 25:
		size = "13"
		X = "25"
	case len(songname) > 20:
		size = "15"
		X = "45"
	case len(songname) > 15:
		size = "16"
	}
	a := arc{
		Sone:     songname + "(" + info.SongID() + ")",
		Score:    strconv.Itoa(info.Content.RecentScore[0].Score),
		Status:   info.atype(),
		Time:     info.Time(),
		P:        info.Pure(),
		F:        info.Far(),
		L:        info.Lost(),
		Rating:   info.Rating(),
		PTT:      info.PTT(),
		Name:     info.Content.Name,
		JOinTime: "",
		Size:     size,
		X:        X,
	}
	err := t.ExecuteTemplate(w, "arc", a)
	if err != nil {
		return fmt.Errorf("createimg: %w", err)
	}
	return nil
}
