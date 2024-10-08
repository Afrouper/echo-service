package main

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	Http          Http                `json:"http,omitempty"`
	Header        map[string][]string `json:"header,omitempty"`
	RemoteAddress string              `json:"remoteAddress,omitempty"`
	Request       Request             `json:"request,omitempty"`
	Authorization JwtPayload          `json:"authorization,omitempty"`
	InstanceName  string              `json:"instanceName,omitempty"`
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

type JwtPayload struct {
	Header  map[string]interface{} `json:"header"`
	Payload interface{}            `json:"payload"`
}

var InstanceName string

func StartEchoService(port int, instanceName string) error {
	log.Printf("Starting echo-service at port %v and instance name %v\n", port, instanceName)
	InstanceName = instanceName
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

	response.InstanceName = InstanceName

	authHeader := req.Header.Get("Authorization")
	if authHeader != "" {
		handleJWT(authHeader, &response.Authorization)
	}
	handleBody(req, &response)

	handleTimeout(req)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Unable to send or decode JSON: %v", err)
	}
}

func handleTimeout(req *http.Request) {
	timeoutValue := req.URL.Query().Get("timeout")
	if timeoutValue == "" {
		timeoutValue = req.Header.Get("X-Timeout")
	}

	if timeoutValue != "" {
		duration, err := time.ParseDuration(timeoutValue)
		if err == nil {
			log.Printf("Cannot parse duration. Check format! %v", err)
		}
		log.Printf("Sleep for %v.", duration)
		time.Sleep(duration)
	}
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

func handleJWT(tokenString string, jwtPayload *JwtPayload) {
	tokenString, _ = strings.CutPrefix(tokenString, "Bearer ")
	token, _, err := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		log.Printf("Error parsing JWT Token: %v", err)
		return
	}
	jwtPayload.Header = token.Header
	jwtPayload.Payload = token.Claims
}
