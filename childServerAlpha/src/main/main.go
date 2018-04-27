package main

import (
	"fmt"
	opentracing "github.com/DroiTaipei/opentracing-go"
	zipkin "github.com/DroiTaipei/zipkin-go-opentracing"

	"net/http"
	"os"
	"strings"
	"time"

	"github.com/DroiTaipei/droipkg/trace"
	"github.com/gin-gonic/gin"
)

const (
	Host          = "10.128.80.39"
	ComponentName = "Alpha"
	ZipkinPort    = 9411
	EntryPort     = 10000
	AlphaPort     = 10001
	BetaPort      = 10002

	ZipkinTopic = "zipkin"

	BagKey = "TestBag"
)

var (
	ZipkinKafkaAddr = []string{
		"10.128.112.186:9092",
		"10.128.112.184:9092",
	}
)

type Response struct {
	Code   int
	Result string
}

func initZipkin() error {
	var tracer opentracing.Tracer
	collector, err := zipkin.NewKafkaCollector(ZipkinKafkaAddr, zipkin.KafkaTopic(ZipkinTopic))
	// HTTP Collector
	//collectorUrl := fmt.Sprintf("http://%s:%d/api/v1/spans", Host, ZipkinPort)
	//collector, err := zipkin.NewHTTPCollector(collectorUrl)
	if err != nil {
		return err
	}
	sampler := zipkin.NewBoundarySampler(0.2, 123)
	tracer, err = zipkin.NewTracer(
		zipkin.NewRecorder(collector, true, fmt.Sprintf("%s:%d", Host, BetaPort), ComponentName),
		zipkin.ClientServerSameSpan(true),
		zipkin.TraceID128Bit(true),
		zipkin.WithSampler(sampler),
	)
	if err != nil {
		return err
	}
	opentracing.SetGlobalTracer(tracer)
	return nil
}

func initJaeger() error {
	// HTTP Collector
	//collectorUrl := fmt.Sprintf("http://%s:%d/api/v1/spans", Host, ZipkinPort)
	//collector, err := zipkin.NewHTTPCollector(collectorUrl)

	host := fmt.Sprintf("%s:%d", Host, EntryPort)
	//sampler := zipkin.NewBoundarySampler(0.2, 0)
	if err := droitrace.InitJaeger(host, ComponentName); err != nil {
		return err
	}
	return nil
}

func reqBeta(c *gin.Context, pSpan opentracing.Span) error {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%d/beta", Host, BetaPort)
	req, _ := http.NewRequest("GET", url, nil)
	// need to create a span and inject it into req's header
	setDroiHeader(c, req)
	span := droitrace.CreateChildSpan(pSpan, req)
	defer span.Finish()
	droitrace.InjectSpan(span, req)
	droitrace.SetRPCClientTag(span)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad status code: %d", res.StatusCode)
	}
	return nil
}

func AlphaHandler(c *gin.Context) {
	span := droitrace.CreateSpanFromReq(c.Request)
	// should call Finish so the tracker will handle it
	defer span.Finish()

	testBagVal := span.BaggageItem(strings.ToLower(BagKey))
	if len(testBagVal) > 0 {
		span.SetTag("test.bag", testBagVal)
	} else {
		span.SetTag("test.bag", "None")
	}
	if err := reqBeta(c, span); err != nil {
		span.SetTag("http.status", http.StatusBadRequest)
		span.SetTag("http.error", err.Error())
		span.SetTag("error", true)
		c.JSON(http.StatusBadRequest, &Response{
			Code:   -1,
			Result: "trace FAILED",
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Code:   0,
		Result: "trace done",
	})
	return
}

func main() {
	if err := initJaeger(); err != nil {
		fmt.Fprintf(os.Stderr, "initJaeger failed, err:%s\n", err)
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET(fmt.Sprintf("/alpha"), AlphaHandler)
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", AlphaPort),
		Handler:      r,
		ReadTimeout:  time.Duration(60) * time.Second,
		WriteTimeout: time.Duration(60) * time.Second,
	}
	if err := s.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "Server started error: %s\n", err)
		os.Exit(0)
	}
}
