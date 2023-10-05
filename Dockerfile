# syntax=docker/dockerfile:1
FROM golang:1.21.1


WORKDIR /app

ENV GOPATH=/

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /netserver

EXPOSE 8080

CMD ["/netserver"]
