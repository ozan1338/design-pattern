package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// function factory Approach
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// another struct approach
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func NewEmployeeFactory2(postion string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{Position: postion, AnnualIncome: annualIncome}
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func main() {
	developerFactory := NewEmployeeFactory("dev", 750000)
	managerFactory := NewEmployeeFactory("manager", 100000)

	developer := developerFactory("adam")
	manager := managerFactory("james")

	fmt.Println(developer)
	fmt.Println(manager)

	// the only advantage using this approach we can later change like annualIncome
	bossFactory := NewEmployeeFactory2("CEO", 20000)
	boss := bossFactory.Create("Same")
	fmt.Println(boss)
}
