package main

import (
	"github.com/tpl-x/conec/internal/server"
)

type app struct {
	pingSvr *server.PingServer
}

func newApp(pingSvr *server.PingServer) *app {
	return &app{pingSvr: pingSvr}
}

func (a *app) ListenAndServe() error {
	return a.pingSvr.ListenAndServe()
}

func main() {
	server, err := wireApp()
	if err != nil {
		panic(err)
	}
	server.ListenAndServe()
}
