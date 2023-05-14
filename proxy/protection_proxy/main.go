package main

import "fmt"

type Driven interface {
	Drive()
}

type Car struct{}

func (this *Car) Drive() {
	fmt.Println("Car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car    Car
	driver *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 18 {
		c.car.Drive()
	} else {
		fmt.Println("Driver Too Young")
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{Car{}, driver}
}

func main() {
	car := NewCarProxy(&Driver{18})
	car.Drive()
}
