package main

import (
	"log"
)

// These classes can be provided from external modules. As long as they have
// methods with the same signature, we can create an interface with those method
// signatures that is meaningful for our client code

type classA struct{}

func (a *classA) Do(m string) {
	log.Println("A:", m)
}

type classB struct{}

func (a *classB) Do(m string) {
	log.Println("B:", m)
}

// create an interface that has meaning for this client

type myDoer interface {
	Do(m string)
}

func doSomething(d myDoer) {
	d.Do("something")
}

func main() {
	doSomething(&classA{})
	doSomething(&classB{})
}
