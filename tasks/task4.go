package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

/*
	4. 	Реализовать постоянную запись данных в канал (главный поток).
		Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
		Необходима возможность выбора количества воркеров при старте.
		Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	mainChannel := make(chan int)
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-sigchnl
		log.Println("Stopping workers...")
		cancel()
		os.Exit(0)
	}()

	StartWorkers(mainChannel, 4, ctx)

	for {
		mainChannel <- rand.Int()
	}

}

func StartWorkers(ch chan int, N int, ctx context.Context) {
	for i := 0; i < N; i++ {
		go func(i int) {
			var c int
			for {
				select {
				case c = <-ch:
					log.Printf("Worker %d: Got: %d", i+1, c)
				case <-ctx.Done():
					log.Printf("Sttopint worker %d...", i)
					return
				}
			}
		}(i)
	}
}
