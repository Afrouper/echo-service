# Simple echo service

[![Build and Test](https://github.com/afrouper/echo-service/actions/workflows/build.yml/badge.svg)](https://github.com/afrouper/echo-service/actions/workflows/build.yml)
[![Build Release](https://github.com/afrouper/echo-service/actions/workflows/release.yaml/badge.svg)](https://github.com/afrouper/echo-service/actions/workflows/release.yaml)
[![CodeQL](https://github.com/afrouper/echo-service/actions/workflows/codeql.yml/badge.svg)](https://github.com/afrouper/echo-service/actions/workflows/codeql.yml)
[![](https://badgen.net/github/release/afrouper/echo-service?icon=github)](https://github.com/afrouper/echo-service/releases/latest)
![](https://badgen.net/github/releases/afrouper/echo-service)
![](https://badgen.net/github/open-issues/afrouper/echo-service)
![](https://badgen.net/docker/size/afrouper/echo-service/latest/amd64)
[![](https://badgen.net/github/license/afrouper/echo-service)](https://raw.githubusercontent.com/afrouper/echo-service/main/LICENSE)

Plain and very simple echo service usefully for testing infrastructure components.

A docker image is available on [Docker Hub](https://hub.docker.com/r/afrouper/echo-service) and
[ghcr.io](https://github.com/Afrouper/echo-service/pkgs/container/echo-service).

## Build & Test
Install the latest go version: https://go.dev/doc/install

Compilation can be done via `go build -v ./...`, Tests via `go test -v ./...`

## Usage
The binary can be started via `./echo-service`. By default, it listens on port 8080. You can change the listen port
by setting the environment variable `echo_service_port`

If using a JWT Bearer token is send the token will be decoded (without verifying!) and send back as JSON.

If called with a Chrome browser this is  the output: 
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