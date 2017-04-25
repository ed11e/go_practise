package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	var urls = []string{"https://www.google.com", "https://www.yahoo.com", "https://www.sky.com", "https://www.bbc.co.uk"}

	for _, url := range urls {
		fmt.Println("Fetching: " + url)
		resp, err := http.Get(url)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			bodyBytes, err2 := ioutil.ReadAll(resp.Body)
			bodyString, _ := string(bodyBytes)
			fmt.Println(bodyString)
		}
	}
}
