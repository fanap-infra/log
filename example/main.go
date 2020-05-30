package main

import (
	"time"

	"gitlab.com/behnama2/log"
	"gitlab.com/behnama2/log/example/rtsp"
)

func main() {
	log.Init(log.DebugLevel)
	defer log.Close()

	log.Warn("Not Found config file")

	log.Infov("GET",
		"url", "http://example.com/data.json",
	)

	log.Errorc("Fetch",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)

	rtsp.GetPacketFunc()
}
