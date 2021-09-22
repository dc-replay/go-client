package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dc-replay/go-client/types"
	"io"
	"net/http"
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

func HttpDo(method string, strurl string, body []byte) (*http.Response, error) {
	cli := http.Client{}
	var err error
	var req *http.Request

	if body != nil && len(body) > 0 {
		req, err = http.NewRequest(method, strurl, io.NopCloser(bytes.NewReader(body)))

	} else {
		req, err = http.NewRequest(method, strurl, nil)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-KEY", apikey)
	req.Header.Set("Content-Type", "application/json")
	return cli.Do(req)
}

func HttpDoJson(i interface{}, method string, strurl string, body []byte) (err error) {
	res, err := HttpDo(method, strurl, body)
	if err != nil {
		return
	}
	err = json.NewDecoder(res.Body).Decode(i)
	return
}

type roboDAO struct {
}

func (d roboDAO) GetAll() ([]*types.Robot, error) {
	ret := make([]*types.Robot, 0)
	err := HttpDoJson(&ret, http.MethodGet, Addr()+fmt.Sprintf("/api/v1/robots"), nil)
	return ret, err
}
func (d roboDAO) Enqueue(id uint) error {
	_, err := HttpDo(http.MethodGet, Addr()+fmt.Sprintf("/api/v1/robots/op/enqueue/%d", id), nil)
	return err
}

var RoboDAO roboDAO

type sqlDAO struct {
}

func (d sqlDAO) SQL(s string) (types.SQLResponse, error) {
	ret := types.SQLResponse{}
	req := types.SQLRequest{Sql: s}
	bs, _ := json.Marshal(req)
	err := HttpDoJson(&ret, http.MethodPost, Addr()+fmt.Sprintf("/api/v1/sql"), bs)
	return ret, err
}

var SQLDAO sqlDAO

type appDAO struct {
}

func (d appDAO) GetAll() ([]*types.App, error) {
	ret := make([]*types.App, 0)
	err := HttpDoJson(&ret, http.MethodGet, Addr()+fmt.Sprintf("/api/v1/apps"), nil)
	return ret, err
}
func (d appDAO) Run(id interface{}) ([]*types.App, error) {
	ret := make([]*types.App, 0)
	err := HttpDoJson(&ret, http.MethodGet, Addr()+fmt.Sprintf("/api/v1/app/run/%v", id), nil)
	return ret, err
}

var AppDAO appDAO
