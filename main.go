package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Mentor struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func main() {
	fmt.Println("Burak Kaan HTTP Client")
	Get("https://jsonplaceholder.typicode.com/posts/1")
	Post("https://postman-echo.com/post")
}

func Get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s", err) // log kütüphanesini kullanmak daha doğru olabilir?
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// string(body) ayrı bir değişkene atamak overkill mi best practice mi?
	fmt.Printf("Get request response body: %s", string(body))
}

func Post(url string) {
	m := Mentor{
		"Erhan",
		"Yakut",
	}
	postBody, _ := json.Marshal(m)
	requestBody := bytes.NewReader(postBody)

	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}


	body, err := ioutil.ReadAll(resp.Body) // byte döner
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("\nPost request response body: %s", string(body))
}
