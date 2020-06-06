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
