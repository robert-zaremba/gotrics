package main

type X struct {
	x int
}

type Y struct {
	X
}

func NewX() *X {
	return &X{0}
}

func NewY() *Y {
	return &Y{X{10}}
}

// Shared functionality which uses specialized method
func (this *X) Basic() {
	this.Specialized()
	println(this.x)
}

// Specialization for X
func (this *X) Specialized() {
	this.x -= 1
}

// Specialization for Y
func (this *Y) Specialized() {
	this.x += 2
}

func main() {
	x := NewX()
	x.Basic()
	y := NewY()
	y.Basic()   // prints 9 instead of 9
	y.X.Basic() // this statement show what is really called by in the previous line
	// The problem is that y.Basic == y.X.Basic which doesn't know anything about y and y.Specialized
}
