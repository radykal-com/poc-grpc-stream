package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting grpc Server")
	go func() {
		RunServer()
	}()

	time.Sleep(3 * time.Second)

	fmt.Println("running client requests")
	RunClient()
	fmt.Println("done")
}
