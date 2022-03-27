package main

import (
	"fmt"
	"link-validator/linkValidator"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://golang.org/")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range linkValidator.GetLinks(resp.Body) {
		fmt.Println(v)
	}
}
