package main

import (
	"crypto/md5"
	"fmt"
)

func md5Checksum(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
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
		got := md5Checksum(tt.s)
		fmt.Printf("md5Checksum(%q) = %v\n", tt.s, got)
	}
}
