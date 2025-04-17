package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

func (s *Router) sendWithOtherHost(c *gin.Context) {

	tracer := opentracing.GlobalTracer()
	rootSpan := newRootSpan("other_host_root_span", c)
	defer rootSpan.Finish()

	fmt.Println("send with other")
	fmt.Println("send header", c.Request.Header)
	fmt.Println()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/receive-one-from-other-host", nil)
	tracer.Inject(rootSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	c.JSON(http.StatusOK, "Success")
}

//"/send-other-host"
//"/receive-two-from-other-host"

func (s *Router) receiveOne(c *gin.Context) {
	fmt.Println("receive header", c.Request.Header)

	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	childSpan := tracer.StartSpan("receive_one_span", opentracing.ChildOf(spanCtx))

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/receive-two-from-other-host", nil)
	tracer.Inject(childSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	defer childSpan.Finish()

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	c.JSON(http.StatusOK, gin.H{"message": "Span receive"})

}

func (s *Router) receiveTwo(c *gin.Context) {
	fmt.Println("receive header", c.Request.Header)

	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	childSpan := tracer.StartSpan("receive_two_span", opentracing.ChildOf(spanCtx))
	defer childSpan.Finish()

	c.JSON(http.StatusOK, gin.H{"message": "Span receive"})
}
