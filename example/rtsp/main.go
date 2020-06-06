package rtsp

import (
	"time"

	"gitlab.com/behnama2/log"
)

func GetPacketFunc() {
	//log.Info("Namitonam fetch konam")
	//
	//log.Infof("Namitonam fetch konam %s", "Hossein")
	log.Warn("Not Found config file")

	log.Infov("GET",
		"url", "http://example.com/data.json",
	)
	log.Errorv("Fetch",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)

	log.Infov("Namitonam",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)

	log.Errorv("Namitonam",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
}
