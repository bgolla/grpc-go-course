package main

import (
	"fmt"
	"log"

	"github.com/bgolla/grpc_go_course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello world")
	lis, err := net.listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceSErver(s, &server{})
}
