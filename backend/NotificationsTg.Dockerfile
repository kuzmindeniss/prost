FROM golang:latest
WORKDIR /app
ENTRYPOINT ["go", "run", "cmd/tg_notifications/main.go"]
