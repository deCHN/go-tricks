package main_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func waiter(i int, block, done chan struct{}) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	fmt.Println(i, "waiting...")
	<-block
	fmt.Println(i, "done!")
	done <- struct{}{}
}

func waiterWithContext(i int, ctx context, done chan struct{}) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	fmt.Println(i, "waiting...")
	select {
	case context.Done():
		done <- struct{}{}
	default:

	}
	fmt.Println(i, "done!")
	done <- struct{}{}
}

func TestWaiter(t *testing.T) {
	block, done := make(chan struct{}), make(chan struct{})
	for i := 0; i < 4; i++ {
		go waiter(i, block, done)
	}
	time.Sleep(5 * time.Second)
	close(block)
	for i := 0; i < 4; i++ {
		<-done
	}
}
