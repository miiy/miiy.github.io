package main

import (
	"log"
	"net/http"
)

func main()  {
	post("")
}
func post(url string)  {
	resp, err := http.Post(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.StatusCode)
	}
}