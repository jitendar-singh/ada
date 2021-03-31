package main

import (
	"encoding/json"
	"fmt"
)

type employees struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	City string `json:"city"`
}

func main() {
	json_string := `{
		"first_name": "Rocky",
		"last_name": "Sting",
		"city": "London"
	}`

	emp1 := new(employees)
	json.Unmarshal([]byte(json_string),emp1)
	fmt.Println(emp1)
}