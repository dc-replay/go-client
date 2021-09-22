//go:build !windows
// +build !windows

package ipcmux

import (
	"github.com/mitchellh/go-ps"
	"log"
	"net"
	"os"
	libpath "path"
	"strconv"
	"time"
)

func init() {

	listenNet = func(path string) (net.Listener, error) {
		ipc_root := os.Getenv("IPC_ROOT")
		if ipc_root == "" {
			ipc_root = os.TempDir()
		}

		path = "replayme_" + path
		path = libpath.Join(ipc_root, path)
		err := os.Remove(path)
		if err != nil {
			if !os.IsNotExist(err) {
				log.Printf("Err removing old uds: %s", err.Error())
			}
		}
		l, err := net.Listen("unix", path)
		log.Printf("UnixSocket::%s=>%v", path, err)
		return l, err

	}
	dialNet = func(path string) (net.Conn, error) {
		ipc_root := os.Getenv("IPC_ROOT")
		if ipc_root == "" {
			ipc_root = os.TempDir()
		}
		path = "replayme_" + path
		path = libpath.Join(ipc_root, path)
		conn, err := net.Dial("unix", path)

		return conn, err
	}
}

const (
	PRCTL_SYSCALL    = 157
	PR_SET_PDEATHSIG = 1
)

func CheckRoot() {

	time.Sleep(time.Second)
	rootpid := os.Getenv(ENV_ROOTPID)
	rootpidi, err := strconv.Atoi(rootpid)
	if err != nil {
		log.Printf("No root process found")
		return
	}
	if rootpidi == 0 || rootpidi == os.Getpid() {
		log.Printf("%s is set to zero, no monitoring on root will take place", ENV_ROOTPID)
		return
	}
	for {
		p, err := ps.FindProcess(rootpidi)
		if p == nil || err != nil {
			log.Printf("aborting")
			os.Exit(1)
		}
		time.Sleep(time.Second)
	}
}
