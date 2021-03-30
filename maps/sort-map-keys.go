package main

import (
	"fmt"
	"sort"
)

func main(){

	employee :=map[string]int{
		"Jitendar":1,
		"Sunny":2,
		"Monalisa":3,
		"Sony":4,
		"Astha":46,
	}
	var keys []string
	for key := range employee{
		keys = append(keys,key)
	}
	sort.Strings(keys)
	fmt.Println(keys)
}