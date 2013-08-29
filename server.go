package main

import (
	"github.com/edmore/esp/service"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"flag"
	"os"
)

var	port = flag.String("port", "", "server listening port; must be set.")

func main() {
	flag.Parse()
	if *port == "" {
		flag.Usage()
		os.Exit(2)
	}

	log.Println("Server starting up ...")
	rpc.Register(new(service.Arith))
	log.Println("Arith Service Registered.")
	rpc.Register(new(service.Ping))
	log.Println("Ping Service Registered.")
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":" + *port)
	log.Println("Server listening on", l.Addr())
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	select {}
}