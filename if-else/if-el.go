package main

import "fmt"

func main() {
	i := 95
	for {
		if i == 100 {
			fmt.Println("JAPAN")
			return
		} else {
			i++
			fmt.Println(i)
		}
	}
}
