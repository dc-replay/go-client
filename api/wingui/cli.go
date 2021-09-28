package wingui

import (
	"fmt"
	"github.com/dc-replay/go-client/api"
)

type Cli struct {
	cli *api.Cli
}

func (c *Cli) HttpCli() *api.Cli {
	return c.cli
}

func (c *Cli) ClipRead() (string, error) {
	ret := ""
	err := c.cli.HttpJsonGet(fmt.Sprintf("/ipc/wingui/clip/read"), &ret)
	return ret, err
}

func (c *Cli) ClipWrite(site string) error {
	err := c.cli.HttpJsonGet(fmt.Sprintf("/ipc/wingui/clip/write?str=%s", site), nil)
	return err
}

func (c *Cli) MouseClick() error {
	err := c.cli.HttpJsonGet(fmt.Sprintf("/ipc/wingui/mouse/click"), nil)
	return err
}
func (c *Cli) MouseMove(x, y int) error {
	err := c.cli.HttpJsonGet(fmt.Sprintf("/ipc/wingui/mouse/move?x=%d&y=%d", x, y), nil)
	return err
}
func (c *Cli) MouseMoveRelative(x, y int) error {
	err := c.cli.HttpJsonGet(fmt.Sprintf("/ipc/wingui/mouse/moverelative?x=%d&y=%d", x, y), nil)
	return err
}
func (c *Cli) MouseClickRelative(x, y int) error {
	err := c.cli.HttpJsonGet(fmt.Sprintf("/ipc/wingui/mouse/clickatrelative?x=%d&y=%d", x, y), nil)
	return err
}
func (c *Cli) MouseClickAt(x, y int) error {
	err := c.cli.HttpJsonGet(fmt.Sprintf("/ipc/wingui/mouse/clickat?x=%d&y=%d", x, y), nil)
	return err
}
func (c *Cli) ScreenClick(f string) error {
	err := c.cli.HttpJsonGet(fmt.Sprintf("/ipc/wingui/screen/click?f=%s", f), nil)
	return err
}
func (c *Cli) ScreenClickCenter(f string) error {
	err := c.cli.HttpJsonGet(fmt.Sprintf("/ipc/wingui/screen/clickcenter?f=%s", f), nil)
	return err
}
func (c *Cli) ScreenFind(f string) (ret map[string]interface{}, err error) {
	ret = make(map[string]interface{})
	err = c.cli.HttpJsonGet(fmt.Sprintf("/ipc/wingui/screen/find?f=%s", f), &ret)
	return
}

func (c *Cli) ScreenWait(f string, m int) (ret map[string]interface{}, err error) {
	ret = make(map[string]interface{})
	err = c.cli.HttpJsonGet(fmt.Sprintf("/ipc/wingui/screen/wait?f=%s&m=%d", f, m), &ret)
	return
}
func (c *Cli) ScreenWaitClick(f string, m int) (ret map[string]interface{}, err error) {
	ret = make(map[string]interface{})
	err = c.cli.HttpJsonGet(fmt.Sprintf("/ipc/wingui/screen/waitclick?f=%s&m=%d", f, m), &ret)
	return
}
func (c *Cli) ScreenWaitClickCenter(f string, m int) (ret map[string]interface{}, err error) {
	ret = make(map[string]interface{})
	err = c.cli.HttpJsonGet(fmt.Sprintf("/ipc/wingui/screen/waitclickcenter?f=%s&m=%d", f, m), &ret)
	return
}

func NewCli() *Cli {
	ret := &Cli{cli: api.NewCli()}
	return ret
}
