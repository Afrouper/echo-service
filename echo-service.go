package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Http Http `json:"http"`
}

type Http struct {
	Method string `json:"method"`
}

func main() {
	port, found := os.LookupEnv("echo-service.port")
	if !found {
		port = "8080"
	}

	log.Printf("Starting echo-service at port %v\n", port)
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Unable to start http server on port %v: %v", port, err)
	}
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	response := Response{}
	response.Http.Method = req.Method

	json.NewEncoder(w).Encode(response)
}
