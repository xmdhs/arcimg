package arcimg

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func getjson(i int) []byte {
	if i <= 0 {
		return nil
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	rep, err := http.NewRequest("GET", "https://arcapi.lowiro.com/blockchain/14/compose/aggregate?calls=%5b%7b+%22endpoint%22%3a+%22%2fuser%2fme%22%2c+%22id%22%3a+0+%7d%5d", nil)
	if err != nil {
		panic(err)
	}
	rep.Header.Set("Accept-Encoding", "identity")
	rep.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	rep.Header.Set("Accept-Language", "zh-cn")
	rep.Header.Set("Accept-Encoding", "gzip")
	rep.Header.Set("Accept", "*/*")
	rep.Header.Set("Authorization", authorization)
	rep.Header.Set("Platform", "ios")
	rep.Header.Set("AppVersion", "3.6.0")
	rep.Header.Set("User-Agent", "Arc-mobile/v3.6.0.0 CFNetwork/811.5.4 Darwin/16.7.0")
	rep.Header.Set("Host", "arcapi.lowiro.com")
	rep.Header.Set("Connection", "Keep-Alive")
	reps, err := client.Do(rep)
	if reps != nil {
		defer reps.Body.Close()
	}
	if err != nil {
		log.Println(err)
		time.Sleep(5 * time.Second)
		return getjson(i - 1)
	}
	var reader io.ReadCloser
	switch reps.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(reps.Body)
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			return getjson(i - 1)
		}
		defer reader.Close()
	default:
		reader = reps.Body
	}
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println(err)
		time.Sleep(5 * time.Second)
		return getjson(i - 1)
	}
	if reps.StatusCode != http.StatusOK {
		log.Println(reps.Status)
		time.Sleep(5 * time.Second)
		return getjson(i - 1)
	}
	return b
}

func get() {
	a := getjson(5)
	if a != nil {
		ajson.Store(a)
	}
}

var authorization string

func init() {
	authorization = os.Getenv("authorization")
	if authorization == "" {
		panic(`authorization == ""`)
	}
}
