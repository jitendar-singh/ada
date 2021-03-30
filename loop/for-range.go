package main

import "fmt"

func main()  {
	strDct := map[string]string{
		"Japan":"Tokio",
		"China":"Beijing",
		"India":"New Delhi",
		"England":"London",
	}
	for index,_ := range strDct{
		fmt.Println("Capital of ",index)
	}

	for _,key:=range strDct{
		fmt.Println(key)
	}
}