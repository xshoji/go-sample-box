package customlogger

import (
	"log"
	"sync"
)

var (
	globalLogger *customLogger
	once         sync.Once
)

type LogLevel int

type customLogger struct {
	logger              *log.Logger
	currentLoggingLevel LogLevel
}

const (
	LogLevelUnknown LogLevel = iota
	LogLevelPanic
	LogLevelError
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug
)

func NewCustomLogger(logger *log.Logger, logLevel LogLevel) *customLogger {
	return &customLogger{
		logger:              logger,
		currentLoggingLevel: logLevel,
	}
}

// InitializeLogger 'Executable once only.'
func InitializeLogger(customLogger *customLogger) {
	once.Do(func() {
		globalLogger = customLogger
	})
}

func Panic(v ...interface{}) {
	logging(LogLevelPanic, v)
	globalLogger.logger.Panic(v)
}

func Error(v ...interface{}) {
	logging(LogLevelError, v)
}

func Warn(v ...interface{}) {
	logging(LogLevelWarn, v)
}

func Info(v ...interface{}) {
	logging(LogLevelInfo, v)
}

func Debug(v ...interface{}) {
	logging(LogLevelDebug, v)
}

// Level based logging in Golang
// https://www.linkedin.com/pulse/level-based-logging-golang-vivek-dasgupta
func logging(loglevel LogLevel, v ...interface{}) {
	if loglevel > globalLogger.currentLoggingLevel {
		return
	}
	level := func() string {
		switch loglevel {
		case LogLevelPanic:
			return "Panic"
		case LogLevelError:
			return "Error"
		case LogLevelWarn:
			return "Warn"
		case LogLevelInfo:
			return "Info"
		case LogLevelDebug:
			return "Debug"
		default:
			return ""
		}
	}()
	globalLogger.logger.Println(append([]interface{}{"[" + level + "]"}, v...)...)
}
