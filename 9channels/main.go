package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
	}

	// channel
	c := make(chan string)

	for _, link := range links {
		go checkLinks(link, c)
	}

	// for {
	// 	go checkLinks(<-c, c)
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLinks(link, c)
		}(l)
	}
}

func checkLinks(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "error")
		// c <- "chanel: err"
		c <- link
		return
	}

	fmt.Println(link, "working fine")
	// c <- "chanel: gud"
	c <- link
}