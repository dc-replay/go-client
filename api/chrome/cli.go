package chrome

import (
	"encoding/json"
	"fmt"
	"github.com/dc-replay/go-client/api"
)

type Cli struct {
	cli *api.Cli
}

func (c *Cli) HttpCli() *api.Cli {
	return c.cli
}

func (c *Cli) Start(to int) error {
	err := c.cli.HttpJsonGet(fmt.Sprintf("/ipc/chrome/start?to=%d", to), nil)
	return err
}
func (c *Cli) StartHeadless() error {
	err := c.cli.HttpJsonGet("/ipc/chrome/startHeadless", nil)
	return err
}
func (c *Cli) Stop() error {
	err := c.cli.HttpJsonGet("/ipc/chrome/stop", nil)
	return err
}
func (c *Cli) New(url string) (string, error) {
	ret := ""
	err := c.cli.HttpJsonGet(fmt.Sprintf("/ipc/chrome/new?url=%s", url), &ret)
	return ret, err
}
func (c *Cli) Close(id string) error {
	err := c.cli.HttpJsonGet("/ipc/chrome/close/"+id, nil)
	return err
}
func (c *Cli) Eval(id string, s string) (map[string]interface{}, error) {
	ret := make(map[string]interface{})
	bs, err := c.cli.HttpRawPost("/ipc/chrome/eval/"+id, []byte(s))
	json.Unmarshal(bs, &ret)
	return ret, err
}

func (c *Cli) Wait(id string, s string, to int) (string, error) {
	ret := ""
	bs, err := c.cli.HttpRawPost(fmt.Sprintf("/ipc/chrome/wait/%s?to=%d", id, to), []byte(s))
	json.Unmarshal(bs, &ret)
	return ret, err
}

func (c *Cli) Send(id string, m string, ps map[string]interface{}) (string, error) {
	ret := ""
	in := map[string]interface{}{
		"method": m,
		"params": ps,
	}
	err := c.cli.HttpJsonPost("/ipc/chrome/eval/"+id, in, &ret)
	return ret, err
}

func NewCli() *Cli {
	ret := &Cli{cli: api.NewCli()}
	return ret
}
