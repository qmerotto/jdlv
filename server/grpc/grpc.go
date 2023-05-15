package grpc

import (
	"fmt"
	"log"
	"net"
	"os/exec"

	"google.golang.org/grpc"
)

func initServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	//protos.NewGridHandlerClient()
	//pb.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func BuildProtos() {
	//output := "C:\\Users\\Kent\\Sources\\jdlv\\server\\grpc\\entities"
	cmd := exec.Command("protoc", "--go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative server\\grpc\\protos\\message.proto")
	// fmt.Sprintf("-I=%s", ".server/grpc/protos"), fmt.Sprintf("--go_out=./server/grpc/protos"), fmt.Sprintf("./server/grpc/protos/"))
	fmt.Println(fmt.Sprintf("command: %s", cmd.String()))
	err := cmd.Run()
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}
}
