// This package presents performance of closures compared to normal functions.
// As a consequence we will see if it's worth to keep a contextual definitions.
// of closures in function body.
// run with:
//    * go run closure.go  # to present a speed ration (compared to non closure version)
//    * go test closure_test.go -bench . # to run test benchmark mode
package main

import (
	"fmt"
	"testing"
	"time"
)

var N = 1000000000

func runOut(a, b int) int { return a + b }

func BenchmarkOutside(*testing.B) {
	var sum = 0
	for i := 0; i < N; i++ {
		sum = runOut(sum, i)
	}
	fmt.Println(sum)
}

func BenchmarkClosure(*testing.B) {
	run := func(a, b int) int { return a + b }
	var sum = 0
	for i := 0; i < N; i++ {
		sum = run(sum, i)
	}
	fmt.Println(sum)
}

func BenchmarkClosureIn(*testing.B) {
	var sum = 0
	for i := 0; i < N; i++ {
		sum = func(i int) int { return sum + i }(i)
	}
	fmt.Println(sum)
}

func do() {
	t := time.Now().UnixNano()
	BenchmarkOutside(nil)
	tend := time.Now().UnixNano()
	base := float64(tend - t)
	fmt.Println("BenchmarkOutside")

	t = time.Now().UnixNano()
	BenchmarkClosure(nil)
	tend = time.Now().UnixNano()
	fmt.Println("BenchmarkClosure  ", float64(tend-t)/base)

	t = time.Now().UnixNano()
	BenchmarkClosureIn(nil)
	tend = time.Now().UnixNano()
	fmt.Println("BenchmarkClosureIn", float64(tend-t)/base)
}

func main() {
	do()
}
