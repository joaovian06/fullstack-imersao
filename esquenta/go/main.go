package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	Name 		string `json:"course"`
	Description string `json:"description"`
	Price 		int    `json:"price"`
}

func (c Course) GetFullInfo() string {
	return fmt.Sprintf("Name: %s, Description: %s, Price: %d", c.Name, c.Description, c.Price)
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) { 
	course := Course{
		Name: "Golang",
		Description: "Golang Course",
		Price: 100,
	}

	json.NewEncoder(w).Encode(course)
}