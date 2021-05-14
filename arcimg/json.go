package arcimg

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type arcinfo struct {
	Content arcinfoContent `json:"content"`
	Status  int            `json:"status"`
}

type arcinfoContent struct {
	Character              int                         `json:"character"`
	Code                   string                      `json:"code"`
	IsCharUncapped         bool                        `json:"is_char_uncapped"`
	IsCharUncappedOverride bool                        `json:"is_char_uncapped_override"`
	IsMutual               bool                        `json:"is_mutual"`
	IsSkillSealed          bool                        `json:"is_skill_sealed"`
	JoinDate               int                         `json:"join_date"`
	Name                   string                      `json:"name"`
	Rating                 int                         `json:"rating"`
	RecentScore            []arcinfoContentRecentScore `json:"recent_score"`
	UserID                 int                         `json:"user_id"`
}

type arcinfoContentRecentScore struct {
	BestClearType     int     `json:"best_clear_type"`
	ClearType         int     `json:"clear_type"`
	Difficulty        int     `json:"difficulty"`
	Health            int     `json:"health"`
	MissCount         int     `json:"miss_count"`
	Modifier          int     `json:"modifier"`
	NearCount         int     `json:"near_count"`
	PerfectCount      int     `json:"perfect_count"`
	Rating            float64 `json:"rating"`
	Score             int     `json:"score"`
	ShinyPerfectCount int     `json:"shiny_perfect_count"`
	SongID            string  `json:"song_id"`
	TimePlayed        int     `json:"time_played"`
	UID               int     `json:"uid"`
}

func json2(jsonn []byte) (arcinfo, error) {
	var arc arcinfo
	if err := json.Unmarshal(jsonn, &arc); err != nil {
		log.Println(err)
		return arc, err
	}
	if arc.Status != 0 {
		return arc, &ErrApiStatus{Code: arc.Status}
	}
	return arc, nil
}

type ErrApiStatus struct {
	Code int
}

func (e *ErrApiStatus) Error() string {
	return "ErrApiStatus code: " + strconv.Itoa(e.Code)
}

func (a *arcinfo) atype() string {
	switch a.Content.RecentScore[0].ClearType {
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
	return convertTimeToFormat(a.Content.RecentScore[0].TimePlayed)
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
	return "PURE: " + strconv.Itoa(a.Content.RecentScore[0].PerfectCount) + "(" +
		strconv.Itoa(a.Content.RecentScore[0].ShinyPerfectCount) + ")"
}

func (a *arcinfo) Far() string {
	return "FAR: " + strconv.Itoa(a.Content.RecentScore[0].NearCount)
}

func (a *arcinfo) Lost() string {
	return "LOST: " + strconv.Itoa(a.Content.RecentScore[0].MissCount)
}

func (a *arcinfo) Rating() string {
	str1 := fmt.Sprintf("%f", a.Content.RecentScore[0].Rating)
	return "Result rating: " + str1
}

func (a *arcinfo) PTT() string {
	ptt := strconv.Itoa(a.Content.Rating)
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
	Difficulty := a.Content.RecentScore[0].Difficulty
	id := a.Content.RecentScore[0].SongID
	switch Difficulty {
	case 0:
		return "PST" + getdifficutie(id, Difficulty)
	case 1:
		return "PRS" + getdifficutie(id, Difficulty)
	case 2:
		return "FTR" + getdifficutie(id, Difficulty)
	case 3:
		return "BYD" + getdifficutie(id, Difficulty)
	default:
		return ""
	}
}
