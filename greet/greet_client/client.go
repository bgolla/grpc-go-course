package main

import "fmt"

import "google.golang.org/grpc"

import "log"

import "github.com/bgolla/grpc_go_course/greet/greetpb"

func main() {
	fmt.Println("Hello I am client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldnt connect : %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %v", c)

}
