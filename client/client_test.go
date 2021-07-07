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
