package main

type X struct {
	x int
}

type Y X

func NewX() *X {
	return &X{0}
}

func NewY() *Y {
	return &Y{10}
}

// Shared functionality which uses specialized method
func (this *X) Basic() {
	this.Specialized()
	println(this.x)
}

// Specialization for X
func (this *X) Specialized() {
	this.x += 1
}

// Specialization for Y
func (this *Y) Specialized() {
	this.x += 2
}

func main() {
	x := NewX()
	x.Basic()
	y := NewY()
	y.Basic() // Error
}
