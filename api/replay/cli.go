package replay

import (
	"fmt"
	"os"

	"github.com/dc-replay/go-client/api"
	"github.com/dc-replay/go-client/types"
	"github.com/dc-replay/go-client/util"
)

type Cli struct {
	cli *api.Cli
}

func (c *Cli) HttpCli() *api.Cli {
	return c.cli
}

func (c *Cli) OpenApp(id uint) error {
	err := c.cli.HttpJsonGet(fmt.Sprintf("/api/v1/app/run/%d", id), nil)
	return err
}

func (c *Cli) ConfigGet(k string) (ret string, err error) {
	err = c.cli.HttpJsonGet("/api/v1/config/byrepo/"+util.Repo()+"/"+k, &ret)
	return
}

func (c *Cli) ConfigGetAll() (ret []types.Config, err error) {
	ret = make([]types.Config, 0)
	err = c.cli.HttpJsonGet("/api/v1/config/byrepo/"+util.Repo(), &ret)
	return
}

func (c *Cli) CronReload() error {
	return c.cli.HttpJsonGet("/ipc/cron/reload", nil)
}

func (c *Cli) Exit() error {
	err := c.cli.HttpJsonGet("/api/v1/exit", nil)
	return err
}

func (c *Cli) MenuGetAllEnabled() (ret []types.Menu, err error) {
	ret = make([]types.Menu, 0)
	err = c.cli.HttpJsonGet("/api/v1/menu/enabled", nil)
	return ret, err
}

func (c *Cli) QueueAdd(job string, bs []byte) error {
	return c.cli.HttpJsonGet("/api/v1/queue/add/"+job, bs)
}
func (c *Cli) QueueGetData(id string) (ret []byte, err error) {
	ret, err = c.cli.HttpRawGet("/api/v1/queue/getrequest/" + id)
	return
}
func (c *Cli) QueueGetMyData() ([]byte, error) {
	return c.QueueGetData(os.Getenv("REPLAY_QUEUEID"))
}
func (c *Cli) QueueAbort() error {
	return c.cli.HttpJsonGet("/api/v1/queue/abort", nil)
}
func (c *Cli) QueueEnqueue(id uint) error {
	return c.cli.HttpJsonGet(fmt.Sprintf("/api/v1/robots/op/enqueue/%d", id), nil)
}

func (c *Cli) ServiceStopAll() error {
	return c.cli.HttpJsonGet("/api/v1/service/op/stopall", nil)
}

type SQLReturn struct {
	Data []map[string]interface{} `json:"data"`
	Err  string                   `json:"err"`
}

func (c *Cli) SQL(s string) (*SQLReturn, error) {
	in := &struct {
		Sql string `json:"sql"`
	}{s}
	out := &SQLReturn{}
	err := c.cli.HttpJsonPost("/api/v1/sql", in, &out)
	return out, err
}

func NewCli() *Cli {
	ret := &Cli{cli: api.NewCli()}
	return ret
}
