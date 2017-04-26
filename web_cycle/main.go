package main

import (
	"fmt"
	"time"
)

func toRun() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(i)
	}
}

func main() {

	// var urls = []string{"https://www.google.com", "https://www.yahoo.com", "https://www.sky.com", "https://www.bbc.co.uk"}

	go toRun()

	var input string
	fmt.Scanln(&input)

	// for _, url := range urls {
	// 	fmt.Println("Fetching: " + url)
	// 	resp, err := http.Get(url)
	//
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	} else {
	// 		body, err := ioutil.ReadAll(resp.Body)
	// 		fmt.Println(string(body))
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	}
	// }
}
