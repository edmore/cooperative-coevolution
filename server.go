package main

import "net"
import "net/http"
import "net/rpc"
import "log"
import "github.com/edmore/esp/service"

func main() {
	arith := new(service.Arith)
	ping := new(service.Ping)

	rpc.Register(arith)
	rpc.Register(ping)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	log.Println("ESP Server listening on", l.Addr())
	if e != nil {
		log.Fatal("listen error:", e)
	}

	http.Serve(l, nil)
}
