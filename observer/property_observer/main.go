package main

import (
	"container/list"
	"fmt"
)

// Observer = monitoring for change aka doctor and Observable = patient

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

type PropertyChange struct {
	Name  string
	Value interface{}
}

type Person struct {
	Observable
	age int
}

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{
			new(list.List),
		},
		age: age,
	}
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	if p.age == age {
		return
	}

	p.age = age
	p.Fire(PropertyChange{"Age", p.age})
}

type TraficManagement struct {
	o Observable
}

func (t *TraficManagement) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 16 {
			fmt.Println("Congrats, you can drive now")
			t.o.Unsubscribe(t)
		}
	}
}

func main() {
	p := NewPerson(14)
	t := &TraficManagement{p.Observable}

	p.Subscribe(t)

	for i := 15; i <= 20; i++ {
		fmt.Println("Setting the age to ", i)
		p.SetAge(i)
	}
}
