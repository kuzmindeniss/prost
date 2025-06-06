FROM golang:latest
WORKDIR /app
ENTRYPOINT ["go", "run", "cmd/tg/main.go"]
