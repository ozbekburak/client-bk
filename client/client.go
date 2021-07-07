package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Mentor struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Post(url string) {
	myMentor := Mentor{
		"Erhan",
		"Yakut",
	}

	data, _ := json.Marshal(myMentor)
	requestBody := bytes.NewBuffer(data)

	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) // returns byte (byte[]).
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("\nResponse of POST request: %s\n", string(body))
}

func Get(url string) interface{} {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	var todo Todo
	json.Unmarshal(body, &todo)
	//fmt.Printf("\nResponse of GET request:\n%+v\n", post) // +v fieldları da yazıyor.
	return todo
}

func Put(url string) {
	todoData := Todo{
		5,
		2,
		"Update the second element of todos",
		true,
	}

	client := &http.Client{}
	data, _ := json.Marshal(todoData)
	requestBody := bytes.NewBuffer(data)

	req, err := http.NewRequest(http.MethodPut, url, requestBody)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var todo Todo
	json.Unmarshal(body, &todo)
	fmt.Printf("\nResponse of PUT request:\n%+v\n", todo)
}

func Delete(url string) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Printf("\nResponse of DELETE request: \nStatus: %v\nBody: %v", resp.Status, string(body)) // 502 dönüyor, silme başarılı diyebiliriz?

}
