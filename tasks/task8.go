package main

import (
	"log"
)

/*
	8. Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.

	Если мы хотим поставить бит в 1, то мы делаем число 2 в степени i - 1 (100000...). и делаем побитовую дизъюнкцию.
	Если мы хотим поставить бит в 0, то мы делаем побитовую конъюнкцию с разницей между исходным числом и числом 2 в степени i - 1 (111101111...)
*/

func main() {
	c := ChangeBit(39, 3, false)

	log.Println(c)

}

func ChangeBit(number int64, i int, bit bool) int64 {
	if bit {
		mask := int64(1 << (int64(i) - 1))
		return number | mask
	} else {
		mask := int64(number - (1 << (i - 1)))
		return number & mask
	}
}
