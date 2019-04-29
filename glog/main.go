package main

import (
	"flag"
	"github.com/golang/glog"
)

func main() {
	fatal := flag.Bool("fatal", false, "[option] fatal")
	flag.Parse()

	if *fatal {
		glog.Fatal("Fatal")
	}
	glog.Error("Error")
	glog.Warning("Warning")
	glog.Info("Info")
}
