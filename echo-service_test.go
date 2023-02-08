package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEchoServiceGet(t *testing.T) {
	responseJson, err := doTestRequest(http.MethodGet, "https://funkyTestServer.org/foo/bar?query1=key1", nil, "")
	if err != nil {
		t.Fatalf("Cannot create test request: %v", err)
	}

	httpJso := responseJson["http"].(map[string]interface{})
	requestJso := responseJson["request"].(map[string]interface{})

	if !strings.EqualFold(requestJso["path"].(string), "/foo/bar") {
		t.Fail()
	}

	if !strings.EqualFold(httpJso["method"].(string), "GET") {
		t.Fail()
	}
}

func TestEchoServicePostWithJson(t *testing.T) {
	type inJson struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	in := inJson{42, "Lutz"}
	b, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("Cannot create json: %v", err)
	}
	responseJson, err := doTestRequest(http.MethodPost, "https://funkyTestServer.org/foo/bar?query1=key1", bytes.NewBuffer(b), "application/json")
	if err != nil {
		t.Fatalf("Cannot create test request: %v", err)
	}

	requestJso := responseJson["request"].(map[string]interface{})
	body := requestJso["body"].(map[string]interface{})
	if !strings.EqualFold("Lutz", body["name"].(string)) {
		t.Error("Name not expected")
	}
	if 42 != body["id"].(float64) {
		t.Error("Id not expected")
	}
}

func TestEchoServicePostWithText(t *testing.T) {
	responseJson, err := doTestRequest(http.MethodPost, "https://funkyTestServer.org/foo/bar?query1=key1", strings.NewReader("Hello Echo"), "plain/text")
	if err != nil {
		t.Fatalf("Cannot create test request: %v", err)
	}

	requestJso := responseJson["request"].(map[string]interface{})
	if !strings.EqualFold("Hello Echo", requestJso["body"].(string)) {
		t.Error("Wrong body")
	}
}

func doTestRequest(method, target string, body io.Reader, ct string) (map[string]interface{}, error) {
	r := httptest.NewRequest(method, target, body)
	w := httptest.NewRecorder()
	if ct != "" {
		r.Header.Set("Content-Type", ct)
	}
	handleRequest(w, r)
	res := w.Result()

	contentType := res.Header.Get("Content-Type")
	if !strings.EqualFold(contentType, "application/json") {
		return nil, fmt.Errorf("invalid Contant-Type: %v", contentType)
	}

	var responseJson map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&responseJson)
	if err != nil {
		return nil, err
	}
	return responseJson, nil
}
