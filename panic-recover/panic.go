package main

import "fmt"

func main() {
	var choice int
	fmt.Println("Enter 1 for student or 2 if a professional")
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		fmt.Println("I am a student")
	case 2:
		fmt.Println("I am a Professional")
	default:
		panic(fmt.Sprintf("Incorrect choice %d", choice))
	}
	defer func() {
		choice := recover()
		fmt.Println(choice)
	}()
}
