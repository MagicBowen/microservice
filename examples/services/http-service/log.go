package main

import (
	"fmt"
	"log"
	"os"
)

func initLogger(logFile *string) {
	outfile, err := os.OpenFile(*logFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(*outfile, "open failed")
		os.Exit(1)
	}
	log.SetOutput(outfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
