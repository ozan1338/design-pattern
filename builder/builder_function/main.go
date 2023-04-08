package main

import "fmt"

type Person struct {
	name, position string
}

type personModification func(*Person)
type PersonBuilder struct {
	actions []personModification
}

func (this *PersonBuilder) Called(name string) *PersonBuilder {
	this.actions = append(this.actions, func(p *Person) {
		p.name = name
	})
	return this
}

func (this *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	this.actions = append(this.actions, func(p *Person) {
		p.position = position
	})
	return this
}

func (this *PersonBuilder) Build() *Person {
	p := Person{}
	for _, action := range this.actions {
		action(&p)
	}
	return &p
}

func main() {
	b := PersonBuilder{}
	p := b.Called("Dmiti").WorksAsA("developer").Build()
	fmt.Println(*p)
}
