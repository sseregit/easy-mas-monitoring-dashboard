package network

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
)

func (s *Router) sendForPanic(c *gin.Context) {
	tracer := opentracing.GlobalTracer()
	rootSpan := newRootSpan("root_span_err", c)
	defer rootSpan.Finish()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/receive-for-error", nil)
	tracer.Inject(rootSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	c.JSON(http.StatusOK, "Success")
}

func (s *Router) receiveForError(c *gin.Context) {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	childSpan := tracer.StartSpan("receive_one_span", opentracing.ChildOf(spanCtx))

	defer childSpan.Finish()

	childSpan.SetTag("error", true)
	childSpan.LogFields(log.String("event", "error"), log.String("message", "우와 에렁"))

	c.JSON(http.StatusOK, "Success")

}
