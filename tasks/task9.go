package main

import (
	"log"
	"sync"
)

/*
	9. Разработать конвейер чисел. Даны два канала:
		в первый пишутся числа (x) из массива,
		во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(2)

	// Горутина на считывание из второго канала
	go func() {
		for {
			if a, ok := <-ch2; ok {
				log.Println(a)
			} else {
				break
			}
		}
		wg.Done()
	}()

	// Горутина на считывание из первого канала
	go func() {
		for {
			a, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- a * a
		}
		close(ch2)
		wg.Done()
	}()

	// Заполнение канала из массива
	for _, v := range slice {
		ch1 <- v
	}

	close(ch1)

	wg.Wait()
}
