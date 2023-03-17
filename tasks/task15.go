package main

import (
	"log"
	"strings"
)

/*
	15. К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
*/

var justString string

func main() {
	log.Println(justString)
	someFunc()
	log.Println(justString, len(justString))
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func createHugeString(size int) string {
	var sb strings.Builder

	for i := 0; i < size; i++ {
		sb.WriteString("s")
	}

	return sb.String()
}
