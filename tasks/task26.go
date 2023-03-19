package main

import "strings"

/*
	26.	Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
		Функция проверки должна быть регистронезависимой.

		Например:
		abcd — true
		abCdefAaf — false
		aabcd — false

		Для проверки данных используем мапу по аналогии с 11 и 12 задачами.
		Предварительно переведём строку в нижний регистр, чтобы нормализовать всех по нему т.к. он не важен.
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
