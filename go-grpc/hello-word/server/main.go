package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "myself/go-grpc/hello-word/helloworld"
	"net"
)

const port = ":50051"

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Name: "Hello " + in.GetName()}, nil
}

type ControlBehavior int32

const (
	Reject ControlBehavior = iota
	Throttling
)

func (s ControlBehavior) String() string {
	switch s {
	case Reject:
		return "Reject"
	case Throttling:
		return "Throttling"
	default:
		return "Undefined"
	}
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	p := []People{
		{
			Age:  1,
			Name: "1",
		},
		{
			Age:  2,
			Name: "2",
		},
	}
	fmt.Println(p)
}

type People struct {
	Age  int      `json:"age"`
	Name string   `json:"name"`
	S    []string `json:"s"`
}

func (p *People) String() string {
	s, _ := json.Marshal(*p)
	return string(s)
}
