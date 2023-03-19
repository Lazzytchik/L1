package main

import (
	"log"
	"sync"
)

/*
	18. Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
		По завершению программа должна выводить итоговое значение счетчика.

		Т.к. несколько горутин записывают к переменной используем Mutex
		Т.к. Нужно подождать завершения горутин используем WaitGroup
*/

type Counter struct {
	X int
}

func (c *Counter) Add(m *sync.Mutex) {
	m.Lock()
	c.X += 1
	m.Unlock()
}

func main() {
	counter := Counter{0}
	var mutex sync.Mutex

	N := 4

	wg := sync.WaitGroup{}
	wg.Add(N)

	for i := 0; i < N; i++ {
		go func(mutex *sync.Mutex, counter *Counter) {
			for j := 0; j < 1000; j++ {
				counter.Add(mutex)
			}
			wg.Done()
		}(&mutex, &counter)
	}

	wg.Wait()

	log.Println(counter.X)
}
