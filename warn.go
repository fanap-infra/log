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

// Warnc args and caller point
func Warnc(args ...interface{}) {
	sugerCaller.Warn(args...)
}

// Warncf format message and caller point
func Warncf(format string, args ...interface{}) {
	sugerCaller.Warnf(format, args...)
}

// Warncv message with key=value ... and caller point
func Warncv(msg string, keysValues ...interface{}) {
	sugerCaller.Warnw(msg, keysValues...)
}
