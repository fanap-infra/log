package log

// Error args
func Error(args ...interface{}) {
	suger.Error(args...)
}

// Errorf format message
func Errorf(format string, args ...interface{}) {
	suger.Errorf(format, args...)
}

// Errorv message with key=value ...
func Errorv(message string, keysValues ...interface{}) {
	suger.Errorw(message, keysValues...)
}

// ErrorIF if err != nil log message and error and return true
func ErrorIF(msg string, err error) bool {
	if err != nil {
		suger.Errorw(msg, "error", err)
		return true
	}

	return false
}

// ErrorIF if err != nil log error and return true
// func ErrorIF(err error) bool {
// 	if err != nil {
// 		suger.Error(err)
// 		return true
// 	}

// 	return false
// }
