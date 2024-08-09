package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)


func main() {
	links := []string {
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://www.amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go handleGetReq(link, c)
	}

	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			handleGetReq(l, c)
		}()
	}
}

func handleGetReq(link string, c chan string) bool {
	resp, err := http.Get(link)
	if err != nil {
		c <- link
		log.Fatal(err)
	}
	fmt.Println(link + " is up!" )
	c <- link

	if(resp.StatusCode == 200) {
		return true
	} else {
		return false
	}
}