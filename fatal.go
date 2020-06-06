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

// FatalIF if err != nil log error and return true
func FatalIF(err error) bool {
	if err != nil {
		suger.Fatal(err)
		return true
	}

	return false
}

// FatalmIF if err != nil log message and error and return true
func FatalmIF(msg string, err error) bool {
	if err != nil {
		suger.Fatalw(msg, "error", err)
		return true
	}

	return false
}
