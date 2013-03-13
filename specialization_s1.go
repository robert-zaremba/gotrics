package main

type X struct {
	x           int
	specialized func()
}

type Y struct {
	X
}

func NewX() *X {
	x := &X{0, nil}
	x.specialized = func() {
		x.x -= 1
	}
	return x
}

// This can also be implemented as a function which returns *X
// So we won't need Y struct definition
// But it is here to show how different structs can handle common Basic method
func NewY() *Y {
	// y := new(Y)
	// y.x = 10
	y := &Y{X{10, nil}}
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
