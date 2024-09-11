package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	port, found := os.LookupEnv("echo_service_port")
	if !found {
		port = "8080"
	}
	instanceName, found := os.LookupEnv("instance_name")
	if !found {
		instanceName = generateRandomString(10)
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Cannot start echo-service: %v", err)
	}
	err = StartEchoService(p, instanceName)
	if err != nil {
		log.Fatalf("Cannot start echo-service: %v", err)
	}
}
