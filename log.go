package main

import (
	"fmt"
	"log"
	"os"
)

func initLog() {

	log.Println("[initLog] changing log destination to ... ")
	relativePathLog := fmt.Sprintf("./%s", LOG_FILE_PATH_RELATIVE)
	log.Printf("[initLog]									 ... %s", relativePathLog)

	file, err := os.OpenFile(relativePathLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println("[initLog] END")
}
