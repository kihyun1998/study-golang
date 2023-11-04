package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
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
		hitURL(url)
	}
}

func hitURL(url string) error {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
	return nil
}
