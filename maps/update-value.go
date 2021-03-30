package main

import "fmt"

func main()  {
	employee := map[string]int{
		"Jitendar":1,
		"Sunny":2,
	}
	fmt.Println("Map before alteration:")
	for key,value := range employee{
		fmt.Println(key,":",value)
	}
	employee["Sunny"]=3
	fmt.Println("Map post updation")
	for key,value := range employee{
		fmt.Println(key,":",value)
	}

}
