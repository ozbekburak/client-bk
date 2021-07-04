package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Burak Kaan HTTP Client")
	Get("https://jsonplaceholder.typicode.com/posts/1")
}

func Get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s", err) // log kütüphanesini kullanmak daha doğru olabilir?
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Printf("String body: %s", sb)
}