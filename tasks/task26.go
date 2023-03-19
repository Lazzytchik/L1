package main

import "strings"

/*
	26.	Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
		Функция проверки должна быть регистронезависимой.

		Например:
		abcd — true
		abCdefAaf — false
		aabcd — false
*/

func main() {

}

func CheckUniqueChar(s string) bool {
	s = strings.ToLower(s)
	seen := make(map[int32]struct{})

	for _, char := range s {
		if _, ok := seen[char]; ok {
			return false
		}
		seen[char] = struct{}{}
	}

	return true
}
