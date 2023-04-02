package main

// interface segregation principle
// it means you should not throw everything into one single interface

type Document struct{}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrint struct{}

func (m MultiFunctionPrint) Print(d Document) {}
func (m MultiFunctionPrint) Fax(d Document)   {}
func (m MultiFunctionPrint) Scan(d Document)  {}

type OldFashionedPrinter struct{}

func (o OldFashionedPrinter) Print(d Document) {}

//Deprecated: ..
func (o OldFashionedPrinter) Fax(d Document) {
	panic("operartion not supported")
}

//Deprecated: ..
func (o OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// ISP Interface segregation principle
// basically states that try to break up an interface into separate
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct{}

func (mp MyPrinter) Print(d Document) {}

type Photocopier struct{}

func (p Photocopier) Print(d Document) {}

func (p Photocopier) Scan(d Document) {}

// Or you can define interface that hold another interface
type MultiFunctionDevice interface {
	Printer
	Scanner
	// Fax
}

// decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (this MultiFunctionMachine) Print(d Document) {
	this.printer.Print(d)
}

func (this MultiFunctionMachine) Scan(d Document) {
	this.scanner.Scan(d)
}

func main() {

}
