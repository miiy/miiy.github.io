package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

/*
# go build clock2.go
# ./clock2 &
[1] 534
# ./netcat1
03:05:58
03:05:59
03:06:00
03:06:01
03:06:02
03:06:03
03:06:04
03:06:05
^C


# ./netcat1
03:06:01
03:06:02
03:06:03
03:06:04
03:06:05
^C

*/