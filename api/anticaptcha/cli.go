package anticaptcha

import (
	"fmt"
	"github.com/dc-replay/go-client/api"
)

type Cli struct {
	*api.Cli
}

func (c *Cli) Recaptchav2(site string, data string) (string, error) {
	ret := ""
	err := c.HttpJsonGet(fmt.Sprintf("/ipc/anticaptcha/recaptchav2?site=%s&data=%s", site, data), &ret)
	return ret, err
}

func (c *Cli) Image2text(site string, data []byte) (string, error) {
	bs, err := c.HttpRawPost(fmt.Sprintf("/ipc/anticaptcha/image2text?site=%s", site), data)
	return string(bs), err
}

func NewCli() *Cli {
	ret := &Cli{Cli: api.NewCli()}
	return ret
}
