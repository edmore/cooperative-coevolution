package main

import (
	"github.com/edmore/esp/service"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	log.Println("Server starting up ...")
	arith := new(service.Arith)
	ping := new(service.Ping)

	rpc.Register(arith)
	log.Println("Arith Service Registered.")
	rpc.Register(ping)
	log.Println("Ping Service Registered.")
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	log.Println("Server listening on", l.Addr())
	if e != nil {
		log.Fatal("listen error:", e)
	}

	http.Serve(l, nil)
}
