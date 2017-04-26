package main

import "fmt"

func getUrls() {
	var urls = []string{"https://www.google.com", "https://www.yahoo.com", "https://www.sky.com", "https://www.bbc.co.uk"}
	for _, url := range urls {
		go func() {
			fmt.Println(url)

		}()
	}
}

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

func main() {
	go getUrls()
	var input string
	fmt.Scanln(&input)

}
