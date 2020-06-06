# Log

```
package main

import "gitlab.com/behnama2/log"

func main() {
	defer log.Sync()
	log.RedirectStdLog()
	log.SetLevel(log.DebugLevel)

	// log.GetLogger()

	log.Warn("Not Found config file")

	log.Infov("GET",
		"url", "http://example.com/data.json",
	)

	log.Errorc("Fetch",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
}
```
