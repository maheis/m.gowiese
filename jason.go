package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func myJson() {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var person Person
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(person.Name)
	fmt.Println(person.Age)
	fmt.Println(person.Email)
}
