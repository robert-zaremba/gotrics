/* by Nate Finch
 * The solution is to refactor the method flow order.
 */

package main

type X struct {
	x int
}

func (x *X) Specialized() {
	x.x -= 1
	x.Basic()
}

func (x *X) Basic() {
	println(x.x)
}

type Y struct {
	X
}

// Y's Log method hides X's, so if you call Y.Log, it'll call this instead
func (y *Y) Specialized() {
	y.x += 2
	y.Basic()
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
	x.Specialized()
	y := NewY()
	y.Specialized()
}

// this is the interface both *X and *Y fulfill
// objects that need to log could take a Logger
// and not need to know what is behind it
type Logger interface {
	Log()
}
