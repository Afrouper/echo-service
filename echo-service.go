package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	Http          Http                `json:"http,omitempty"`
	Header        map[string][]string `json:"header,omitempty"`
	RemoteAddress string              `json:"remoteAddress,omitempty"`
	Request       Request             `json:"request,omitempty"`
}

type Http struct {
	Method        string `json:"method,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
	Host          string `json:"host,omitempty"`
	ContentLength int64  `json:"contentLength,omitempty"`
}

type Request struct {
	RequestURI  string      `json:"requestURI,omitempty"`
	Path        string      `json:"path,omitempty"`
	QueryString string      `json:"queryString,omitempty"`
	Body        interface{} `json:"body,omitempty"`
}

func StartEchoService(port int) error {
	log.Printf("Starting echo-service at port %v\n", port)
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
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

	response.Request.RequestURI = req.RequestURI
	response.Request.QueryString = req.URL.RawQuery
	response.Request.Path = req.URL.Path

	response.RemoteAddress = req.RemoteAddr

	response.Header = req.Header

	handleBody(req, &response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleBody(req *http.Request, res *Response) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Cannot create request body: %v", err)
		res.Request.Body = "N/A"
		return
	}

	contentType := req.Header.Get("Content-Type")
	if strings.EqualFold("application/json", contentType) {
		err = json.Unmarshal(b, &res.Request.Body)
		if err != nil {
			log.Printf("Cannot create json for request body: %v", err)
			res.Request.Body = "N/A"
		}
	} else {
		res.Request.Body = string(b)
	}
}
