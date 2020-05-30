package rtsp

import (
	"time"

	"gitlab.com/behnama2/log"
)

func GetPacketFunc() {
	//log.Info("Namitonam fetch konam")
	//
	//log.Infof("Namitonam fetch konam %s", "Hossein")

	log.Infoc("Namitonam",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)

	log.Errorc("Namitonam",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
}
