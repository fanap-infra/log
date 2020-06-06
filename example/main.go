package main

import (
	"time"

	"github.com/fanap-infra/log"
	"github.com/fanap-infra/log/example/rtsp"
)

func main() {
	defer log.Sync()
	log.RedirectStdLog()
	log.Config(log.DebugLevel, true)

	// log.GetLogger()
	log.Warn("Not Found config file")

	log.Infov("GET",
		"url", "http://example.com/data.json",
	)
	log.Errorv("Fetch",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)

	rtsp.GetPacketFunc()
}
