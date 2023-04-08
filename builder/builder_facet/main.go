package main

import "fmt"

type Person struct {
	// address
	StreetAddress, PostCode, City string

	// job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (this *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	this.person.StreetAddress = streetAddress
	return this
}

func (this *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	this.person.City = city
	return this
}

func (this *PersonAddressBuilder) PostalCode(postalCode string) *PersonAddressBuilder {
	this.person.PostCode = postalCode
	return this
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (this *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	this.person.CompanyName = companyName
	return this
}

func (this *PersonJobBuilder) Position(position string) *PersonJobBuilder {
	this.person.Position = position
	return this
}

func (this *PersonJobBuilder) Earning(earning int) *PersonJobBuilder {
	this.person.AnnualIncome = earning
	return this
}

func (this *PersonBuilder) Build() *Person {
	return this.person
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().At("Tebet Timur 4A").PostalCode("28028").In("South Jakarta").
		Works().At("Dans").Position("Programmer").Earning(750000)

	person := pb.Build()
	fmt.Println(*person)
}
