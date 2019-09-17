package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(finishReqBuffered(time.Millisecond * 40))
	fmt.Println(finishReqUnbuffered(time.Millisecond * 40))

	<-time.After(time.Second)
}

func finishReqUnbuffered(timeout time.Duration) string {
	ch := make(chan string)
	go func() {
		ch <- result() // Blocks if timeout finishes first
		close(ch)
		fmt.Println("unblocked")
	}()

	select {
	case result := <-ch:
		return result
	case <-time.After(timeout):
		return "timed out"
	}
}

func finishReqBuffered(timeout time.Duration) string {
	ch := make(chan string, 1)
	go func() {
		ch <- result()
		close(ch)
		fmt.Println("unblocked")
	}()

	select {
	case result := <-ch:
		return result
	case <-time.After(timeout):
		return "timed out"
	}
}

func result() string {
	<-time.After(time.Millisecond * 200)
	return "did not time out"
}
