package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	// Define boot arguments.
	paramsLogLevel = flag.Int("l", 4 /*      */, "[optional] Log level (0:Panic, 1:Error, 2:Warn, 3:Info, 4:Debug)")
	paramsHelp     = flag.Bool("h", false /* */, "\nhelp")
	// Define logger: date, time, microseconds, directory and file path are always outputted.
	logger         = log.New(os.Stdout, "[Logger] ", log.Lshortfile|log.LstdFlags)
	loggerLogLevel = Debug
)

type LogLevel int

const (
	Panic LogLevel = iota
	Error
	Warn
	Info
	Debug
)

// Level based logging in Golang
// https://www.linkedin.com/pulse/level-based-logging-golang-vivek-dasgupta
func logging(loglevel LogLevel, logLogger *log.Logger, v ...interface{}) {
	if loglevel > loggerLogLevel {
		return
	}
	level := func() string {
		switch loglevel {
		case Panic:
			return "Panic"
		case Error:
			return "Error"
		case Warn:
			return "Warn"
		case Info:
			return "Info"
		case Debug:
			return "Debug"
		default:
			return ""
		}
	}()
	logLogger.Println(append([]interface{}{"[" + level + "]"}, v...)...)
	if loglevel == Panic {
		logLogger.Panic(fmt.Sprintln(append([]interface{}{"[" + level + "]"}, v...)...))
	}
}

func main() {

	// 引数のパース
	flag.Parse()
	// Required parameter [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *paramsHelp {
		flag.Usage()
		os.Exit(0)
	}

	// set log level
	loggerLogLevel = LogLevel(*paramsLogLevel)

	logging(Error, logger, "Error log", "error", 222, true)
	logging(Warn, logger, "Warn log", "warn", 333, false)
	logging(Info, logger, "Info log", "info", 444, false)
	logging(Debug, logger, "Debug log", "debug", 555, false)
	logging(Panic, logger, "Panic log", "panic", 111, false)
}
