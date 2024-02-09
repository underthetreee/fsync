package main

import (
	"log"
	"net"

	"github.com/underthetreee/fsync/internal/sync"
	"google.golang.org/grpc"
)

func main() {
	gRPCServer := grpc.NewServer()
	sync.Register(gRPCServer)
	l, _ := net.Listen("tcp", ":50051")

	log.Println("gRPC server is listening on", l.Addr())
	if err := gRPCServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}
