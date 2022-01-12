package main

import (
	"math/rand"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpcstream/protobuf"
)

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

func (g *getServiceServer) Get(req *protobuf.Request, stream protobuf.GetService_GetServer) error {
	var i int32
	var wg sync.WaitGroup
	for i < req.GetCount() {
		wg.Add(1)
		go func(c int32) {
			wait := time.Duration(rand.Intn(10)) * time.Second
			time.Sleep(wait)

			_ = stream.Send(&protobuf.Response{Id: c})

			wg.Done()
		}(i)
		i++
	}

	wg.Wait()

	return nil
}
