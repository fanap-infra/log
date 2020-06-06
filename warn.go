package log

// Warn args
func Warn(args ...interface{}) {
	suger.Warn(args...)
}

// Warnf format message
func Warnf(format string, args ...interface{}) {
	suger.Warnf(format, args...)
}

// Warnv message with key=value ...
func Warnv(message string, keysValues ...interface{}) {
	suger.Warnw(message, keysValues...)
}
