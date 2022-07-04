/*
 * @Autor: 郭彬
 * @Description:client
 * @Date: 2021-10-18 14:07:38
 * @LastEditTime: 2021-10-18 14:57:28
 * @FilePath: \gotest\src\grpc-helloworld\greeter_client\main.go
 */

package main

import (
	"log"
	"os"

	"gotest/src/grpc-helloworld/helloworld"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
