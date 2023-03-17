package main

import "testing"

func TestChangeBit(t *testing.T) {
	var from int64 = 32
	var to int64 = 40

	if result := ChangeBit(from, 4, true); result != to {
		t.Fatalf("Wrong answer: want %d, got %d", to, result)
	}
}

func TestChangeBit2(t *testing.T) {
	var from int64 = 48
	var to int64 = 32

	if result := ChangeBit(from, 5, false); result != to {
		t.Fatalf("Wrong answer: want %d, got %d", to, result)
	}
}
