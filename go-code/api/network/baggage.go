package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

func (s *Router) sendForBaggage(c *gin.Context) {
	tracer := opentracing.GlobalTracer()
	rootSpan := newRootSpan("baggage_root_span", c)
	defer rootSpan.Finish()

	rootSpan.SetBaggageItem("greeting", "bagabagegav")

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/receive-for-baggage", nil)
	tracer.Inject(rootSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	c.JSON(http.StatusOK, "Success")
}

func (s *Router) receiveBaggage(c *gin.Context) {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	childSpan := tracer.StartSpan("receive_baggage_span", opentracing.ChildOf(spanCtx))

	defer childSpan.Finish()

	fmt.Println(childSpan.BaggageItem("greeting"))

	c.JSON(http.StatusOK, "Success")

}
