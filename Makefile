.SILENT:

.PHONY: server 
server:
	go build -o bin/server cmd/server/main.go
	./bin/server

.PHONY: proto 
proto:
	mkdir -p pkg
	protoc --go_out=./pkg --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative proto/filesync.proto