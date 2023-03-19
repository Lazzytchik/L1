package main

import "testing"

/*
	abcd — true
	abCdefAaf — false
	aabcd — false
*/

func Test261(t *testing.T) {
	s := "abcd"

	if got := CheckUniqueChar(s); !got {
		t.Fatalf("Wrong Answer: Got: %t Want: true", got)
	}
}

func Test262(t *testing.T) {
	s := "abCdefA"

	if got := CheckUniqueChar(s); got {
		t.Fatalf("Wrong Answer: Got: %t Want: false", got)
	}
}

func Test263(t *testing.T) {
	s := "aabcd"

	if got := CheckUniqueChar(s); got {
		t.Fatalf("Wrong Answer: Got: %t Want: false", got)
	}
}
