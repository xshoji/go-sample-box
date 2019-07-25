package main

import (
	"flag"
	"fmt"
	"github.com/apex/log"
	"io/ioutil"
	"os"
	"path/filepath"
)

var helpFlag = flag.Bool("help", false, "help")
var directoryPathFlag = flag.String("dir", "", "[required] dir")

func init() {
	flag.BoolVar(helpFlag, "h", false, "= -help")
	flag.StringVar(directoryPathFlag, "d", "", "= -dir")
}

func main() {

	flag.Parse()
	if *helpFlag || *directoryPathFlag == "" {
		flag.Usage()
		os.Exit(0)
	}

	files, err := ioutil.ReadDir(*directoryPathFlag)
	if err != nil {
		log.Fatal("[ err ] " + *directoryPathFlag + " not exists.")
		os.Exit(1)
	}
	fmt.Println("dir: ", *directoryPathFlag)

	for _, file := range files {
		fileFullPath := filepath.Join(*directoryPathFlag, file.Name())
		if file.IsDir() {
			fmt.Println("[ dir ] " + file.Name() + " (" + fileFullPath + ")")
		} else {
			fmt.Println("[ file ] " + file.Name() + " (" + fileFullPath + ")")
		}
	}
}
