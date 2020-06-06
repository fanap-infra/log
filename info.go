package log

// Info args
func Info(args ...interface{}) {
	suger.Info(args...)
}

// Infof format message
func Infof(format string, args ...interface{}) {
	suger.Infof(format, args...)
}

// Infov message with key=value ...
func Infov(message string, keysValues ...interface{}) {
	suger.Infow(message, keysValues...)
}
