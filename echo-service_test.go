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
	"time"
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

func TestJWT(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "https://funkyTestServer.org/foo/bar?query1=key1", nil)
	w := httptest.NewRecorder()
	r.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJjdXN0b21DbGFpbSI6IkZvbyJ9.WSt10RZi2AGyQfYsR2AyeH6yUwG89hWsX-cZyVNYMTU")
	handleJwtReply(t, w, r)

	r = httptest.NewRequest(http.MethodGet, "https://funkyTestServer.org/foo/bar?query1=key1", nil)
	w = httptest.NewRecorder()
	r.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJjdXN0b21DbGFpbSI6IkZvbyJ9.WSt10RZi2AGyQfYsR2AyeH6yUwG89hWsX-cZyVNYMTU")
	handleJwtReply(t, w, r)
}

func handleJwtReply(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
	handleRequest(w, r)
	res := w.Result()

	contentType := res.Header.Get("Content-Type")
	if !strings.EqualFold(contentType, "application/json") {
		t.Fatalf("invalid Contant-Type: %v", contentType)
	}

	var responseJson map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&responseJson)
	if err != nil {
		t.Fatalf("JWT error %v", err)
	}

	authorization := responseJson["authorization"].(map[string]interface{})
	if authorization == nil {
		t.Fatalf("Authorization is nil.")
	}
	payload := authorization["payload"].(map[string]interface{})
	header := authorization["header"].(map[string]interface{})

	if payload == nil {
		t.Fatalf("Payload is nil.")
	}
	if header == nil {
		t.Fatalf("Header is nil.")
	}

	if !strings.EqualFold("1234567890", payload["sub"].(string)) {
		t.Fatalf("Claim 'sub' has wrong content %s", payload["sub"].(string))
	}
	if !strings.EqualFold("Foo", payload["customClaim"].(string)) {
		t.Fatalf("Claim 'customClaim' has wrong content %s", payload["customClaim"].(string))
	}

	if !strings.EqualFold("HS256", header["alg"].(string)) {
		t.Fatalf("Claim 'alg' has wrong content %s", header["alg"].(string))
	}
}

func TestTimeout(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "https://funkyTestServer.org/foo/bar?timeout=2s", nil)
	w := httptest.NewRecorder()

	start := time.Now()
	handleRequest(w, r)
	duration := time.Now().Sub(start)

	if duration < 2*time.Second {
		t.Fatalf("Wrong timeout (%v) while calling with Query.", duration)
	}

	r = httptest.NewRequest(http.MethodGet, "https://funkyTestServer.org/foo/bar", nil)
	r.Header.Set("X-Timeout", "2000ms")
	w = httptest.NewRecorder()

	start = time.Now()
	handleRequest(w, r)
	duration = time.Now().Sub(start)

	if duration < 2*time.Second {
		t.Fatalf("Wrong timeout (%v) while calling with Query.", duration)
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
