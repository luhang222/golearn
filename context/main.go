package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx := context.Background()
	ctxWithCancel, cancelFunc := context.WithCancel(ctx)
	defer func() {
		fmt.Println("Main Defer: canceling context")
		cancelFunc()
	}()
	go func() {
		sleepRandom("Main", nil)
		cancelFunc()
		fmt.Println("Main Sleep Completed. Canceling context")
	}()
	doWorkContext(ctxWithCancel)
}

func sleepRandom(fromFunction string, ch chan int) {
	defer func() {
		fmt.Println(fromFunction, "sleepRandom completed")
	}()
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomNumber := r.Intn(100)
	sleepTime := randomNumber + 100
	fmt.Println(fromFunction, "Starting sleep for ", sleepTime, "ms")
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	fmt.Println(fromFunction, "Waking up, slept for ", sleepTime, "ms")
	if ch != nil {
		ch <- sleepTime
	}
}

func sleepRandomContext(ctx context.Context, ch chan bool) {
	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()
	sleeptimeChan := make(chan int)
	go sleepRandom("sleepRandomContext", sleeptimeChan)
	select {
	case <-ctx.Done():
		fmt.Println("sleepRandomContext: Time to return")
	case sleepTime := <-sleeptimeChan:
		fmt.Println("Slept for ", sleepTime, "ms")
	}
}

func doWorkContext(ctx context.Context) {
	ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)
	defer func() {
		fmt.Println("doWorkContext complete")
		cancelFunction()
	}()

	ch := make(chan bool)
	go sleepRandomContext(ctxWithTimeout, ch)
	select {
	case <-ctx.Done():
		fmt.Println("doWorkContext: Time to return")
	case <-ch:
		fmt.Println("sleepRandomContext returned")
	}
}
