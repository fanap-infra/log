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
