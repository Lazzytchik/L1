package main

import (
	"log"
)

/*
	8. Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.
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
