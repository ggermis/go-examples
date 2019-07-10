package main

import (
	"log"
	"sync"
	"sync/atomic"
)

// Counter that will be incremented using a safe operation (sync.atomic)
var safe_counter int32 = 0

// Counter that will simply be incremented by addition
var unsafe_counter int32 = 0

func doSomething(wg *sync.WaitGroup) {
	defer wg.Done()
	unsafe_counter = unsafe_counter + 1
	atomic.AddInt32(&safe_counter, 1)
}

func main() {
	log.Printf("INIT: [unsafe: %5d, safe: %5d]\n", unsafe_counter, safe_counter)

	wg := new(sync.WaitGroup)
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go doSomething(wg)
	}
	wg.Wait()

	log.Printf("END : [unsafe: %5d, safe: %5d]\n", unsafe_counter, safe_counter)
}
