package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// simulate work
func work(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(".. starting work", i)
	duration := time.Duration(rand.Intn(5)) * time.Microsecond
	time.Sleep(duration)
	fmt.Println(".. finished work", i, "duration", duration)
}

// start all work in parallel
func orchestrate(done chan<- bool) {
	var wg sync.WaitGroup
	for i := 1; i <= 25; i++ {
		wg.Add(1)
		go work(i, &wg)
	}
	wg.Wait()
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("Begin")

	go orchestrate(done)

	// block until written to
	<-done
	fmt.Println("End")
}
