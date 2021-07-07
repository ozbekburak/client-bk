package main

import (
	"fmt"

	"github.com/ozbekburak/client-bk/client"
)



func main() {

	todoData := client.Todo{
		5,
		2,
		"Update the second element of todos",
		true,
	}

	fmt.Println("Burak Kaan HTTP Client")

	fmt.Printf("\nResponse of POST request: %s\n", string(client.Post("https://postman-echo.com/post")))
	fmt.Printf("\nResponse of GET request:\n%+v\n", client.Get("https://jsonplaceholder.typicode.com/todos/1"))
	fmt.Printf("\nResponse of PUT request:\n%+v\n", client.Put("https://jsonplaceholder.typicode.com/todos/2", todoData))

	status, body := client.Delete("https://jsonplaceholder.typicode.com/todos/3")
	fmt.Printf("\nResponse of DELETE request:\nStatus: %v\nBody: %v", status, body)
}
