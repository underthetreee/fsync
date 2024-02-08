.SILENT:

.PHONY: run 
run:
	go build -o bin/fsync cmd/main.go
	./bin/fsync