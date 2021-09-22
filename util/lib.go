package util

import (
	"github.com/dc-replay/go-client/constants"
	"os"
	"strings"
)

var apikey string

func init() {
	apikey = os.Getenv("REPLAY_APIKEY")
}

func Addr() string {
	ret := os.Getenv("REPLAY_ADDR")
	if ret == "" {
		ret = "http://localhost:8080"
	} else if strings.HasPrefix(ret, ":") {
		ret = "http://localhost" + ret
	}

	return ret
}

func DataDir() string {
	ret := os.Getenv(constants.ENV_REPLAY_DATADIR)
	return ret
}

func Repo() string {
	repo := os.Getenv("REPLAY_REPO")
	if repo == "" {
		path, _ := os.Getwd()
		path = strings.Replace(path, "\\", "/", -1)
		pathparts := strings.Split(path, "/")
		repo = pathparts[len(pathparts)-1]

	}
	return repo
}

//
//func HttpDo(method string, strurl string, body []byte) (*http.Response, error) {
//	cli := http.Client{}
//	var err error
//	var req *http.Request
//
//	if body != nil && len(body) > 0 {
//		req, err = http.NewRequest(method, strurl, io.NopCloser(bytes.NewReader(body)))
//
//	} else {
//		req, err = http.NewRequest(method, strurl, nil)
//	}
//	if err != nil {
//		return nil, err
//	}
//	req.Header.Set("X-API-KEY", apikey)
//	req.Header.Set("Content-Type", "application/json")
//	return cli.Do(req)
//}
//
//func HttpDoJson(i interface{}, method string, strurl string, body []byte) (err error) {
//	res, err := HttpDo(method, strurl, body)
//	if err != nil {
//		return
//	}
//	err = json.NewDecoder(res.Body).Decode(i)
//	return
//}
