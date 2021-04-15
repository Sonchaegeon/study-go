package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}
	for _, url := range urls {
		hitURL(url)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking: " + url)
	res, err := http.Get(url)
	if err == nil || res.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil;
}