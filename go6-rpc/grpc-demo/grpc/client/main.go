package main

import (
	"client/protocol"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	dial, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := protocol.NewHelloServiceClient(dial)
	resp, err := client.Hello(context.Background(), &protocol.Request{Value: "jadejones"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
