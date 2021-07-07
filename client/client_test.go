package client

import "testing"

func TestGet(t *testing.T) {
	data := Todo{
		UserID:    1,
		ID:        1,
		Title:     "delectus aut autem",
		Completed: false,
	}
	result := Get("https://jsonplaceholder.typicode.com/todos/1")
	if result != data {
		t.Errorf("Expected %v, Actual %v", data, result)
	}
}

func TestDelete(t *testing.T) {
	status, body := Delete("https://jsonplaceholder.typicode.com/todos/4")
	if status != "200 OK" && body != ""{
		t.Errorf("Expected status %v, Actual status: %v", "200 OK", status)
		t.Errorf("Expected body %v, Actual body: %v", "", body)
	}
}

func TestPut(t *testing.T)  {
	data := Todo{
		UserID: 10,
		ID: 3,
		Title: "Update data for test",
		Completed: true,
	}
	result := Put("https://jsonplaceholder.typicode.com/todos/3", data)
	if result != data {
		t.Errorf("Expected %v, Actual %v", data, result)
	}

}
