package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Age   int `json:"age"`
	Class int `json:"class"`
	Roll  int `json:"roll"`
}

func main() {
	json_string := `{
	"age": 1,
	"class": 4,
	"roll": 23
	}`
	Ramesh := new(Student)
	json.Unmarshal([]byte(json_string), Ramesh)
	fmt.Println(Ramesh)
}
