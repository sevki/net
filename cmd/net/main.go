package main

import (
	"log"
	"net/http"

	"sevki.org/net/net"
	"sevki.org/net/transport/v1"
)

func main() {
	netSvc := net.New()
	srv := v1.NewNetHTTPServer(netSvc)
	log.Fatal(http.ListenAndServe(":8080", srv))
}
