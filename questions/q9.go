package main

import "log"

/*
	9. Сколько существует способ задать переменную типа slice или map?

*/

func main() {
	slice1 := []int{1, 2, 3}
	log.Println(slice1)

	slice2 := make([]int, 2, 3)
	log.Println(slice2)

	var slice3 = new([]int)
	log.Println(slice3)
}
