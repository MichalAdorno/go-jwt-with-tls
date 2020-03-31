FROM golang:1.14.1-alpine3.11 AS builder
RUN mkdir /source
COPY ["./","/source/"]
RUN cd /source && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o main . && \
    cp /source/main /source/build
FROM alpine:3.11.5 AS production
RUN mkdir /target
COPY --from=builder ["/source/build/main","/source/config.yaml","/source/server.crt","/source/server.key","/target/"]
RUN chmod a+x /target/main
EXPOSE 8080
ENTRYPOINT ["/target/main"]
CMD ["-config=/target/config.yaml"]
