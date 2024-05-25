package main

import (
	"bufio"
	"fmt"
	"os"
	"log"

	"lspgo/rpc"
)

func main() {
	fmt.Println("Starting...")
	logger := getLogger("D:\\Dev\\lspgo\\log.txt")
	logger.Println("logger started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg)
	}
}

func handleMessage(_ any) {}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	return log.New(logfile, "[lspgo]", log.Ldate|log.Ltime|log.Lshortfile)
}
