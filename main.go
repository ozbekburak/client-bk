package main

import (
	"fmt"
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

}