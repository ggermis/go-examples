package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.be/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.Request)
	fmt.Printf("%+v\n", resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", body)
}
