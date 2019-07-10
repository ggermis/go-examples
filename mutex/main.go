package main

import (
	"log"
	"sync"
)

type safeCounter struct {
	// protect v by mutex m
	m *sync.Mutex
	v int
}

func (c *safeCounter) increment() {
	c.m.Lock()
	defer c.m.Unlock()
	c.v++
}

// alias an existing type to make the variable naming more consistent
type unsafeCounter = int

var wg *sync.WaitGroup

var safe_counter *safeCounter
var unsafe_counter unsafeCounter

func doSomething() {
	defer wg.Done()
	safe_counter.increment()
	unsafe_counter++
}

func main() {
	wg = new(sync.WaitGroup)

	safe_counter = &safeCounter{m: &sync.Mutex{}, v: 0}
	unsafe_counter = 0

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go doSomething()
	}

	wg.Wait()
	log.Printf("Counters: [safe: %d, unsafe: %d]", safe_counter.v, unsafe_counter)
}
