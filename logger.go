package log

import "go.uber.org/zap"

// Logger is used for logging formatted messages.
type Logger interface {
	// Printf must have the same semantics as log.Printf.
	Printf(format string, args ...interface{})

	Trace(msg string)
	Tracef(format string, args ...interface{})

	Debug(msg string)
	Debugf(format string, args ...interface{})
	Debugv(message string, keysValues ...interface{})

	Info(msg string)
	Infof(format string, args ...interface{})
	Infov(message string, keysValues ...interface{})

	Warn(msg string)
	Warnf(format string, args ...interface{})
	Warnv(message string, keysValues ...interface{})

	Error(msg string)
	Errorf(format string, args ...interface{})
	Errorv(message string, keysValues ...interface{})
}

// GetLogger standard logger
func GetLogger() Logger {
	return &l
}

// GetScope standard logger
func GetScope(name string) Logger {
	return &logger{
		suger: suger.Named(name),
	}
}

type logger struct {
	suger *zap.SugaredLogger
}

var l logger = logger{suger: suger}

// Printf must have the same semantics as log.Printf.
func (l *logger) Printf(format string, args ...interface{}) { l.suger.Errorf(format, args...) }

func (l *logger) Trace(msg string)                          {}
func (l *logger) Tracef(format string, args ...interface{}) {}

func (l *logger) Debug(msg string)                          { l.suger.Debug(msg) }
func (l *logger) Debugf(format string, args ...interface{}) { l.suger.Debugf(format, args...) }
func (l *logger) Debugv(message string, keysValues ...interface{}) {
	l.suger.Debugw(message, keysValues...)
}

func (l *logger) Info(msg string)                          { l.suger.Info(msg) }
func (l *logger) Infof(format string, args ...interface{}) { l.suger.Infof(format, args...) }
func (l *logger) Infov(message string, keysValues ...interface{}) {
	l.suger.Infow(message, keysValues...)
}

func (l *logger) Warn(msg string)                          { l.suger.Warn(msg) }
func (l *logger) Warnf(format string, args ...interface{}) { l.suger.Warnf(format, args...) }
func (l *logger) Warnv(message string, keysValues ...interface{}) {
	l.suger.Warnw(message, keysValues...)
}

func (l *logger) Error(msg string)                          { l.suger.Error(msg) }
func (l *logger) Errorf(format string, args ...interface{}) { l.suger.Errorf(format, args...) }
func (l *logger) Errorv(message string, keysValues ...interface{}) {
	l.suger.Errorw(message, keysValues...)
}
