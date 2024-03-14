FROM golang:latest

WORKDIR /app/server

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080 8080