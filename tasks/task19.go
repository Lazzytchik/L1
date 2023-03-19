package main

import "log"

/*
	19. Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
		Символы могут быть unicode.

		Т.к. Символы могут быть Unicode используем runes.
*/

func main() {
	s := "ঐআঊ3"
	log.Println(Reverse(s))
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
