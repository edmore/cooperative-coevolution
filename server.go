package main

import (
	"flag"
	"fmt"
	"github.com/edmore/esp/service"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

var port = flag.Int("port", 9999, "default server listening port; can be reset.")

func main() {
	flag.Parse()
	log.Println("Server starting up ...")
	rpc.Register(new(service.Arith))
	log.Println("Arith Service Registered.")
	rpc.Register(new(service.Ping))
	log.Println("Ping Service Registered.")
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Println("Server listening on", l.Addr())
	go http.Serve(l, nil)
	select {}
}
