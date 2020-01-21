package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	head("http://api.github.com")
}

func head(url string)  {
	resp, err := http.Head(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal("error.")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)
	fmt.Println(string(body))
}

/*
200
map[Access-Control-Allow-Origin:[*] Access-Control-Expose-Headers:[ETag, Link, Location, Retry-After, X-GitHub-OTP, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset, X-OAuth-Scopes, X-Accepted-OAuth-Scopes, X-Poll-Interval, X-GitHub-Media-Type] Cache-Control:[public, max-age=60, s-maxage=60] Content-Length:[2153] Content-Security-Policy:[default-src 'none'] Content-Type:[application/json; charset=utf-8] Date:[Wed, 08 Jan 2020 08:28:32 GMT] Etag:["af8c3ea475a768b4c2a0113f6377c65f"] Referrer-Policy:[origin-when-cross-origin, strict-origin-when-cross-origin] Server:[GitHub.com] Status:[200 OK] Strict-Transport-Security:[max-age=31536000; includeSubdomains; preload] Vary:[Accept Accept-Encoding] X-Content-Type-Options:[nosniff] X-Frame-Options:[deny] X-Github-Media-Type:[github.v3; format=json] X-Github-Request-Id:[C3C4:1983:218F15:2C6151:5E1592AF] X-Ratelimit-Limit:[60] X-Ratelimit-Remaining:[59] X-Ratelimit-Reset:[1578475712] X-Xss-Protection:[1; mode=block]]
*/