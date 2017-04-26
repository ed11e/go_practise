package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getUrls() {
	var urls = []string{"https://www.google.com", "https://www.yahoo.com", "https://www.sky.com", "https://www.bbc.co.uk"}
	for _, url := range urls {
		fmt.Println("Fetching: " + url)
		resp, err := http.Get(url)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func checkDone() {
	fmt.Println("Press Enter to terminate the program")
}

func main() {
	go getUrls()
	var input string
	fmt.Scanln(&input)

}
