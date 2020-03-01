
# Layer to build binary
FROM golang:latest  AS server_builder
ENV GO111MODULE=on
WORKDIR /go/src/github.com/rounak316/blazetunnel
COPY go.mod .
COPY go.sum .
RUN go mod download
WORKDIR /go/src/github.com/rounak316/blazetunnel
# COPY . .


