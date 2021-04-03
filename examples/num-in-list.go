package examples

import "fmt"

func numinList(list []int, key int) bool {
	if len(list) == 0 {
		fmt.Println("Empty list")
		return false
	} else {
		for _, value := range list {
			if value == key {
				return true
			}
		}
		return false
	}
}
