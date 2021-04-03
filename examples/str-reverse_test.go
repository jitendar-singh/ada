package examples

import (
	"fmt"
	"testing"
)

func TestStrRev(t *testing.T) {
	tests := []struct {
		word     string
		expected string
	}{
		{"sunny", "ynnus"},
		{"abcd", "dcb"},
		{"日本語", "語本日"},
	}
	fmt.Println("Starting Reverse string Test")
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v:", tc.word), func(t *testing.T) {
			got := strReverse(tc.word)
			if got != tc.expected {
				t.Fatalf("Reverse() = %v; want %v; got %v", tc.word, got, tc.expected)
			}
		})
	}

}
