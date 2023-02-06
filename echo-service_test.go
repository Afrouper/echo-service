package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEchoService(t *testing.T) {
	t.Logf("Start test...")

	r := httptest.NewRequest(http.MethodPost, "https://funkyTestServer.org/foo/bar?query1=key1", nil)
	w := httptest.NewRecorder()
	handleRequest(w, r)
	res := w.Result()

	contentType := res.Header.Get("Content-Type")
	if !strings.EqualFold(contentType, "application/json") {
		t.Errorf("Invalid Contant-Type: %v", contentType)
	}

	var responseJson map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&responseJson)
	if err != nil {
		t.Error(err)
	}

	path := responseJson["path"].(string)
	if !strings.EqualFold(path, "/foo/bar") {
		t.Fail()
	}
}
