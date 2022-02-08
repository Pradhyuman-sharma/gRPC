package main

import (
	"context"
	"fmt"
	pb "grpc/gen/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	client := pb.NewTestApiClient(conn)
	res, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Hello everyone!"})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)
	fmt.Println(res.Msg)
}
