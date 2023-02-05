package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Http          Http                `json:"http,omitempty"`
	Header        map[string][]string `json:"header,omitempty"`
	RemoteAddress string              `json:"remoteAddress,omitempty"`
	RequestURI    string              `json:"requestURI,omitempty"`
	Path          string              `json:"path,omitempty"`
	QueryString   string              `json:"queryString,omitempty"`
}

type Http struct {
	Method        string `json:"method,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
	Host          string `json:"host,omitempty"`
	ContentLength int64  `json:"contentLength,omitempty"`
}

func StartEchoService(port string) error {
	log.Printf("Starting echo-service at port %v\n", port)
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Printf("Unable to start http server on port %v: %v\n", port, err)
	}
	return err
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	response := Response{}

	response.Http.Method = req.Method
	response.Http.ContentLength = req.ContentLength
	response.Http.Protocol = req.Proto
	response.Http.Host = req.Host

	response.RemoteAddress = req.RemoteAddr
	response.RequestURI = req.RequestURI
	response.QueryString = req.URL.RawQuery
	response.Path = req.URL.Path

	response.Header = req.Header

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
