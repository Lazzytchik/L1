package main

import (
	"log"
	"time"
)

/*
	25. Реализовать собственную функцию sleep
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
