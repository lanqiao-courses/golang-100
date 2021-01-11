package main

import (
	"crypto/sha256"
	"fmt"
)

func Sha256Checksum(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
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
		got := Sha256Checksum(tt.s)
		fmt.Printf("Sha256Checksum(%q) = %v\n", tt.s, got)
	}
}
