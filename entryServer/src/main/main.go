package main

import (
	"fmt"
	opentracing "github.com/DroiTaipei/opentracing-go"
	zipkin "github.com/DroiTaipei/zipkin-go-opentracing"

	"net/http"
	"os"
	"sync"
	"time"

	droitrace "github.com/DroiTaipei/droitrace"
	"github.com/gin-gonic/gin"
)

const (
	Host          = "10.128.80.39"
	ComponentName = "Entry"
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

type DummyLog struct {
	Code    int
	Message string
}

type Response struct {
	Code   int
	Result string
}

func initZipkin() error {
	// HTTP Collector
	//collectorUrl := fmt.Sprintf("http://%s:%d/api/v1/spans", Host, ZipkinPort)
	//collector, err := zipkin.NewHTTPCollector(collectorUrl)

	// Kafka Collector
	collector, err := zipkin.NewKafkaCollector(ZipkinKafkaAddr, zipkin.KafkaTopic(ZipkinTopic))
	if err != nil {
		return err
	}
	host := fmt.Sprintf("%s:%d", Host, EntryPort)
	sampler := zipkin.NewBoundarySampler(0.2, 0)
	if err := droitrace.InitZipkin(collector, sampler, host, ComponentName); err != nil {
		return err
	}
	return nil
}

func initJaeger() error {
	// HTTP Collector
	//collectorUrl := fmt.Sprintf("http://%s:%d/api/v1/spans", Host, ZipkinPort)
	//collector, err := zipkin.NewHTTPCollector(collectorUrl)
	//sampler := zipkin.NewBoundarySampler(0.2, 0)
	if err := droitrace.InitJaeger(1, ComponentName); err != nil {
		return err
	}
	return nil
}

func reqAlpha(c *gin.Context, pSpan opentracing.Span, wg *sync.WaitGroup) error {
	defer wg.Done()

	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%d/alpha", Host, AlphaPort)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// need to create a span and inject it into req's header
	setDroiHeader(c, req)
	span := droitrace.CreateChildSpan(pSpan, req)
	defer span.Finish()
	droitrace.SetRPCClientTag(span)
	droitrace.InjectSpan(span, req)

	for k, v := range req.Header {
		fmt.Println(k, v)
	}

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

func reqBeta(c *gin.Context, pSpan opentracing.Span, wg *sync.WaitGroup) error {
	defer wg.Done()

	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%d/beta", Host, BetaPort)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Step 3.
	// Before request other components (ex: Beta here), start a childOf/followFrom span
	// by droitrace.CreateChildSpan()/droitrace.CreateFollowsFromSpan().
	// Call droitrace.SetRPCClientTag() will help you set client-related tag.
	// Call droitrace.InjectSpan() will inject the span on the carrier, the request header
	// *** Remember to Finish any span you create ***
	setDroiHeader(c, req)
	span := droitrace.CreateChildSpan(pSpan, req)
	defer span.Finish()
	span.LogKV("TestLogKey", DummyLog{Code: 100, Message: "Log test"})
	droitrace.SetRPCClientTag(span)
	droitrace.InjectSpan(span, req)
	// Done
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

func HelloTraceHandler(c *gin.Context) {
	var wg sync.WaitGroup

	// Step 2.
	// start a span by droitrace.CreateSpanFromReq()
	// The function will extract the span context carriered by request.
	// - not exist or corrupted: create a root span (new trace)
	// - exist: create a new span inherit the carried span, attach server-related tags (if droi headers exist, attach to tags autoly)
	span := droitrace.CreateSpanFromReq(c.Request)
	// should call Finish so the tracker will handle it
	defer span.Finish()
	span.SetBaggageItem(BagKey, "Hello,This bag will tranverse the components on the trace")

	wg.Add(1)
	go reqAlpha(c, span, &wg)
	wg.Add(1)
	go reqBeta(c, span, &wg)
	wg.Wait()

	c.JSON(http.StatusOK, &Response{
		Code:   0,
		Result: "trace done",
	})
	return
}

func main() {
	// Step 1.
	// init zipkin global tracker to record finished span and send it to zipkin server.
	// tick timer up or meet batch size will trigger send spans
	if err := initJaeger(); err != nil {
		fmt.Fprintf(os.Stderr, "initJaeger failed, err:%s\n", err)
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET(fmt.Sprintf("/hellotrace"), HelloTraceHandler)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", EntryPort),
		Handler:      r,
		ReadTimeout:  time.Duration(60) * time.Second,
		WriteTimeout: time.Duration(60) * time.Second,
	}
	if err := s.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "Server started error: %s\n", err)
		os.Exit(0)
	}
}
