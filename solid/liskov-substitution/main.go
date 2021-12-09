package main

import "fmt"

// Liskov Substitution principle
// You should be able to substitute the embedded type
// instead of embedded part

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (sq *Square) SetWidth(width int) {
	sq.width = width
	sq.height = width
}

func (sq *Square) SetHeight(height int) {
	sq.width = height
	sq.height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)

	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected area of ", expectedArea, ", but got of ", actualArea, "\n")
}

// Solution #1
type Square2 struct {
	size int
}

func (sq *Square2) Rectangle() Rectangle {
	return Rectangle{sq.size, sq.size}
}

// If some of type properties are similar or identical in their behaviour, they
// shouldn't be affected independenly. Object should have idempotent behaviour
// in any instances of it, no matter how we create it or modify its fields or methods

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(2)
	UseIt(sq)
}
