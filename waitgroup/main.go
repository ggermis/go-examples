package main

import (
	"log"
	"sync"
)

var wg sync.WaitGroup

func doSomething(i int) {
	defer wg.Done()
	log.Printf(" .. doing %2d\n", i)
}

func main() {
	log.Println("Staring...")
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go doSomething(i)
	}
	wg.Wait()
	log.Println("Finished")
}
