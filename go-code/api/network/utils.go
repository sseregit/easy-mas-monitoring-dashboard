package network

import "github.com/gin-gonic/gin"

type R int8

const (
	GET R = iota
	POST
	DELETE
	PUT
)

func (n *Network) Router(r R, path string, handler gin.HandlerFunc) {
	switch r {
	case GET:
		n.engin.GET(path, handler)
	case POST:
		n.engin.POST(path, handler)
	case DELETE:
		n.engin.DELETE(path, handler)
	case PUT:
		n.engin.PUT(path, handler)
	}

}
