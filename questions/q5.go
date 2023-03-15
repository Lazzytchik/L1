package main

/*
	5. Какой размер у структуры struct{}{}
	0 байт
*/

import (
	"fmt"
	"unsafe"
)

type S1 struct {
	f1 int
}

func main() {
	s1 := S1{}
	s2 := struct{}{}

	fmt.Printf("s1 size: %v\n", unsafe.Sizeof(s1))
	fmt.Printf("s2 size: %v\n", unsafe.Sizeof(s2))
}
