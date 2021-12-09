package main

// Interface Segregation principle (ISP)
// Don't put too much into interface -
// split it into multiple interfaces
// YAGNI - You Are not Going to Need It

type Document struct{}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (mfp MultiFunctionPrinter) Print(d Document) {

}

func (mfp MultiFunctionPrinter) Fax(d Document) {

}

func (mfp MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionedPrinter struct{}

func (ofp OldFashionedPrinter) Print(d Document) {
	// OK
}

// Deprecated: METHOD NOT SUPPORTED BY TYPE
func (ofp OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Deprecated: METHOD NOT SUPPORTED BY TYPE
func (ofp OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// ISP

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct{}

func (mp MyPrinter) Print(d Document) {

}

type Photocopier struct{}

func (p Photocopier) Scan(d Document) {
	panic("implement me")
}

func (p Photocopier) Print(d Document) {
	panic("implement me")
}

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

func (mfm MultiFunctionMachine) Print(d Document) {
	mfm.printer.Print(d)
	mfm.scanner.Scan(d)
}

func main() {
	// ofp := OldFashionedPrinter{}
	// ofp.Scan()
}
