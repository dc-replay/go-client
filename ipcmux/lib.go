package ipcmux

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
)

const (
	ENV_IPC     = "IPC"
	ENV_ROOTPID = "ROOT_PID"
)

var listenNet func(path string) (net.Listener, error)

var dialNet func(path string) (net.Conn, error)

func Dial(p string) (net.Conn, error) {
	return dialNet(p)
}
func SetName(n string) {
	os.Setenv(ENV_IPC, n)
}
func ResolveBinName() string {

	bin := os.Getenv(ENV_IPC)
	if bin == "" {
		bin = os.Args[0]

		if runtime.GOOS == "windows" {
			bin = strings.Replace(bin, "\\", "/", -1)
		}
		parts := strings.Split(bin, "/")
		bin = parts[len(parts)-1]
		bin = strings.Replace(bin, ".exe", "", 1)
		log.Printf("No IPC env var found, using %s to bind ipc", bin)
	}
	return bin
}

func Listen(p string) (net.Listener, error) {
	log.Printf("IPC:: Will listen on: [%s]", p)
	return listenNet(p)
}

func NewClient() *http.Client {
	ret := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, b string) (net.Conn, error) {
				ps := strings.Split(b, ":")
				return dialNet(ps[0])
			},
		},
	}
	return ret
}

func Serve(mux *http.ServeMux) error {
	go CheckRoot()

	bin := ResolveBinName()

	l, err := Listen(bin)
	if err != nil {
		return err
	}
	server := http.Server{
		Handler: mux,
	}
	err = server.Serve(l)
	return err
}

func ServeNamed(n string, mux *http.ServeMux) error {
	go CheckRoot()

	l, err := Listen(n)
	if err != nil {
		return err
	}
	server := http.Server{
		Handler: mux,
	}
	err = server.Serve(l)
	return err
}

func ServeDefault() error {
	return Serve(http.DefaultServeMux)
}
