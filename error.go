package log

// Error args
func Error(args ...interface{}) {
	suger.Error(args...)
}

// Errorf format message
func Errorf(format string, args ...interface{}) {
	suger.Errorf(format, args...)
}

// Errorv message with key=value ...
func Errorv(message string, keysValues ...interface{}) {
	suger.Errorw(message, keysValues...)
}

// ErrorIF error
func ErrorIF(err error) {
	if err != nil {
		suger.Error(err)
	}
}

// ErrormIF error with message
func ErrormIF(msg string, err error) {
	if err != nil {
		suger.Errorw(msg, "error", err)
	}
}
