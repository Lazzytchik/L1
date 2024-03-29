package main

import (
	"log"
	"sync"
)

/*
	7. Реализовать конкурентную запись данных в map.

	Maps are not safe for concurrent use: it’s not defined what happens when you read and write to them simultaneously.
	If you need to read from and write to a map from concurrently executing goroutines,
		the accesses must be mediated by some kind of synchronization mechanism.
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
