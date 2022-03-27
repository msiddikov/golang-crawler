package main

import (
	"fmt"
	"time"

	crawler "github.com/msiddikov/golang-crawler"
)

func main() {
	url := "https://go.dev"
	started := time.Now()
	fmt.Printf("Started crawling %s \n", url)
	res := crawler.Start(url, 5, 100)
	var valid, invalid uint
	for _, v := range res {
		if v.IsValid {
			fmt.Printf("[VALID] url: %s \n", v.Url)
			valid += 1
		} else {
			fmt.Printf("[INVALID] url: %s error: %s\n", v.Url, v.Err)
			invalid += 1
		}

	}
	d := time.Now().Sub(started)
	fmt.Printf("Found %v links in : %v seconds \nValid links: %v \nBroken links: %v", len(res), d.Seconds(), valid, invalid)
}
