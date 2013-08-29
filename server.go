package main

import (
	"flag"
	"fmt"
	"github.com/edmore/esp/service"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

var port = flag.String("port", "9999", "default server listening port; can be reset.")

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

	l, e := net.Listen("tcp", ":"+*port)
	log.Println("Server listening on", l.Addr())
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server running ...")
	})
	go http.Serve(l, nil)
	select {}
}
