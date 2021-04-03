package examples

func SumList(list []int) int {
	var sum int
	for _, value := range list {
		sum = sum + value
	}
	return sum
}
