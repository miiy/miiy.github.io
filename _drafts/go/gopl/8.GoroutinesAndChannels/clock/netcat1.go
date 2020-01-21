package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy (dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

/*
# go build netcat1.go
# ./netcat1
02:59:03
02:59:04
02:59:05
02:59:06
^C

# ./netcat1
02:59:08
02:59:09
02:59:10
02:59:11
^C

 */