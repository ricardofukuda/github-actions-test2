FROM golang:1.26 AS builder

WORKDIR /app
COPY . .
RUN go build -o app .

# multi stage docker
FROM ubuntu:24.04
WORKDIR /app
COPY --chown=1000:1000 --chmod=700 --from=builder /app/app .

USER 1000:1000

ENTRYPOINT ["./app"]