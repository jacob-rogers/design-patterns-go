package main

import (
	"fmt"
)

// OCP
// open for extension, but closed for modification
// Specification

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct{}

func (f *Filter) FilterByColors(products []Product, colors []Color) []*Product {
	result := make([]*Product, 0)

	for i, p := range products {
		for _, c := range colors {
			if p.color == c {
				result = append(result, &products[i])
			}
		}
	}

	return result
}

func (f *Filter) FilterBySizes(products []Product, sizes []Size) []*Product {
	result := make([]*Product, 0)

	for i, p := range products {
		for _, s := range sizes {
			if p.size == s {
				result = append(result, &products[i])
			}
		}
	}

	return result
}

func (f *Filter) FilterByColorsAndSizes(products []Product, colors []Color, sizes []Size) []*Product {
	result := make([]*Product, 0)

	for i, p := range products {
		for _, c := range colors {
			for _, s := range sizes {
				if p.color == c && p.size == s {
					result = append(result, &products[i])
				}
			}
		}
	}

	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (cs ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == cs.color
}

type SizeSpecification struct {
	size Size
}

func (ss SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == ss.size
}

type BetterFilter struct{}

func (bf *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, p := range products {
		if spec.IsSatisfied(&p) {
			result = append(result, &products[i])
		}
	}

	return result
}

type Param struct {
	color Color
	size  Size
}

type SpecificationSet struct {
	specs []Specification
}

func (sset SpecificationSet) IsSatisfied(p *Product) bool {
	for _, spec := range sset.specs {
		if !spec.IsSatisfied(p) {
			return false
		}
	}

	return true
}

func main() {
	apple := Product{"Apple", green, small}
	house := Product{"House", blue, large}
	tree := Product{"Tree", green, large}

	products := []Product{apple, house, tree}

	// Old usage of directly changed object (filtering method were added)
	fmt.Printf("Green products (old):\n")
	f := Filter{}
	for _, p := range f.FilterByColors(products, []Color{green}) {
		fmt.Printf(" - %s is green\n", p.name)
	}

	fmt.Printf("\n##########################################################\n\n")

	fmt.Printf("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	bf1 := BetterFilter{}
	for _, p := range bf1.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", p.name)
	}

	fmt.Printf("\n##########################################################\n\n")

	largeSpec := SizeSpecification{large}
	lgSpec := SpecificationSet{[]Specification{largeSpec, greenSpec}}
	bf2 := BetterFilter{}
	fmt.Printf("Green and large products (new):\n")
	for _, p := range bf2.Filter(products, lgSpec) {
		fmt.Printf(" - %s is green and large\n", p.name)
	}
}
