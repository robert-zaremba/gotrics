package main

type Dependency interface {
	Specialized(*int)
}

// Box is a Wrapper / Container
type Box struct {
	Dependency
	x int
}

func (this *Box) Basic() {
	this.Specialized(&this.x)
	println(this.x)
}

type X struct{}
type Y struct{}

func (this X) Specialized(a *int) {
	*a -= 1
}

func (this Y) Specialized(a *int) {
	*a += 2
}

func main() {
	x := Box{X{}, 0}
	x.Basic()
	y := Box{Y{}, 10}
	y.Basic()
}
