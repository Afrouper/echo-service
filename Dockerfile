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

FROM scratch

LABEL org.opencontainers.image.description = "Simple http echo-service"

COPY --from=BUILDER /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=BUILDER /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=BUILDER /etc/passwd /etc/passwd
COPY --from=BUILDER /etc/group /etc/group

USER appuser:appuser

ENTRYPOINT ["/echo-server"]