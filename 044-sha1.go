package main

import (
	"crypto/sha1"
	"fmt"
)

func Sha1Checksum(s string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}
func main() {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "empty string",
			s:    "",
		},
		{
			name: "hello world",
			s:    "hello world",
		},
	}
	for _, tt := range tests {
		got := Sha1Checksum(tt.s)
		fmt.Printf("sha1Checksum(%q) = %v\n", tt.s, got)
	}
}
