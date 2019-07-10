package main

//go:generate protoc --go_out=. ./addressbook/addressbook.proto

import (
	pb "codenut.org/lab/go/examples/protobuf/addressbook"
	"github.com/golang/protobuf/proto"
	"log"
)

// To add a method to an existing type, we need to wrap it
type AddressBook struct {
	pb.AddressBook
}

func (b *AddressBook) addPerson(person *pb.Person) {
	b.People = append(b.People, person)
}

func main() {
	// initialize addressbook using wrapper type
	book := &AddressBook{}
	book.addPerson(&pb.Person{Name: "Person A", Email: "person.a@example.com"})
	book.addPerson(&pb.Person{Name: "Person B", Email: "person.b@example.com"})

	// Marshal
	data, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	log.Printf("%v\n", book)

	// Raw data
	log.Printf("%v\n", data)

	// Unmarshal
	new_book := &pb.AddressBook{}
	if err := proto.Unmarshal(data, new_book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	log.Printf("%v\n", new_book)
}
