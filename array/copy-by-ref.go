package main

import "fmt"

func main() {

	names := [...]string{
		"Sunny", "Sony", "Nono",
	}
	fullNames := &names
	fmt.Println("Names before rectification:", names)
	fmt.Println("FullNames before rectification:", *fullNames)

	names[0] = "Jitendar"
	fmt.Println("Names after rectification:", names)
	fmt.Println("FullNames after rectification:", *fullNames)

}
