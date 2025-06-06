FROM golang:latest
WORKDIR /app
ENTRYPOINT ["go", "run", "cmd/bff/main.go"]
