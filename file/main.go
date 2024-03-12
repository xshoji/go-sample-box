package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filePath := flag.String("file", "", "[require] A file path")
	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *filePath == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("scanFileContents")
	scanFileContents(filePath)
	fmt.Println("")

	fmt.Println("readStringFileContents")
	readStringFileContents(filePath)
	fmt.Println("")

	fmt.Println("readAllFileContents")
	readAllFileContents(filePath)
	fmt.Println("")

	fmt.Println("ReadFile")
	contents, err := os.ReadFile(*filePath)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v", string(contents))
	fmt.Println("")
}

// - [Go でファイルを1行ずつ読み込む（csv ファイルも） - Qiita](https://qiita.com/ikawaha/items/28186d965780fab5533d)
func scanFileContents(filePath *string) {
	file, err := os.Open(*filePath)
	if err != nil {
		log.Panic(err)
	}
	defer createFileCloseDeferFunc(file)()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func readStringFileContents(filePath *string) {
	file, err := os.Open(*filePath)
	if err != nil {
		log.Panic(err)
	}
	defer createFileCloseDeferFunc(file)()

	reader := bufio.NewReaderSize(file, 1024)
	for line := ""; err == nil; line, err = reader.ReadString('\n') {
		fmt.Print(line)
	}
	if err != io.EOF {
		panic(err)
	}
}

func readAllFileContents(filePath *string) {
	file, err := os.Open(*filePath)
	if err != nil {
		log.Panic(err)
	}
	defer createFileCloseDeferFunc(file)()

	contents, err := io.ReadAll(file)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v", string(contents))
}

func createFileCloseDeferFunc(file *os.File) func() {
	return func() {
		fileCloseErr := file.Close()
		if fileCloseErr != nil {
			log.Panic(fileCloseErr)
		}
	}
}
