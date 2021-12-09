package main

import "fmt"

type Person struct {
	// address
	StreetAddress, Postcode, City string

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
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

// PersonAddressBuilder methods
func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	b.person.Postcode = postcode
	return b
}

// PersonJobBuilder methods
func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName
	return b
}

func (b *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

func (b *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	b.person.AnnualIncome = annualIncome
	return b
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func main() {
	pb1 := NewPersonBuilder()
	pb1.
		Lives().
		At("Tverskaya, 17").
		In("Moscow").
		WithPostcode("121011").
		Works().
		At("Kaspersky Lab").
		AsA("Security Analyst").
		Earning(123000)

	pb2 := NewPersonBuilder()
	pb2.
		Lives().
		At("Naberezhnaya Moiki, 41").
		In("St.Petersburg").
		WithPostcode("177205").
		Works().
		At("ELG Prom").
		AsA("Programmer").
		Earning(178000)

	p1 := pb1.Build()
	p2 := pb2.Build()

	fmt.Printf("%+v\n", p1)
	fmt.Printf("%+v\n", p2)
}
