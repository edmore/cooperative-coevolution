package main

import (
	"fmt"
	"github.com/edmore/esp/service"
	"log"
	"net/rpc"
        "os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "[Server IP Address]")
		os.Exit(1)
	}
	ipAddr := os.Args[1]

	client, err := rpc.DialHTTP("tcp", ipAddr)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var pong, ping string

	for ping != "exit" {
		// Arith Service
		args := &service.Args{7, 8}
		var reply int
		err = client.Call("Arith.Mul", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

		// Ping Service

		fmt.Printf("\nPlease enter your ping string :\n")
		fmt.Scanf("%s", &ping)

		err = client.Call("Ping.Pong", ping, &pong)
		if err != nil {
			log.Fatal("ping error:", err)
		}
		fmt.Printf(pong)
	}
}
