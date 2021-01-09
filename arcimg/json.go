package arcimg

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type arcinfo struct {
	Value []value `json:"value"`
}

type value struct {
	Avalue avalue `json:"value"`
}

type avalue struct {
	Friends []friends `json:"friends"`
}

type friends struct {
	Name        string        `json:"name"`
	Rating      int           `json:"rating"`
	Recentscore []recentscore `json:"recent_score"`
}

type recentscore struct {
	BestClearType     int     `json:"best_clear_type"`
	ClearType         int     `json:"clear_type"`
	Difficulty        int     `json:"difficulty"`
	MissCount         int     `json:"miss_count"`
	Modifier          int     `json:"modifier"`
	NearCount         int     `json:"near_count"`
	PerfectCount      int     `json:"perfect_count"`
	Rating            float64 `json:"rating"`
	Score             int     `json:"score"`
	ShinyPerfectCount int     `json:"shiny_perfect_count"`
	SongID            string  `json:"song_id"`
	TimePlayed        int     `json:"time_played"`
}

func Json2(jsonn []byte) (arcinfo, error) {
	var arc arcinfo
	if err := json.Unmarshal(jsonn, &arc); err != nil {
		log.Println(err)
		return arc, err
	}
	return arc, nil
}

func (a *arcinfo) atype() string {
	switch a.Value[0].Avalue.Friends[0].Recentscore[0].ClearType {
	case 0:
		return "Track Lost"
	case 1:
		return "Normal Clear"
	case 2:
		return "Full Recall"
	case 3:
		return "Pure Memory"
	case 4:
		return "Easy Clear"
	case 5:
		return "Hard Clear"
	default:
		return ""
	}
}

func (a *arcinfo) Time() string {
	return convertTimeToFormat(a.Value[0].Avalue.Friends[0].Recentscore[0].TimePlayed)
}

func convertTimeToFormat(timetamp int) string {
	curTime := time.Now().Unix()
	timetamp = timetamp / 1000
	time := int(curTime) - timetamp
	if time < 60 && time >= 0 {
		return "now"
	} else if time >= 60 && time < 3600 {
		if time/60 == 1 {
			return strconv.Itoa(time/60) + " minute ago"
		}
		return strconv.Itoa(time/60) + " minutes ago"

	} else if time >= 3600 && time < 3600*24 {
		if time/3600 == 1 {
			return strconv.Itoa(time/3600) + " hour ago"
		}
		return strconv.Itoa(time/3600) + " hours ago"

	} else if time >= 3600*24 && time < 3600*24*30 {
		if time/3600/24 == 1 {
			return strconv.Itoa(time/3600/24) + " day ago"
		}
		return strconv.Itoa(time/3600/24) + " days ago"

	} else if time >= 3600*24*30 && time < 3600*24*30*12 {
		if time/3600/24/30 == 1 {
			return strconv.Itoa(time/3600/24/30) + " month ago"
		}
		return strconv.Itoa(time/3600/24/30) + " months ago"

	} else if time >= 3600*24*30*12 {
		if time/3600/24/30/12 == 1 {
			return strconv.Itoa(time/3600/24/30/12) + " year ago"
		}
		return strconv.Itoa(time/3600/24/30/12) + " years ago"

	}
	return "now"
}

func (a *arcinfo) Pure() string {
	return "PURE: " + strconv.Itoa(a.Value[0].Avalue.Friends[0].Recentscore[0].PerfectCount) + "(" +
		strconv.Itoa(a.Value[0].Avalue.Friends[0].Recentscore[0].ShinyPerfectCount) + ")"
}

func (a *arcinfo) Far() string {
	return "FAR: " + strconv.Itoa(a.Value[0].Avalue.Friends[0].Recentscore[0].NearCount)
}

func (a *arcinfo) Lost() string {
	return "LOST: " + strconv.Itoa(a.Value[0].Avalue.Friends[0].Recentscore[0].MissCount)
}

func (a *arcinfo) Rating() string {
	str1 := fmt.Sprintf("%f", a.Value[0].Avalue.Friends[0].Recentscore[0].Rating)
	return "Result rating: " + str1
}

func (a *arcinfo) PTT() string {
	ptt := strconv.Itoa(a.Value[0].Avalue.Friends[0].Rating)
	i := len(ptt)
	if i == 3 {
		return "PTT: " + ptt[0:1] + "." + ptt[1:]
	}
	if i == 4 {
		return "PTT: " + ptt[0:2] + "." + ptt[2:]
	}
	return "PTT: " + ptt
}

func (a *arcinfo) SongID() string {
	Difficulty := a.Value[0].Avalue.Friends[0].Recentscore[0].Difficulty
	switch Difficulty {
	case 0:
		return "PST"
	case 1:
		return "PRS"
	case 2:
		return "FTR"
	case 3:
		return "BYD"
	default:
		return ""
	}
}
