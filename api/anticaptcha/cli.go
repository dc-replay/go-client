package anticaptcha

import (
	"github.com/dc-replay/go-client/api"
)

type Cli struct {
	*api.Cli
}

type Req struct {
	Site string `json:"site"`
	Data string `json:"data"`
	Img  []byte `json:"img"`
	To   int    `json:"to"`
}

func (c *Cli) Recaptchav2(site string, data string) (string, error) {
	ret := ""
	req := &Req{
		Site: site,
		Data: data,
		Img:  nil,
		To:   300,
	}
	err := c.HttpJsonPost("/ipc/anticaptcha/recaptchav2", req, &ret)
	return ret, err
}

func (c *Cli) Image2text(site string, data []byte) (string, error) {
	ret := ""
	req := &Req{
		Site: site,
		Img:  data,
		To:   300,
	}
	err := c.HttpJsonPost("/ipc/anticaptcha/image2text", req, &ret)
	return ret, err
}

func NewCli() *Cli {
	ret := &Cli{Cli: api.NewCli()}
	return ret
}
