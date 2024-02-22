.SILENT:

.PHONY: server 
server:
	go build -o bin/server cmd/server/main.go
	./bin/server

.PHONY: client 
client:
	go build -o bin/client cmd/client/main.go
	./bin/client

.PHONY: proto 
proto:
	mkdir -p pkg
	protoc --go_out=./pkg --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative proto/filesync.proto

.PHONY: compose-up 
compose-up:
	docker compose up -d

.PHONY: compose-down
compose-down:
	docker compose down --remove-orphans