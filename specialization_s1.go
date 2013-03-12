package main

type X struct {
	x           int
	specialized func()
}

type Y struct {
	X
}

func NewX() *X {
	x := &X{x: 0}
	x.specialized = func() {
		x.x -= 1
	}
	return x
}

func NewY() *Y {
	y := new(Y)
	y.x = 10
	y.specialized = func() {
		y.x += 2
	}
	return y
}

// Shared functionality which uses specialized method
func (this *X) Basic() {
	this.specialized()
	println(this.x)
}

func main() {
	x := NewX()
	x.Basic()
	y := NewY()
	y.Basic()
}
