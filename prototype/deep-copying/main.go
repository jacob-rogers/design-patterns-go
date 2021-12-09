package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

// We need this receiver function with copies of pointer struct's data
func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

// DeepCopy function - here we also cover all need field by its DeepCopy method, or copy()
// built-in function, like for [*person.Friends] slice
func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)

	return &q
}

func main() {
	john := Person{
		"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Matt"},
	}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
