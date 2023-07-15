# Go gRPC PING PONG

## Install compiler

```bash
go get -u google.golang.org/protobuf
go get -u google.golang.org/protobuf/proto
```

## Install protocol compiler plugins for Go

> https://grpc.io/docs/languages/go/quickstart/#prerequisites

## Run Complier

```bash
$ protoc --go_out=. --go-grpc_out=. ./proto/*.proto
```

## Run Server

```bash
$ go run server/main.go
```
