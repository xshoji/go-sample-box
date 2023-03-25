package main

import (
	"flag"
	"github.com/xshoji/go-sample-box/log-global/customlogger"
	"log"
	"os"
)

var (
	// Define boot arguments.
	paramsLogLevel = flag.Int("l", 1 /*      */, "[optional] Log level (1:Panic, 2:Error, 3:Warn, 4:Info, 5:Debug)")
	paramsHelp     = flag.Bool("h", false /* */, "\nhelp")
)

func main() {

	// 引数のパース
	flag.Parse()
	// Required parameter [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *paramsHelp {
		flag.Usage()
		os.Exit(0)
	}
	// go - Correct approach to global logging - Stack Overflow https://stackoverflow.com/questions/18361750/correct-approach-to-global-logging
	// Initialize logger (Executable once only.)
	customlogger.InitializeLogger(customlogger.NewCustomLogger(
		log.New(os.Stdout, "[Logger] ", log.Lshortfile|log.LstdFlags),
		log.New(os.Stderr, "[ErrorLogger] ", log.Lshortfile|log.LstdFlags),
		customlogger.LogLevel(*paramsLogLevel),
	))
	customlogger.InitializeLogger(customlogger.NewCustomLogger(log.New(os.Stdout, "[Aaaa] ", log.Lshortfile), nil, customlogger.LogLevel(*paramsLogLevel)))
	customlogger.InitializeLogger(customlogger.NewCustomLogger(log.New(os.Stdout, "[Bbbb] ", log.LstdFlags), nil, customlogger.LogLevel(*paramsLogLevel)))

	customlogger.Error("Error log", "error", 222, true)
	customlogger.Warn("Warn log", "warn", 333, false)
	customlogger.Info("Info log", "info", 444, false)
	customlogger.Debug("Debug log", "debug", 555, false)
	customlogger.Panic("Panic log", "panic", 111, false)
}
