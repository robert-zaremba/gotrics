/* this approach resambles cookie pattern
 * it's like s1 but instead creating closures we implement specialized function ahead
 *   and set them as a parameters in initializators
 */

package main

type X struct {
	x           int
	specialized func(*int)
}

func specializedX(a *int) {
	*a -= 1
}

func specializedY(a *int) {
	*a += 2
}

func NewX() X {
	return X{0, specializedX}
}

func NewY() X {
	return X{10, specializedY}
}

// Shared functionality which uses specialized method
func (this *X) Basic() {
	this.specialized(&this.x)
	println(this.x)
}

func main() {
	x := NewX()
	x.Basic()
	y := NewY()
	y.Basic()
}
