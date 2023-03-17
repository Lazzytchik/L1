package main

import "log"

/*
	13. Поменять местами два числа без создания временной переменной.
*/

func main() {
	a := 2
	b := 3

	log.Println(a, b)

	a, b = b, a

	log.Println(a, b)
}
