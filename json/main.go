package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         int    `json:"age"`
	hiddenField string `json:"not_shown"`
}

func main() {
	p := person{FirstName: "John", LastName: "Doe", Age: 42, hiddenField: "fields need to be exporterd to marshal"}
	fmt.Printf("%+v\n", p)

	// Marshal object -> JSON (byte array)
	b, _ := json.Marshal(p)
	fmt.Println(string(b))

	// Marshal object -> JSON (writer eg. http, stdout, network, ...)
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(p)

	// Unmarshal JSON -> object
	var people []person
	j := []byte(`[{"first_name": "Jane", "last_name": "Doe", "age": 37}]`)
	_ = json.Unmarshal(j, &people)
	fmt.Printf("%+v\n", people)

}
