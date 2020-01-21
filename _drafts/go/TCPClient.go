package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s host:port", os.Args[0])
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveIPAddr("ip4", service)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tcpAddr.String())
	conn, err := net.Dial("tcp", tcpAddr.String())
	if err != nil {
		log.Fatal(err)
	}
	n, err := conn.Write([]byte("HEAD / HTTP/1.1 \r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(n)
}