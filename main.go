package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Burak Kaan HTTP Client")
	Get("https://www.google.com")
}

func Get(url string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	response, err := c.Get(url)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Printf("Body: %s", body)
}
