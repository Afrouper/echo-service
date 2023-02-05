package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestEchoService(t *testing.T) {
	fmt.Printf("Start test...")

	res, err := http.Get("http://localhost:8080/foo/bar?query1=value1")
	if err != nil {
		t.Errorf("Could not make HTTP request: %v", err)
	}
	contentType := res.Header.Get("Content-Type")
	if !strings.EqualFold(contentType, "application/json") {
		t.Errorf("Invalid Contant-Type: %v", contentType)
	}

	var responseJson map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&responseJson)
	if err != nil {
		t.Error(err)
	}

	path := responseJson["path"].(string)
	if !strings.EqualFold(path, "/foo/bar") {
		t.Fail()
	}
}
