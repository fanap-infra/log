package log

// Logger is used for logging formatted messages.
type Logger interface {
	// Printf must have the same semantics as log.Printf.
	Printf(format string, args ...interface{})
}

// GetLogger standard logger
func GetLogger() Logger {
	return &l
}

type logger struct{}

var l logger

// Printf must have the same semantics as log.Printf.
func (l *logger) Printf(format string, args ...interface{}) {
	Errorf(format, args...)
}
