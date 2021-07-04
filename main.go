package main

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

// PostData güzel bir isimlendirme değil gibi ama Post() fonksiyon ismiyle çakışıyor
type PostData struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	fmt.Println("Burak Kaan HTTP Client")
	Post("https://postman-echo.com/post")
	Get("https://jsonplaceholder.typicode.com/posts/1")
	Put("https://jsonplaceholder.typicode.com/posts/2")
	Delete("https://jsonplaceholder.typicode.com/posts/3")
}


func Post(url string) {
	myMentor := Mentor{
		"Erhan",
		"Yakut",
	}

	data, _ := json.Marshal(myMentor)
	requestBody := bytes.NewReader(data)

	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		fmt.Printf("Error: %v", err) // log kütüphanesini kullanmak daha doğru mu olur? (log.Fatalln)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) // byte döner (byte[]). bodyBytes daha iyi bir adlandırma mı olur?
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	// string(body) ayrı bir değişkene atamak overkill mi best practice mi?
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

	// struct olarak göstermek için unmarshal edip postdata'yı dolduruyoruz. Mantıklı mı? yoksa direkt string(body) mi
	// kullanmak iyi olurdu?
	var post PostData
	json.Unmarshal(body, &post)
	fmt.Printf("\nResponse of GET request:\n%+v\n", post) // +v fieldları da yazıyor.
}

func Put(url string){
	postData := PostData{
		5,
		2,
		"Update the second element of posts",
		"This is our new body for posts/2",
	}

	client := &http.Client{}
	data, _ := json.Marshal(postData)
	requestBody := bytes.NewBuffer(data)

	req, err := http.NewRequest(http.MethodPut, url, requestBody) // http.Put yok.
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

func Delete(url string)  {
	postData := PostData{
		5,
		3,
		"Delete the post with an id of 3",
		"Goodbye, body.",
	}

	client := &http.Client{}
	data, _ := json.Marshal(postData)
	requestBody := bytes.NewBuffer(data)

	req, err := http.NewRequest(http.MethodDelete, url, requestBody) // http.Put yok.
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