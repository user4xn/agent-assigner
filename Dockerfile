# Dockerfile
FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o agent-assigner main.go

EXPOSE 8080

ENTRYPOINT ["./agent-assigner"]
