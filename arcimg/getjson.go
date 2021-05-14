package arcimg

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func getJson(uid string) ([]byte, error) {
	u := url.URL{}
	u.Host = apiaddress.Host
	u.Scheme = apiaddress.Scheme
	u.Path = "/v4/user/info"
	v := url.Values{}
	v.Set("usercode", uid)
	v.Set("recent", "1")
	u.RawQuery = v.Encode()
	b, err := httpGet(u.String())
	if err != nil {
		return nil, fmt.Errorf("getJson: %w", err)
	}
	return b, nil
}

var client = http.Client{
	Timeout: 5 * time.Second,
}

func httpGet(url string) ([]byte, error) {
	reps, err := client.Get(url)
	if reps != nil {
		defer reps.Body.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("httpGet: %w", err)
	}
	b, err := io.ReadAll(reps.Body)
	if err != nil {
		return nil, fmt.Errorf("httpGet: %w", err)
	}
	return b, nil
}

var apiaddress *url.URL

func init() {
	a := os.Getenv("apiaddress")
	if a == "" {
		panic(`apiaddress == ""`)
	}
	var err error
	apiaddress, err = url.Parse(a)
	if err != nil {
		panic(err)
	}
}
