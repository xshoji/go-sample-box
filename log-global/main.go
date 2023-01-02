package main

import (
	"flag"
	"github.com/xshoji/go-sample-box/sync-once/customlogger"
	"log"
	"os"
)

var (
	// Define boot arguments.
	argsLogLevel = flag.Int("l", 1 /*      */, "[optional] Log level (1:Panic, 2:Error, 3:Warn, 4:Info, 5:Debug)")
	argsHelp     = flag.Bool("h", false /* */, "\nhelp")
)

func main() {

	// 引数のパース
	flag.Parse()
	// Required parameter [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *argsHelp {
		flag.Usage()
		os.Exit(0)
	}
	// Initialize logger (Executable once only.)
	customlogger.InitializeLogger(customlogger.NewCustomLogger(log.New(os.Stdout, "[Logger] ", log.Llongfile|log.LstdFlags), customlogger.LogLevel(*argsLogLevel)))
	customlogger.InitializeLogger(customlogger.NewCustomLogger(log.New(os.Stdout, "[Aaaa] ", log.Llongfile), customlogger.LogLevel(*argsLogLevel)))
	customlogger.InitializeLogger(customlogger.NewCustomLogger(log.New(os.Stdout, "[Bbbb] ", log.LstdFlags), customlogger.LogLevel(*argsLogLevel)))

	customlogger.Error("Error log", "error", 222, true)
	customlogger.Warn("Warn log", "warn", 333, false)
	customlogger.Info("Info log", "info", 444, false)
	customlogger.Debug("Debug log", "debug", 555, false)
	customlogger.Panic("Panic log", "panic", 111, false)
}
