// Go has some tail calls optimization:
//    http://stackoverflow.com/questions/12102675/tail-call-optimization-in-go
//    http://blog.gopheracademy.com/recursion
// Run and observ your memory!
// What about non tail recursions? Is there any limit?

package main

func rec(i int) int {
	if i%100000 == 0 {
		println(i)
	}
	x := rec(i + 1)
	println("a", rec(x+2))
	return i
}

func tail_rec(i int) {
	if i%100000 == 0 {
		println(i)
	}
	rec(i + 1)
}

func main() {
	//rec(1)
	tail_rec(1)
}
