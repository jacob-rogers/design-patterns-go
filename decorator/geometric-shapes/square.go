package main

import "fmt"

type Square struct {
	Size float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square wiith side %f", s.Size)
}

func (s *Square) Resize(factor float32) {
	s.Size *= factor
}
