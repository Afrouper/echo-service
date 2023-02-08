FROM golang:alpine AS builder

WORKDIR $GOPATH/src/github.com/Afrouper/echo-service
COPY *.go .
COPY go.mod .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/echo-service

FROM scratch
COPY --from=builder /go/bin/echo-service /echo-service

EXPOSE 8080
USER 1001

ENTRYPOINT ["/echo-service"]