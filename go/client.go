package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-example/go/example"
	
)

func main() {
	conn, err := grpc.Dial(":50050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c:= example.NewExampleServiceClient(conn)

	responseHello, err := c.SayHello(context.Background(), &example.HelloRequest{Greeting: "Prakash"})

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", responseHello.Reply)

	responseAddition, err := c.DoAddition(context.Background(), &example.NumberRequest{FirstNumber: 1, SecondNumber: 2})

	if err != nil {
		log.Fatalf("Error when calling DoAddition: %s", err)
	}

	log.Printf("Response from server: %d", responseAddition.Result)

	responseSubtraction, err := c.DoSubtraction(context.Background(), &example.NumberRequest{FirstNumber: 1, SecondNumber: 2})

	if err != nil {
		log.Fatalf("Error when calling DoSubtraction: %s", err)
	}

	log.Printf("Response from server: %d", responseSubtraction.Result)
}
