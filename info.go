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

// Infoc args and caller point
func Infoc(args ...interface{}) {
	sugerCaller.Info(args...)
}

// Infocf format message and caller point
func Infocf(format string, args ...interface{}) {
	sugerCaller.Infof(format, args...)
}

// Infocv message with key=value ... and caller point
func Infocv(msg string, keysValues ...interface{}) {
	sugerCaller.Infow(msg, keysValues...)
}
