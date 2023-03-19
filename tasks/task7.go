package main

import (
	"log"
	"sync"
)

/*
	7. Реализовать конкурентную запись данных в map.
*/

func main() {
	m := make(map[int]int)

	var mutex sync.Mutex
	var wg = sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for i := 0; i < 100; i++ {
			mutex.Lock()
			m[i] = i
			mutex.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 100; i < 200; i++ {
			mutex.Lock()
			m[i] = i
			mutex.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()

	log.Println(m)
}
