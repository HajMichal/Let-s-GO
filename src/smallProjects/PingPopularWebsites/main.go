package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	links := []string {
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://www.amazon.com",
	}

	for _, link := range links {
		go handleGetReq(link)
	}
}

func handleGetReq(link string) bool {
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(link + " is up!" )
	if(resp.StatusCode == 200) {
		return true
	} else {
		return false
	}
}