package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptrace"
	"time"
)

func main() {
	tr := &http.Transport{
		// Client keeps idle conns around for a long time (bigger than server).
		IdleConnTimeout:     10 * time.Minute,
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 10,

		// Optional: keep TCP keepalive fairly long (not the same as HTTP idle pooling,
		// but included since you mentioned "high keep alive").
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 5 * time.Minute,
		}).DialContext,
	}

	c := &http.Client{Transport: tr}

	url := "http://127.0.0.1:8080/"

	// 1) First request creates a connection
	log.Println("REQ #1 (creates conn)")
	doReq(c, url)

	// 2) Let it sit idle long enough for the *server* to close it
	log.Println("sleeping 70s so server closes idle conn (server IdleTimeout=60s)")
	time.Sleep(70 * time.Second)

	// 3) Second request tries to reuse the pooled conn -> EOF / closed conn error
	log.Println("REQ #2 (client tries to reuse idle pooled conn) -> expect boom")
	doReq(c, url)
}

func doReq(c *http.Client, url string) {
	ctx := context.Background()

	trace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			log.Printf("GotConn: reused=%v wasIdle=%v idleTime=%v",
				info.Reused, info.WasIdle, info.IdleTime)
		},
		PutIdleConn: func(err error) {
			log.Printf("PutIdleConn: err=%v", err)
		},
	}

	req, _ := http.NewRequestWithContext(httptrace.WithClientTrace(ctx, trace), "GET", url, nil)
	resp, err := c.Do(req)
	if err != nil {
		log.Printf("Do error: %T: %v", err, err)
		return
	}
	defer resp.Body.Close()

	var buf [64]byte
	n, rerr := resp.Body.Read(buf[:])
	log.Printf("status=%s read=%q rerr=%v", resp.Status, string(buf[:n]), rerr)
	fmt.Println()
}
