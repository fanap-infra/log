package log

// Logger is used for logging formatted messages.
type Logger interface {
	// Printf must have the same semantics as log.Printf.
	Printf(format string, args ...interface{})

	Trace(msg string)
	Tracef(format string, args ...interface{})
	Debug(msg string)
	Debugf(format string, args ...interface{})
	Info(msg string)
	Infof(format string, args ...interface{})
	Warn(msg string)
	Warnf(format string, args ...interface{})
	Error(msg string)
	Errorf(format string, args ...interface{})
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

func (l *logger) Trace(msg string)                          {}
func (l *logger) Tracef(format string, args ...interface{}) {}

func (l *logger) Debug(msg string)                          { Debug(msg) }
func (l *logger) Debugf(format string, args ...interface{}) { Debugf(format, args...) }

func (l *logger) Info(msg string)                          { Info(msg) }
func (l *logger) Infof(format string, args ...interface{}) { Infof(format, args...) }

func (l *logger) Warn(msg string)                          { Warn(msg) }
func (l *logger) Warnf(format string, args ...interface{}) { Warnf(format, args...) }

func (l *logger) Error(msg string)                          { Error(msg) }
func (l *logger) Errorf(format string, args ...interface{}) { Errorf(format, args...) }
