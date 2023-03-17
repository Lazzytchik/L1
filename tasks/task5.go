package main

/*
	5.	Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
		По истечению N секунд программа должна завершаться.
*/

import (
	"log"
	"time"
)

func main() {
	MakeBridge(2)
}

func MakeBridge(n int) {
	var man = make(chan int)
	var quit = make(chan struct{})

	go printChannel(man)
	go fillChannel(man, quit)

	time.Sleep(time.Duration(int64(n)) * time.Second)
	quit <- struct{}{}
}

func printChannel(man chan int) {
	for {
		log.Println(<-man)
	}
}

func fillChannel(ch chan int, cancel chan struct{}) {
	for {
		select {
		case ch <- 1:
		case <-cancel:
			break
		}
	}
}
