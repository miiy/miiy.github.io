package main

import (
	"fmt"
	"log"
	"strings"
	"io/ioutil"
	"net/http"
)

func main()  {
	get("http://api.github.com")
}

func get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("StatusCode: %d\n", resp.StatusCode)
	fmt.Printf("Header: %T\n", resp.Header)
	fmt.Printf("Server: %s\n", resp.Header.Get("Server"))
	for k, v := range resp.Header {
		fmt.Printf("%s: %s\n", k, strings.Join(v, ","))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body:%s\n", string(body))
}

/*

http.Get:

func Get(url string) (resp *Response, err error) {
	return DefaultClient.Get(url)
}

http.Response

type Response struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
	// Keys in the map are canonicalized (see CanonicalHeaderKey).
	Header Header
	Body io.ReadCloser
	ContentLength int64
	TransferEncoding []string
	Close bool
	Uncompressed bool
	Trailer Header
	Request *Request
	TLS *tls.ConnectionState
}

http.Header

type Header map[string][]string

ioutil.ReadAll

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func ReadAll(r io.Reader) ([]byte, error) {
	return readAll(r, bytes.MinRead)
}


*/