package main

import (
	"fmt"
)

func main() {
	// Allow setting multiple messages on the channel from same routine
	c := make(chan string, 3)

	c <- "a"
	c <- "b"
	c <- "c"

	fmt.Println(<-c)
	fmt.Println(<-c)
}
