package examples

import (
	"fmt"
	"testing"
)

func Test_sumList(t *testing.T) {
	tests := []struct {
		list        []int
		expectedSum int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 36},
		{[]int{-1, -2, -3, -4, -5, -6, -7, -8}, -36},
		{[]int{}, 0},
		{[]int{-1, 0}, -1},
		{[]int{0, 0}, 0},
	}
	fmt.Println("Starting sum-List test")
	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v)", tc.list), func(t *testing.T) {
			got := SumList(tc.list)
			if got != tc.expectedSum {
				t.Fatalf("NumList()= %v; sum= %v", got, tc.expectedSum)
			}
		})
	}
}
