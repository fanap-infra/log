package log

type Level int

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var levelText = []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
