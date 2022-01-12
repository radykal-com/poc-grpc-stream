package main

import (
	"context"
	"math/rand"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpcstream/protobuf"
)

var waitTimes = []time.Duration{
	time.Duration(rand.Intn(10)) * time.Second,
	time.Duration(rand.Intn(10)) * time.Second,
	time.Duration(rand.Intn(10)) * time.Second,
	time.Duration(rand.Intn(10)) * time.Second,
	time.Duration(rand.Intn(10)) * time.Second,
	time.Duration(rand.Intn(10)) * time.Second,
	time.Duration(rand.Intn(10)) * time.Second,
	time.Duration(rand.Intn(10)) * time.Second,
	time.Duration(rand.Intn(10)) * time.Second,
	time.Duration(rand.Intn(10)) * time.Second,
	time.Duration(rand.Intn(10)) * time.Second,
}

func RunServer() *grpc.Server {
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	protobuf.RegisterGetServiceServer(grpcServer, newGetService())

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

	return grpcServer
}

func newGetService() protobuf.GetServiceServer {
	return &getServiceServer{}
}

type getServiceServer struct {
	protobuf.UnimplementedGetServiceServer
}

func (g *getServiceServer) GetStream(req *protobuf.Request, stream protobuf.GetService_GetStreamServer) error {
	var i int32
	var wg sync.WaitGroup
	for i < 10 {
		wg.Add(1)
		go func(c int32) {
			wait := waitTimes[c]
			time.Sleep(wait)

			_ = stream.Send(&protobuf.Response{Id: []int32{c}})

			wg.Done()
		}(i)
		i++
	}

	wg.Wait()

	return nil
}

func (g *getServiceServer) Get(_ context.Context, req *protobuf.Request) (*protobuf.Response, error) {
	var i int32
	var wg sync.WaitGroup
	for i < 10 {
		wg.Add(1)
		go func(c int32) {
			wait := waitTimes[c]
			time.Sleep(wait)

			wg.Done()
		}(i)
		i++
	}

	wg.Wait()

	return &protobuf.Response{Id: []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}, nil
}
