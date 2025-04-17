package network

import (
	"github.com/gin-gonic/gin"
	"go-code/jaeger"
)

type Network struct {
	client *jaeger.Client
	engin  *gin.Engine
}

func NewNetwork(client *jaeger.Client) *Network {
	n := &Network{client: client, engin: gin.New()}

	return n
}

func (n *Network) Start() {
	n.engin.Run(":8080")
}
