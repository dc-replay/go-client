package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

type Cli struct {
	cli     *http.Client
	headers map[string]string
}

func (c *Cli) Addr() string {
	ret := os.Getenv("REPLAY_ADDR")
	if ret == "" {
		ret = "https://localhost:8443"
	} else if strings.HasPrefix(ret, ":") {
		ret = "https://localhost" + ret
	}

	return ret
}
func (c *Cli) AddHeader(k string, v string) {
	c.headers[k] = v
}
func (c *Cli) DelHeader(k string) {
	delete(c.headers, k)
}
func (c *Cli) HttpDo(method string, strurl string, body []byte) (*http.Response, error) {
	if !strings.HasPrefix(strurl, "http") {
		addr := c.Addr()
		if strings.HasSuffix(addr, "/") {
			addr = addr[:len(addr)-1]
		}
		if !strings.HasPrefix(strurl, "") {
			strurl = "/" + strurl
		}
		strurl = addr + strurl
	}
	req, err := http.NewRequest(method, strurl, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	if c.headers != nil {
		for k, v := range c.headers {
			req.Header.Set(k, v)
		}
	}
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
	return c.cli.Do(req)
}
func (c *Cli) HttpDoJson(method string, strurl string, i interface{}, o interface{}) (err error) {
	bs, err := json.Marshal(i)
	if err != nil {
		return err
	}
	res, err := c.HttpDo(method, strurl, bs)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if o != nil {
		err = json.NewDecoder(res.Body).Decode(o)
	}
	return
}
func (c *Cli) HttpJsonGet(strurl string, o interface{}) error {
	return c.HttpDoJson(http.MethodGet, strurl, nil, o)
}
func (c *Cli) HttpJsonDelete(strurl string, o interface{}) error {
	return c.HttpDoJson(http.MethodDelete, strurl, nil, o)
}
func (c *Cli) HttpJsonHead(strurl string, o interface{}) error {
	return c.HttpDoJson(http.MethodHead, strurl, nil, o)
}
func (c *Cli) HttpJsonPost(strurl string, i interface{}, o interface{}) error {
	return c.HttpDoJson(http.MethodPost, strurl, i, o)
}
func (c *Cli) HttpJsonPut(strurl string, i interface{}, o interface{}) error {
	return c.HttpDoJson(http.MethodPut, strurl, i, o)
}
func (c *Cli) HttpJsonPatch(strurl string, i interface{}, o interface{}) error {
	return c.HttpDoJson(http.MethodPatch, strurl, i, o)
}
func (c *Cli) HttpRawGet(strurl string) ([]byte, error) {
	res, err := c.HttpDo(http.MethodGet, strurl, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bs, _ := io.ReadAll(res.Body)
	return bs, nil
}
func (c *Cli) HttpRawDelete(strurl string) ([]byte, error) {
	res, err := c.HttpDo(http.MethodDelete, strurl, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bs, _ := io.ReadAll(res.Body)
	return bs, nil
}
func (c *Cli) HttpRawPost(strurl string, i []byte) ([]byte, error) {
	res, err := c.HttpDo(http.MethodPost, strurl, i)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bs, _ := io.ReadAll(res.Body)
	return bs, nil
}
func (c *Cli) HttpRawPut(strurl string, i []byte) ([]byte, error) {
	res, err := c.HttpDo(http.MethodPut, strurl, i)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bs, _ := io.ReadAll(res.Body)
	return bs, nil
}

func NewCli() *Cli {
	apikey = os.Getenv("REPLAY_APIKEY")
	ret := &Cli{
		cli:     &http.Client{},
		headers: make(map[string]string),
	}
	ret.AddHeader("X-API-KEY", apikey)
	return ret
}
