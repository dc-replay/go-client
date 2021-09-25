package excel

import (
	"github.com/dc-replay/go-client/api"
)

const (
	CELLTYPE_STRING = "s"
	CELLTYPE_INT    = "i"
	CELLTYPE_BOOL   = "b"
)

type Cli struct {
	cli *api.Cli
}

type Req struct {
	File    string `json:"file"`
	Sheet   string `json:"sheet"`
	Cel     string `json:"cel"`
	Val     string `json:"val"`
	Celtype string `json:"celtype"`
}

func (c *Cli) Read(fname string, sheet string) ([][]string, error) {
	req := &Req{
		File:  fname,
		Sheet: sheet,
	}
	res := make([][]string, 0)
	err := c.cli.HttpJsonPost("/ipc/excel/read", req, &res)
	return res, err
}
func (c *Cli) Write(fname string, sheet string, cell string, val string, celtype string) error {
	req := &Req{
		File:    fname,
		Sheet:   sheet,
		Cel:     cell,
		Val:     val,
		Celtype: celtype,
	}
	res := ""
	err := c.cli.HttpJsonPost("/ipc/excel/write", req, &res)
	return err
}

func (c *Cli) New(fname string, sheet string) (string, error) {
	req := &Req{
		File:  fname,
		Sheet: sheet,
	}
	res := ""
	err := c.cli.HttpJsonPost("/ipc/excel/new", req, &res)
	return res, err
}

func NewCli() *Cli {
	ret := &Cli{cli: api.NewCli()}
	return ret
}
