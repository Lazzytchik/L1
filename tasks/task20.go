package main

import (
	"log"
	"strings"
)

/*
	20. Разработать программу, которая переворачивает слова в строке.
		Пример: «snow dog sun — sun dog snow».
*/

func main() {
	s := "snow dog sun"
	log.Println(ReverseSentence(s))
}

func ReverseSentence(s string) string {
	splited := strings.Fields(s)
	result := strings.Builder{}

	for _, word := range splited {
		result.WriteString(word)
		result.WriteString(" ")
	}

	return result.String()
}
