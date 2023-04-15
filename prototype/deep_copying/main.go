package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	John := Person{"John", &Address{"123", "London", "UK"}}

	// Cannot do this cause Address is pointer if you modify variable of John got
	// Change too
	// Jane := John
	// Jane.Name = "Jane"
	// Jane.Address.City = "Manchester"

	// deep copying

	Jane := John
	Jane.Address = &Address{
		John.Address.StreetAddress,
		John.Address.City,
		John.Address.Country,
	}

	Jane.Name = "Jane"
	Jane.Address.City = "Manchester"

	fmt.Println(Jane, Jane.Address)
	fmt.Println(John, John.Address)
}
