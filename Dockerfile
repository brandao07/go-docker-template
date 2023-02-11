FROM golang:latest

WORKDIR /usr/src/app

# Init
RUN go mod init github.com/brandao07/go-docker-template
# Fiber
RUN go get github.com/gofiber/fiber/v2
# Hot-Reload
RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy

