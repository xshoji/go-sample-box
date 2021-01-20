package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	// Required parameters
	argumentDirectoryPath = flag.String("d", "" /*    */, "[required] dir")
	// Optional parameters
	argumentHelp = flag.Bool("h", false /*            */, "\nhelp")
)

func main() {

	flag.Parse()
	if *argumentHelp || *argumentDirectoryPath == "" {
		flag.Usage()
		os.Exit(0)
	}

	listUpFiles(*argumentDirectoryPath)
	listUpAllFiles(*argumentDirectoryPath)
}

func listUpFiles(root string) {
	fmt.Println(">> listUpFiles")
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal("[ err ] " + root + " not exists.")
		os.Exit(1)
	}
	fmt.Println("dir: ", root)

	for _, file := range files {
		fileFullPath := filepath.Join(root, file.Name())
		if file.IsDir() {
			fmt.Println("[ dir ] " + file.Name() + " (" + fileFullPath + ")")
		} else {
			fmt.Println("[ file ] " + file.Name() + " (" + fileFullPath + ")")
		}
	}
}

func listUpAllFiles(root string) {
	fmt.Println(">> listUpAllFiles")
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, filePath := range files {
		fmt.Println(filePath)
	}
}
