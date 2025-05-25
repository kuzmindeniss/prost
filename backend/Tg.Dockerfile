FROM golang:latest
RUN go install github.com/air-verse/air@latest
WORKDIR /app
ENTRYPOINT ["air", "-d", "-c", "cmd/tg/.air.toml"]
