package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc"
	"grpcstream/protobuf"
)

func RunStreamClient() {
	dial, err := grpc.Dial(
		"127.0.0.1:9000",
		grpc.WithDefaultCallOptions(),
		grpc.WithInsecure(),
	)

	client := protobuf.NewGetServiceClient(dial)

	start := time.Now()
	stream, err := client.GetStream(context.Background(), &protobuf.Request{})
	if err != nil {
		panic(err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(fmt.Sprintf("reveived error: %s", err.Error()))
		}

		for _, res := range response.GetId() {
			fmt.Println(fmt.Sprintf("received response: %d: %s elapsed", res, time.Now().Sub(start).String()))
			executeLoadOnResponse()
			fmt.Println(fmt.Sprintf("load executed for response: %d: %s elapsed", res, time.Now().Sub(start).String()))
		}
	}
	fmt.Println(fmt.Sprintf("stream total time %s", time.Now().Sub(start).String()))
}

func RunClient() {
	dial, err := grpc.Dial(
		"127.0.0.1:9000",
		grpc.WithDefaultCallOptions(),
		grpc.WithInsecure(),
	)

	client := protobuf.NewGetServiceClient(dial)

	start := time.Now()
	response, err := client.Get(context.Background(), &protobuf.Request{})
	if err != nil {
		panic(err)
	}

	for _, res := range response.GetId() {
		fmt.Println(fmt.Sprintf("received response: %d: %s elapsed", res, time.Now().Sub(start).String()))
		executeLoadOnResponse()
		fmt.Println(fmt.Sprintf("load executed for response: %d: %s elapsed", res, time.Now().Sub(start).String()))

	}
	fmt.Println(fmt.Sprintf("get total time %s", time.Now().Sub(start).String()))
}

func executeLoadOnResponse() {
	time.Sleep(10 * time.Millisecond)
}
