package main

import "fmt"

// Bridge

// Circle, Square
// Raster, Vector

// RasterCircle, VectorCircle, RasterSquare, VectorSquare ...

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
	//
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for a circle of radius", radius)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius += factor
}

func main() {
	rr := RasterRenderer{}
	vr := VectorRenderer{}

	circle1 := NewCircle(&rr, 5)
	circle1.Draw()
	circle1.Resize(2)
	circle1.Draw()

	circle2 := NewCircle(&vr, 5)
	circle2.Draw()
	circle2.Resize(4)
	circle2.Draw()

}
