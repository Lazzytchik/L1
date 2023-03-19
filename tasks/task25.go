package main

import (
	"log"
	"time"
)

/*
	25. Реализовать собственную функцию sleep

	Заставим исполнять программу до тех пор, пока текущее время не станет равным или больше чем текущее время + параметр длительности.
*/

func main() {
	duration := 2 * time.Second

	log.Println("Before")
	sleep(duration)
	log.Println("After")
}

func sleep(duration time.Duration) {
	to := time.Now().Add(duration)
	for time.Now().Unix() < to.Unix() {
	}
}
