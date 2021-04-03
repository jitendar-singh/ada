package main

import (
	"fmt"
	"sort"
)

func main() {

	employee := map[string]int{"Jitendar": 4, "Sony": 2, "Monlaisa": 1, "Gonalisa": 3}
	for keys, values := range employee {
		fmt.Println(keys, ":", values)
	}

	var values []int
	for _, value := range employee {
		values = append(values, value)
	}
	sort.Ints(values)
	fmt.Println(values)
}
