package main

import (
	"fmt"
	"github.com/edmore/esp/service"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Arith Service
	args := &service.Args{7, 8}
	var reply int
	err = client.Call("Arith.Mul", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

	// Ping Service
	var pong, ping string
	fmt.Printf("\nPlease enter your ping string :\n")
	fmt.Scanf("%s", &ping)

	err = client.Call("Ping.Pong", ping, &pong)
	if err != nil {
		log.Fatal("ping error:", err)
	}
	fmt.Printf(pong)
}
