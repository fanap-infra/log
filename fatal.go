package log

// Fatal args
func Fatal(args ...interface{}) {
	suger.Fatal(args...)
}

// Fatalf format message
func Fatalf(format string, args ...interface{}) {
	suger.Fatalf(format, args...)
}

// Fatalv message with key=value ...
func Fatalv(message string, keysValues ...interface{}) {
	suger.Fatalw(message, keysValues...)
}

// Fatalc args and caller point
func Fatalc(args ...interface{}) {
	sugerCaller.Fatal(args...)
}

// Fatalcf format message and caller point
func Fatalcf(format string, args ...interface{}) {
	sugerCaller.Fatalf(format, args...)
}

// Fatalcv message with key=value ... and caller point
func Fatalcv(msg string, keysValues ...interface{}) {
	sugerCaller.Fatalw(msg, keysValues...)
}

// FatalIF error
func FatalIF(err error) {
	if err != nil {
		suger.Fatal(err)
	}
}

// FatalmIF error with message
func FatalmIF(msg string, err error) {
	if err != nil {
		suger.Fatalw(msg, "error", err)
	}
}

// FatalcIF error and caller point
func FatalcIF(err error) {
	if err != nil {
		sugerCaller.Fatal(err)
	}
}

// FatalcmIF error with message and caller point
func FatalcmIF(msg string, err error) {
	if err != nil {
		sugerCaller.Fatalw(msg, "error", err)
	}
}
