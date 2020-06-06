package log

// Printf error with format string
func Printf(format string, args ...interface{}) {
	suger.Errorf(format, args...)
}
