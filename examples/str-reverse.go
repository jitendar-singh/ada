package examples

func strReverse(str string) string {
	var res string
	for _, char := range str {
		res = string(char) + res
	}
	return res
}
