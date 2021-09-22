package pubsub

import (
	"encoding/json"
	"fmt"
	"github.com/dc-replay/go-client/ipcmux"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"os"
	"os/user"
	"strings"
	"time"
)

const XMSGTYPE = "X-MSGTYPE"
const XPUBSUBID = "X-PUBSUBID"

var PSID = ""
var dialer *websocket.Dialer

func init() {
	dialer = &websocket.Dialer{}
	hname, _ := os.Hostname()
	pid := os.Getppid()
	u, _ := user.Current()
	PSID = fmt.Sprintf("psid_%s_%s_%d", hname, u.Username, pid)

	dialer.NetDial = func(network, addr string) (net.Conn, error) {
		ps := strings.Split(addr, ":")
		c, err := ipcmux.Dial(ps[0])
		return c, err
	}
}

type PSCli struct {
	q      string
	con    *websocket.Conn
	ch     chan []byte
	closed bool
}

func (p *PSCli) GetConn() *websocket.Conn {
	if p.con == nil {

		var counter = 0
		for counter < 100 {
			con, _, err := dialer.Dial("ws://pubsub/?q="+p.q, nil)

			if err != nil {
				log.Printf("Error connecting WS: %s", err.Error())
				counter++
				time.Sleep(time.Second)
				continue
			}
			p.con = con
			p.con.SetCloseHandler(func(code int, text string) error {
				log.Printf("Connection closed: %d: %s", code, text)
				p.con = nil
				return nil
			})
			break
		}

	}
	return p.con
}

func (p *PSCli) Send(i interface{}) error {
	bs, err := json.Marshal(i)
	if err != nil {
		return err
	}
	err = p.GetConn().WriteMessage(websocket.TextMessage, bs)
	if err != nil {
		p.con.Close()
		p.con = nil
	}
	return err
}

func (p *PSCli) Chan() chan []byte {
	return p.ch
}

func (p *PSCli) Read(i interface{}) error {
	_, bs, err := p.GetConn().ReadMessage()
	if err != nil {
		if p.con != nil {
			p.con.Close()
			p.con = nil
		}
		return err
	}
	return json.Unmarshal(bs, i)
}

func NewCli(q string) (*PSCli, error) {
	ret := &PSCli{}
	ret.q = q
	return ret, nil
}
