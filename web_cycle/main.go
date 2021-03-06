package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var urls = []string{"https://www.google.com", "https://www.yahoo.com", "https://www.sky.com", "https://www.bbc.co.uk"}
	resultsChan := make(chan string)
	for _, url := range urls {
		go func() {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			} else {
				body, err := ioutil.ReadAll(resp.Body)
				resultsChan <- string(body)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}

	for item := range urls {
		fmt.Println(<-resultsChan)
		fmt.Println(item)
	}
}
