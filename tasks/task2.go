package main

import (
	"log"
	"sync"
)

/* Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.*/

func main() {
	wg := sync.WaitGroup{}
	np := 2

	wg.Add(np)

	numbers := []int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9, 11}

	for i := 0; i < np; i++ {
		go sqrt(i, np, numbers, &wg)
	}

	wg.Wait()

}

func sqrt(start int, step int, s []int, wg *sync.WaitGroup) {
	for i := start; i < len(s); i += step {
		log.Println(s[i] * s[i])
	}
	wg.Done()
}
