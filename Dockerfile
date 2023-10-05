# syntax=docker/dockerfile:1
FROM golang:1.21.1
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build 

