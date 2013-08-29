package main

import (
	"flag"
	"fmt"
	"github.com/edmore/esp/service"
	"log"
	"net/rpc"
	"os"
)

var (
	ip   = flag.String("ip", "", "server IP address; must be set.")
	port = flag.String("port", "", "server port; must be set.")
)

func main() {
	flag.Parse()
	if *ip == "" || *port == "" {
		flag.Usage()
		os.Exit(2)
	}

	client, err := rpc.DialHTTP("tcp", *ip+":"+*port)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var pong, ping string

	// Arith Service
	args := &service.Args{7, 8}
	var reply int
	err = client.Call("Arith.Mul", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

	// Ping Service
	for ping != "exit" {
		fmt.Printf("\nPlease enter your ping string :\n")
		fmt.Scanf("%s", &ping)

		err = client.Call("Ping.Pong", ping, &pong)
		if err != nil {
			log.Fatal("ping error:", err)
		}
		fmt.Printf(pong)
	}
}
