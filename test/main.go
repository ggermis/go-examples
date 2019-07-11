package main

type A struct {
	Name string
}

func (a *A) DoSomething() string {
	return "something"
}
