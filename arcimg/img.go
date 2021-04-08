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
}

func createimg(w io.Writer, info *arcinfo) error {
	songname := getsongname(info.Value[0].Avalue.Friends[0].Recentscore[0].SongID)
	a := arc{
		Sone:     songname + "(" + info.SongID() + ")",
		Score:    strconv.Itoa(info.Value[0].Avalue.Friends[0].Recentscore[0].Score),
		Status:   info.atype(),
		Time:     info.Time(),
		P:        info.Pure(),
		F:        info.Far(),
		L:        info.Lost(),
		Rating:   info.Rating(),
		PTT:      info.PTT(),
		Name:     info.Value[0].Avalue.Friends[0].Name,
		JOinTime: "",
	}
	err := t.ExecuteTemplate(w, "arc", a)
	if err != nil {
		return fmt.Errorf("createimg: %w", err)
	}
	return nil
}
