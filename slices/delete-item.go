package main

import "fmt"

func remove(slice []string, s int) []string {
	// slice unpacking slice[s+1:]...
	return append(slice[:s], slice[s+1:]...)
}

func main() {

	var names = make([]string, 5)
	names[0] = "sunny"
	names[1] = "lola kutti"
	names[2] = "Monlaisa"
	names = remove(names, 1)
	fmt.Println(names)
}

//# Dry run
//["sunny","sony","tuni","guni","muni"]
//s=2
//slice[:2] = "sunny","sony" // previuos elements to the key
//slice[s+1:] = "guni","muni"// elements post the key
