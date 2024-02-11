package main

import (
	"log"

	manager "github.com/underthetreee/fsync/pkg/file_manager"
)

func main() {
	m, err := manager.NewManager()
	if err != nil {
		log.Fatal(err)
	}

	// gRPCServer := grpc.NewServer()
	// sync.Register(gRPCServer)
	// l, _ := net.Listen("tcp", ":50051")

	// log.Println("gRPC server is listening on", l.Addr())
	// if err := gRPCServer.Serve(l); err != nil {
	// 	log.Fatal(err)
	// }
}
