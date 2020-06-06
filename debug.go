package log

// Debug args
func Debug(args ...interface{}) {
	suger.Debug(args...)
}

// Debugf format message
func Debugf(format string, args ...interface{}) {
	suger.Debugf(format, args...)
}

// Debugv message with key=value ...
func Debugv(message string, keysValues ...interface{}) {
	suger.Debugw(message, keysValues...)
}

// Debugc args and caller point
func Debugc(args ...interface{}) {
	sugerCaller.Debug(args...)
}

// Debugcf format message and caller point
func Debugcf(format string, args ...interface{}) {
	sugerCaller.Debugf(format, args...)
}

// Debugcv message with key=value ... and caller point
func Debugcv(msg string, keysValues ...interface{}) {
	sugerCaller.Debugw(msg, keysValues...)
}
