package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"lspgo/lsp"
	"lspgo/rpc"
)

func main() {
	fmt.Println("Starting...")
	logger := getLogger("D:\\Dev\\lspgo\\log.txt")
	logger.Println("logger started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}
		handleMessage(logger, method, content)
	}
}

func handleMessage(logger *log.Logger, method string, content []byte) {
	logger.Printf("Received message with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("Couldn't parse content: %s", err)
		}
		logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	return log.New(logfile, "[lspgo]", log.Ldate|log.Ltime|log.Lshortfile)
}
