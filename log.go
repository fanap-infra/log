package log

// Debug simple
func Debug(args ...interface{}) {
	suger.Debug(args)
}

// Debugf format text with value
func Debugf(template string, args ...interface{}) {
	suger.Debugf(template, args...)
}

// Debugv message with key=value ...
func Debugv(msg string, keysAndValues ...interface{}) {
	suger.Debugw(msg, keysAndValues...)
}

// Debugv message with caller info and key=value ...
func Debugc(msg string, keysAndValues ...interface{}) {
	sugerCaller.Debugw(msg, keysAndValues...)
}

// Info simple
func Info(args ...interface{}) {
	suger.Info(args...)
}

// Infof format text with value
func Infof(template string, args ...interface{}) {
	suger.Infof(template, args...)
}

func Infov(msg string, keysAndValues ...interface{}) {
	suger.Infow(msg, keysAndValues...)
}

func Infoc(msg string, keysAndValues ...interface{}) {
	sugerCaller.Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	suger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	suger.Warnf(template, args...)
}

func Warnv(msg string, keysAndValues ...interface{}) {
	suger.Warnw(msg, keysAndValues...)
}

func Warnc(msg string, keysAndValues ...interface{}) {
	sugerCaller.Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	suger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	suger.Errorf(template, args...)
}

func Errorv(msg string, keysAndValues ...interface{}) {
	suger.Errorw(msg, keysAndValues...)
}

func Errorc(msg string, keysAndValues ...interface{}) {
	sugerCaller.Errorw(msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	suger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	suger.Fatalf(template, args...)
}

func Fatalv(msg string, keysAndValues ...interface{}) {
	suger.Fatalw(msg, keysAndValues...)
}

func Fatalc(msg string, keysAndValues ...interface{}) {
	sugerCaller.Fatalw(msg, keysAndValues...)
}
