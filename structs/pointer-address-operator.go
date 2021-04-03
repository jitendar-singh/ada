package main

import "fmt"

type plant struct {
	name    string
	indoor  bool
	flowers bool
	fruits  bool
}

func main() {
	var plant1 = &plant{}
	plant1.name = "Rubber Plant"
	plant1.indoor = true
	plant1.flowers = false
	plant1.fruits = false

	var plant2 = &plant{}
	plant2.name = "Erica Palm"
	plant2.indoor = false

	fmt.Println(plant1)
	fmt.Println(plant2)
}
