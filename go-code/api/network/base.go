package network

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

func newRootSpan(name string, c *gin.Context) opentracing.Span {
	tracer := opentracing.GlobalTracer()
	spenCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sendSpan := tracer.StartSpan(name, ext.RPCServerOption(spenCtx))

	defer sendSpan.Finish()
	return sendSpan
}

func (s *Router) send(c *gin.Context) {
	newRootSpan("send_root_span", c)

	c.JSON(http.StatusOK, "Success")
}

func (s *Router) sendWithTag(c *gin.Context) {

}

func (s *Router) sendWithChild(c *gin.Context) {

}
