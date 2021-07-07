package main

import (
	"fmt"

	"github.com/ozbekburak/client-bk/client"
)

func main() {
	fmt.Println("Burak Kaan HTTP Client")
	client.Post("https://postman-echo.com/post")
	client.Get("https://jsonplaceholder.typicode.com/posts/1")
	client.Put("https://jsonplaceholder.typicode.com/posts/2")
	client.Delete("https://jsonplaceholder.typicode.com/posts/3")
}
