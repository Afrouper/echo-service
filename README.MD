# Simple echo service

[![Current Release](https://badgen.net/github/release/Afrouper/echo-service?icon=github)](https://github.com/Afrouper/echo-service/releases/latest)
![License](https://badgen.net/github/license/Afrouper/echo-service)

Plain and very simple echo service usefully for testing infrastructure  components.

A docker image is avalable on [Docker Hub](https://hub.docker.com/repository/docker/afrouper/echo-service)

## Build & Test
Install the latest go version: https://go.dev/doc/install

Compilation can be done via `go build -v ./...`, Tests via `go test -v ./...`

## Usage
The binary can be started via `./echo-service`. It listens to port 8080. If called with a Chrome browser this is
the output: 
```json
{
  "http": {
    "method": "GET",
    "protocol": "HTTP/1.1",
    "host": "localhost:8080"
  },
  "header": {
    "Accept": [
      "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
    ],
    "Accept-Encoding": [
      "gzip, deflate, br"
    ],
    "Accept-Language": [
      "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7"
    ],
    "Connection": [
      "keep-alive"
    ],
    "Dnt": [
      "1"
    ],
    "Sec-Ch-Ua": [
      "\"Not_A Brand\";v=\"99\", \"Google Chrome\";v=\"109\", \"Chromium\";v=\"109\""
    ],
    "Sec-Ch-Ua-Mobile": [
      "?0"
    ],
    "Sec-Ch-Ua-Platform": [
      "\"macOS\""
    ],
    "Sec-Fetch-Dest": [
      "document"
    ],
    "Sec-Fetch-Mode": [
      "navigate"
    ],
    "Sec-Fetch-Site": [
      "none"
    ],
    "Sec-Fetch-User": [
      "?1"
    ],
    "Upgrade-Insecure-Requests": [
      "1"
    ],
    "User-Agent": [
      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
    ]
  },
  "remoteAddress": "[::1]:60259",
  "request": {
    "requestURI": "/foo/bar?q=v",
    "path": "/foo/bar",
    "queryString": "q=v",
    "body": ""
  }
}
```