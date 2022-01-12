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

	fmt.Println("-----------------------------------------")
	fmt.Println("running stream client requests")
	RunStreamClient()
	fmt.Println("stream client done")
	fmt.Println("-----------------------------------------")
	fmt.Println("-----------------------------------------")
	fmt.Println("running get client requests")
	RunClient()
	fmt.Println("get client done")
	fmt.Println("-----------------------------------------")
}
