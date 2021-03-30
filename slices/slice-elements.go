package main

import "fmt"

func main()  {
	var countries = []string{"India","Pakisthan","China","Tibet","Afganisthan"}
	// Entire Slice
	fmt.Println(countries)
	// Slice from 2nd element
	fmt.Println(countries[1:])
	//slice from 2nd till 4th element
	fmt.Println(countries[1:len(countries)-1])
	//last two elements
	fmt.Println(countries[len(countries)-2:])
}
