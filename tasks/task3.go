package main

import (
	"log"
	"sync"
)

/* Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.*/

func main() {
	wg := sync.WaitGroup{}
	np := 2

	wg.Add(np)
	m := sync.Mutex{}

	numbers := []int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9, 11}
	sum := 0

	for i := 0; i < np; i++ {
		go sqrtSum(i, np, numbers, &sum, &m, &wg)
	}

	wg.Wait()

	log.Println(sum)

}

func sqrtSum(start int, step int, s []int, sum *int, m *sync.Mutex, wg *sync.WaitGroup) {
	for i := start; i < len(s); i += step {
		m.Lock()
		*sum = *sum + s[i]*s[i]
		m.Unlock()
	}
	wg.Done()
}
