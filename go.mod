module github.com/dc-replay/go-client

go 1.17

//replace go.digitalcircle.com.br/lib/ipcmux => ../../lib/ipcmux
//
//replace go.digitalcircle.com.br/lib/pubsub => ../../lib/pubsub

require (
	github.com/Microsoft/go-winio v0.5.0
	github.com/gorilla/websocket v1.4.2
	github.com/mitchellh/go-ps v1.0.0
)

require golang.org/x/sys v0.0.0-20210910150752-751e447fb3d0 // indirect
