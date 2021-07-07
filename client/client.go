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

type PostData struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
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

func Get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	var post PostData
	json.Unmarshal(body, &post)
	fmt.Printf("\nResponse of GET request:\n%+v\n", post) // +v fieldları da yazıyor.
}

func Put(url string) {
	postData := PostData{
		5,
		2,
		"Update the second element of posts",
		"This is our new body for posts/2",
	}

	client := &http.Client{}
	data, _ := json.Marshal(postData)
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

	var post PostData
	json.Unmarshal(body, &post)
	fmt.Printf("\nResponse of PUT request:\n%+v\n", post)
}

func Delete(url string) {
	postData := PostData{
		5,
		3,
		"Delete the post with an id of 3",
		"Goodbye, body.",
	}

	client := &http.Client{}
	data, _ := json.Marshal(postData)
	requestBody := bytes.NewBuffer(data)

	req, err := http.NewRequest(http.MethodDelete, url, requestBody)
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

	var post PostData
	json.Unmarshal(body, &post)
	fmt.Printf("\nResponse of DELETE request:\n%+v\n", post)
}
