FROM alpine AS BUILDER

RUN apk update && \
    apk add --no-cache git ca-certificates tzdata && \
    update-ca-certificates

ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

#WORKDIR $GOPATH/src/github.com/Afrouper/echo-service
#COPY . .

#RUN go version &&\
#    go get -d -v &&\
#    CGO_ENABLED=0 GOOS=linux go build \
#    -ldflags='-w -s -extldflags "-static"' -a \
#    -o /go/bin/echo-server . &&\
#    go test -v


FROM scratch

#LABEL org.opencontainers.image.description = "Simple http echo-service"

COPY --from=BUILDER /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=BUILDER /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=BUILDER /etc/passwd /etc/passwd
COPY --from=BUILDER /etc/group /etc/group

#COPY --from=BUILDER /go/bin/echo-server /echo-server
#COPY dist/echo-service_linux_amd64_v1 /echo-server

USER appuser:appuser

ENTRYPOINT ["/echo-server"]