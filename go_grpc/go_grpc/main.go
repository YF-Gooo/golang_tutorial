package main

//client.go

import (
	"fmt"

	pd "grpc_test/rpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = ":50051"
)

func main() {
	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()
	c := pd.NewGreeterClient(conn)

	r, _ := c.SayHello(context.Background(), &pd.HelloRequest{Name: "test"})
	fmt.Println(r.Message)
}
