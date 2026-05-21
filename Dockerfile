FROM golang:1.26 AS builder

WORKDIR /app
COPY app.go go.mod .
RUN go build -o app .

# multi stage docker
FROM ubuntu:24.04
WORKDIR /app
COPY --chown=1000:1000 --chmod=700 --from=builder /app/app .

USER 1000:1000

ENTRYPOINT ./app