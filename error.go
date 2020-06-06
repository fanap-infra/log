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

// Errorc args and caller point
func Errorc(args ...interface{}) {
	sugerCaller.Error(args...)
}

// Errorcf format message and caller point
func Errorcf(format string, args ...interface{}) {
	sugerCaller.Errorf(format, args...)
}

// Errorcv message with key=value ... and caller point
func Errorcv(msg string, keysValues ...interface{}) {
	sugerCaller.Errorw(msg, keysValues...)
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

// ErrorcIF error and caller point
func ErrorcIF(err error) {
	if err != nil {
		sugerCaller.Error(err)
	}
}

// ErrorcmIF error with message and caller point
func ErrorcmIF(msg string, err error) {
	if err != nil {
		sugerCaller.Errorw(msg, "error", err)
	}
}
