package main

import (
	"time"

	"gitlab.com/behnama2/log"
	"gitlab.com/behnama2/log/example/rtsp"
)

func main() {
	defer log.Sync()
	log.RedirectStdLog()
	log.SetLevel(log.DebugLevel)

	// log.GetLogger()


	log.Warn("Not Found config file")

	log.Infov("GET",
		"url", "http://example.com/data.json",
	)

	log.Errorcv("Fetch",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)

	rtsp.GetPacketFunc()
}
