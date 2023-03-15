package main

import (
	"fmt"
	"sync"
)

/*
	11. Что выведет данная программа и почему?

	Программа завершится deadlock'ом т.к. все горутины завершены, однако WaitGroup всё ещё ждёт завершения работы горутин.
	Это связанно с тем, что мы передаём в анонимную функцию копию WaitGroup'ы, а не ссылку на оригинальную.

*/

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}

	wg.Wait()

	fmt.Println("exit")
}
