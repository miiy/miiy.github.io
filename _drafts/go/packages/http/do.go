package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {

	options("http://api.github.com")
}

func options(url string)  {
	req, err := http.NewRequest("OPTIONS", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Custom-Header", "Custom-Value")
	fmt.Println("Request header:", req.Header)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Header: ", resp.Header)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body:%s\n", string(body))
}