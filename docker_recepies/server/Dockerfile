
# Layer to build binary
FROM golang:latest  AS server_builder
LABEL stage=builder
ENV GO111MODULE=on
WORKDIR /go/src/github.com/rounak316/blazetunnel
COPY go.mod .
COPY go.sum .
RUN go mod download
WORKDIR /go/src/github.com/rounak316/blazetunnel
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' .
FROM alpine AS weaviate
RUN apk add ca-certificates
COPY --from=server_builder /go/bin/blazetunnel /bin/blazetunnel

# Layer to execute binary
FROM alpine  AS runner
COPY --from=server_builder /go/bin/blazetunnel /bin/blazetunnel
EXPOSE 2723
EXPOSE 80
ENTRYPOINT [ "/bin/blazetunnel"] 
