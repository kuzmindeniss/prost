FROM golang:latest
WORKDIR /app
CMD ["go", "run", "cmd/tg/main.go"]
