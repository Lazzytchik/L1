package main

import (
	"context"
	"fmt"
	"time"
)

/*
	Способы завершить горутину.
	 - По закрытию канала
	 - С помощью селекта и другого канала
	 - С помощью context
	Завершить горутину напрямую из другой горутины нельзя.
*/

func main() {
	ChannelStop()
	SemaphoreStop()
	ContextStop()
}

func ChannelStop() {
	//	Заканчивает по закрытию канала
	ch := make(chan string, 6)
	go func() {
		for {
			v, ok := <-ch
			if !ok {
				fmt.Println("Finish")
				return
			}
			fmt.Println(v)
		}
	}()

	ch <- "hello"
	ch <- "world"
	close(ch)
	time.Sleep(time.Second)
}

func SemaphoreStop() {
	//	Заканчивает с помощью другого канала.
	ch := make(chan string, 6)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case ch <- "foo":
			case <-done:
				close(ch)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()

	for i := range ch {
		fmt.Println("Received: ", i)
	}

	fmt.Println("Finish")
}

func ContextStop() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				ch <- struct{}{}
				return
			default:
				fmt.Println("foo...")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-ch
	fmt.Println("Finish")
}
