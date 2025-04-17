package api

import (
	"go-code/api/network"
	"go-code/jaeger"
)

type App struct {
	Client *jaeger.Client
}

func NewApp(serviceName string) {
	a := &App{}

	a.Client = jaeger.NewClient(serviceName)
	n := network.NewNetwork(a.Client)
	n.Start()
}
