package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	port, found := os.LookupEnv("echo-service.port")
	if !found {
		port = "8080"
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Cannot start echo-service: %v", err)
	}
	StartEchoService(p)
}
