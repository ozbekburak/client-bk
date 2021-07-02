package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Burak Kaan HTTP Client")
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	Get(&c, "https://api.github.com/")
}

func Get(c *http.Client, url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
	fmt.Printf("Body : %s \n ", body)
	fmt.Printf("Response status : %s \n", resp.Status)
}
