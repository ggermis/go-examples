package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg *sync.WaitGroup

type workerConfig struct {
	action string
}

type workerResult struct {
	config   workerConfig
	status   string
	duration time.Duration
}

func doWork(id int, in <-chan workerConfig, out chan<- workerResult) {
	defer close(out)
	log.Println("starting worker", id)
	for config := range in {
		log.Println("worker", id, "=> starting work for", config.action)

		// simulate work
		duration := time.Duration(rand.Intn(1000)+200) * time.Millisecond
		time.Sleep(duration)

		log.Println("worker", id, "=> posting result for", config.action)
		out <- workerResult{config: config, status: "ok", duration: duration}

		log.Println("worker", id, "=> done")
	}
}

func handleResult(in <-chan workerResult) {
	defer wg.Done()
	result := <-in
	log.Println("-> result for", result.config.action, "->", result.status, ":", result.duration)
}

func main() {
	log.Println("Starting...")

	rand.Seed(time.Now().UnixNano())

	configs := []workerConfig{
		{action: "a"},
		{action: "b"},
		{action: "c"},
		{action: "d"},
		{action: "e"},
		{action: "f"},
		{action: "g"},
		{action: "h"},
		{action: "i"},
		{action: "j"},
	}

	// channel to read work from
	in := make(chan workerConfig)
	defer close(in)

	// channel to process results from
	out := make(chan workerResult)

	// start worker threads
	for i := 1; i <= 7; i++ {
		go doWork(i, in, out)
	}

	// perform work + handle response
	wg = new(sync.WaitGroup)
	for _, config := range configs {
		wg.Add(1)
		in <- config
		time.Sleep(time.Duration(rand.Intn(100)+50) * time.Millisecond)
		go handleResult(out)
	}

	wg.Wait()
	log.Println("Done.")
}
