// Solution by Nate Finch

package main

type X struct {
	x int
}

func (x *X) Log() {
	x.x -= 1
	x.print()
}

func (x *X) print() {
	println(x.x)
}

type Y struct {
	X
}

// Y's Log method hides X's, so if you call Y.Log, it'll call this instead
func (y *Y) Log() {
	y.x += 2
	y.print()
}

// Initializers can return both final type or interface - Logger

func NewX() *X {
	return &X{0}
}

func NewY() *Y {
	return &Y{X{10}}
}

func main() {
	x := NewX()
	x.Log()
	y := NewY()
	y.Log()
}

// this is the interface both *X and *Y fulfill
// objects that need to log could take a Logger
// and not need to know what is behind it
type Logger interface {
	Log()
}
