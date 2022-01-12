package main

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/grpc"
	"grpcstream/protobuf"
)

func RunClient() {
	dial, err := grpc.Dial(
		"127.0.0.1:9000",
		grpc.WithDefaultCallOptions(),
		grpc.WithInsecure(),
	)

	client := protobuf.NewGetServiceClient(dial)

	stream, err := client.Get(context.Background(), &protobuf.Request{Count: 10})
	if err != nil {
		panic(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(fmt.Sprintf("reveived error: %s", err.Error()))
		}
		fmt.Println(fmt.Sprintf("received response: %d", res.GetId()))
	}
}
