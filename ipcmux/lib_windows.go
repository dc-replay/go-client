//go:build windows
// +build windows

package ipcmux

import (
	"errors"
	"github.com/Microsoft/go-winio"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	listenNet = func(path string) (net.Listener, error) {
		path = strings.Replace(path, "/", "\\", -1)
		log.Printf("Listening to: " + path)

		chcon := make(chan bool)
		var conn net.Listener
		var err error
		go func() {
			path = "\\\\.\\pipe\\" + path
			conn, err = winio.ListenPipe(path, &winio.PipeConfig{
				SecurityDescriptor: "",
				MessageMode:        false,
				InputBufferSize:    0,
				OutputBufferSize:   0,
			})
			if err != nil {
				log.Printf("ipcmux::listenNet: %s", err.Error())
			}
			go func() {
				go func() {
					to := time.Second * time.Duration(3)
					fconn, err := winio.DialPipe(path, &to)
					if err != nil {
						log.Printf("Error on pipe warm up: %s", err.Error())
					} else {
						fconn.Close()
					}
					log.Printf("Pipe warm up done")
				}()
			}()
			chcon <- true
		}()

		select {
		case <-chcon:
			return conn, err
		case <-time.After(time.Second * 5):
			err = errors.New("Timeout exceeded for Listening: " + path)
			return nil, err
		}

		return conn, err
	}
	dialNet = func(path string) (net.Conn, error) {
		path = strings.Replace(path, "/", "\\", -1)
		path = "\\\\.\\pipe\\" + path
		chcon := make(chan bool)
		var conn net.Conn
		var err error
		//Log("Dialing to: " + path)
		go func() {
			to := time.Second * time.Duration(15)
			conn, err = winio.DialPipe(path, &to)
			//conn, err = npipe.Dial(path)
			chcon <- true
		}()

		select {
		case <-chcon:
			return conn, err
		case <-time.After(time.Second * 5):
			err = errors.New("Timeout exceeded for conn: " + path)
			return nil, err
		}

		return conn, err
	}
}

func CheckRoot() {
	time.Sleep(time.Second)
	rootpid := os.Getenv(ENV_ROOTPID)
	rootpidi, err := strconv.Atoi(rootpid)
	if err != nil {
		log.Printf("No root process found")
	}
	if rootpidi == 0 || rootpidi == os.Getpid() {
		log.Printf("%s is set to zero, no monitoring on root will take place", ENV_ROOTPID)
		return
	}

	p, err := os.FindProcess(rootpidi)

	if err != nil {
		log.Printf("Error finding process: %s", err.Error())
		os.Exit(1)
	}
	if p == nil {
		log.Printf("Process not found.")
		os.Exit(1)
	}
	_, err = p.Wait()
	if err != nil {
		log.Printf(err.Error())
	}
	log.Printf("Parent process %s is gone. Finishing here.", os.Getenv(ENV_ROOTPID))
	os.Exit(0)

}
