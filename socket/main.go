package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
)

var (
	argsTcpEndpoint = flag.String("t", "" /*    */, "[req] tcp end point ( e.g. 172.217.175.3:8080 )")
	argsPayload     = flag.String("b", "" /*    */, "[req] payload ( e.g. GET / HTTP/1.0\\r\\n )")
	argsHelp        = flag.Bool("h", false /*   */, "\nhelp")
	argsDebug       = flag.Bool("d", false /*   */, "\ndebug")
)

// Usage:
// go run main.go -t 172.217.175.46:80 -b "GET / HTTP/1.1\r\nHost: 172.217.175.46\r\nUser-Agent: golangscript\r\nAccept: */*\r\nConnection: Close\r\n\r\n"
//
// > Socketプログラミング · Build web application with Golang
// > https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/08.1.html
func main() {

	flag.Parse()
	if *argsHelp || *argsTcpEndpoint == "" || *argsPayload == "" {
		flag.Usage()
		os.Exit(0)
	}
	requestData := strings.Replace(*argsPayload, "\\n", "\n", -1)
	requestData = strings.Replace(requestData, "\\r", "\r", -1)

	// Create TcpAddress
	tcpAddress, err := net.ResolveTCPAddr("tcp4", *argsTcpEndpoint)
	handleError(err, *argsTcpEndpoint+" is unexpected TCP end point format.")
	log.Println(tcpAddress)

	// Establish tcp connection
	tcpConnection, err := net.DialTCP("tcp", nil, tcpAddress)
	handleError(err, "Tcp connection could not establish.")
	log.Println(tcpConnection)

	// Make kill signal channel
	deferFunc, goroutineDeferFunc := createDeferFunc(tcpConnection, make(chan os.Signal, 1))
	defer deferFunc()
	go goroutineDeferFunc()

	// Write data
	log.Printf("Write data : \n%v", requestData)
	_, err = tcpConnection.Write([]byte(requestData))
	handleError(err, "Error occurred in tcpConnection.Write()")

	// Read data
	result, err := ioutil.ReadAll(tcpConnection)
	handleError(err, "Error occurred in ioutil.ReadAll(tcpConnection)")
	log.Printf("Read data : \n%v", string(result))
	os.Exit(0)

}

func handleError(err error, errorMessage string) {
	if err != nil {
		log.Fatalf("%v ( err: %v )\n", errorMessage, err)
	}
}

func createDeferFunc(tcpConnection *net.TCPConn, signalChannel chan os.Signal) (func(), func()) {
	deferFunc := func() {
		log.Println("Call defer function")
		tcpConnection.Close()
	}

	// Make kill signal channel
	signal.Notify(signalChannel, os.Kill, os.Interrupt)
	return deferFunc, func() {
		<-signalChannel
		log.Println("Catch signals")
		deferFunc()
		os.Exit(1)
	}
}
