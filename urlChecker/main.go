package main

import (
	"fmt"
	"net/http"
)

type request struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan request)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, ":", status)
	}
}

func hitURL(url string, c chan<- request) {
	_, err := http.Get(url)
	status := "OK"
	if err != nil {
		status = "FAILED"
	}
	c <- request{url: url, status: status}

}
