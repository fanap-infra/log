package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level zapcore.Level

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = Level(zapcore.DebugLevel)
	// InfoLevel is the default logging priority.
	InfoLevel = Level(zapcore.InfoLevel)
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = Level(zapcore.WarnLevel)
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = Level(zapcore.ErrorLevel)
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel = Level(zapcore.DPanicLevel)
	// PanicLevel logs a message, then panics.
	PanicLevel = Level(zapcore.PanicLevel)
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = Level(zapcore.FatalLevel)
)

var zapLogger *zap.Logger
var suger *zap.SugaredLogger
var sugerCaller *zap.SugaredLogger
var loggerLevel zapcore.Level = zapcore.WarnLevel

func init() {
	// WarnLevel is default
	defaultInit()
}

// Init Logger
func defaultInit() {
	// The bundled Config struct only supports the most common configuration
	// options. More complex needs, like splitting logs between multiple files
	// or writing to non-file outputs, require use of the zapcore package.
	//
	// In this example, imagine we're both sending our logs to Kafka and writing
	// them to the console. We'd like to encode the console output and the Kafka
	// topics differently, and we'd also like special treatment for
	// high-priority logs.

	// First, define our level-handling logic.
	levelEnabler := zap.LevelEnablerFunc(levelEnablerFunc)

	//highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	//	return lvl >= level
	//})
	//lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	//	return lvl < level
	//})

	// Assume that we have clients for two Kafka topics. The clients implement
	// zapcore.WriteSyncer and are safe for concurrent use. (If they only
	// implement io.Writer, we can use zapcore.AddSync to add a no-op Sync
	// method. If they're not safe for concurrent use, we can add a protecting
	// mutex with zapcore.Lock.)
	// topicDebugging := zapcore.AddSync(ioutil.Discard)
	// topicErrors := zapcore.AddSync(ioutil.Discard)

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	//consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.
	// kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	// consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	config := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// config := zapcore.EncoderConfig{
	// 	// Keys can be anything except the empty string.
	// 	TimeKey:        "T",
	// 	LevelKey:       "L",
	// 	NameKey:        "N",
	// 	CallerKey:      "C",
	// 	MessageKey:     "M",
	// 	StacktraceKey:  "S",
	// 	LineEnding:     zapcore.DefaultLineEnding,
	// 	EncodeLevel:    zapcore.CapitalLevelEncoder,
	// 	EncodeTime:     zapcore.ISO8601TimeEncoder,
	// 	EncodeDuration: zapcore.StringDurationEncoder,
	// 	EncodeCaller:   zapcore.ShortCallerEncoder,
	// }

	consoleEncoder := zapcore.NewConsoleEncoder(config)

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		// zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, levelEnabler),
		// zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),
		//zapcore.NewCore(consoleEncoder, consoleDebugging, levelEnabler),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	zapLogger = zap.New(core)

	suger = zapLogger.Sugar()
	sugerCaller = zapLogger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

// Sync calls the underlying Core's Sync method, flushing any buffered log
func Sync() {
	zapLogger.Sync()
}

func levelEnablerFunc(level zapcore.Level) bool {
	return level >= loggerLevel
}

// SetLevel of log will been enable
func SetLevel(level Level) {
	loggerLevel = zapcore.Level(level)
}

// RedirectStdLog std log to this to Info Level
// It returns a function to restore the original prefix and flags and reset the
// standard library's output to os.Stderr.
func RedirectStdLog() func() {
	return zap.RedirectStdLog(zapLogger)
}

// RedirectStdLogAt std log to this at log level
// It returns a function to restore the original prefix and flags and reset the
// standard library's output to os.Stderr.
func RedirectStdLogAt(level Level) func() {
	f, _ := zap.RedirectStdLogAt(zapLogger, zapcore.Level(level))
	return f
}
