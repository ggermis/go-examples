package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go func() {
		c <- "test"
	}()
	fmt.Println(<-c)
}
