package main

import "os"

func main() {
	port, found := os.LookupEnv("echo-service.port")
	if !found {
		port = "8080"
	}
	StartEchoService(port)
}
