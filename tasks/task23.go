package main

import "log"

/*
	23. Удалить i-ый элемент из слайса.

	Совмещаем всё что до выбранного элемента в слайсе с тем что после него и возвращаем.
*/

func main() {
	s := []int{2, 5, 8, 1, 9, 2, 8, 4}
	log.Println(s)

	s = Delete(s, 5)
	log.Println(s)
}

func Delete[T comparable](slice []T, i int) []T {
	return append(slice[:i], slice[i+1:]...)
}
