package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

func (this *Address) DeepCopy() *Address {
	return &Address{
		this.StreetAddress,
		this.City,
		this.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (this *Person) DeepCopy() *Person {
	p := *this
	p.Address = p.Address.DeepCopy()
	copy(p.Friends, this.Friends)
	return &p
}

func main() {
	John := Person{"John", &Address{"123", "London", "UK"}, []string{"Chris", "Jane"}}

	Jane := John.DeepCopy()
	Jane.Name = "Jane"
	Jane.Address.City = "Manchester"
	Jane.Friends = append(Jane.Friends, "John")

	fmt.Println(Jane, Jane.Address)
	fmt.Println(John, John.Address)
}
