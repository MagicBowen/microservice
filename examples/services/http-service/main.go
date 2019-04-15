package main

import (
	"flag"
	"log"
)

var (
	logFile = flag.String("log", "output.log", "Log file name")
)

func main() {
	flag.Parse()
	// initLogger(logFile)

	err := rpc.initial("localhost:8899")
	if err != nil {
		log.Fatalf("gRPC init failed")
		return
	}
	defer rpc.release()

	initHTTPServer(":8866")
}
