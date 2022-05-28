package main

import (
	"errors"
	"fmt"
	http "net/http"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.naver.com",
		"https://www.yahoo.com",
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	var results = make(map[string]string)

	for _, url := range urls {
		err := hitUrl(url)

		var result = "OK"
		if err != nil {
			result = "FAILED"
		}

		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

var errRequestFailed = errors.New("request failed")

func hitUrl(url string) error {
	fmt.Println("Checking", url)

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}

	return nil
}
