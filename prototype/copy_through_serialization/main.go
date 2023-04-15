package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (this *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(this)

	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)

	result := Person{}

	_ = d.Decode(&result)

	return &result
}

func main() {
	John := Person{"John", &Address{"123", "London", "UK"}, []string{"Chris", "Jane"}}

	Jane := John.DeepCopy()
	Jane.Name = "Jane"
	Jane.Address.City = "Manchester"
	Jane.Friends = []string{"John"}

	fmt.Println(Jane, Jane.Address)
	fmt.Println(John, John.Address)
}
