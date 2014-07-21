/* This package presents how Go handles structure passing to the functions
 */

package main

import "fmt"

const LEN uint = 1e7

type X struct {
	pole int
	arr  [LEN]int
}

func (*X) String() string {
	return "big string"
}

func main() {
	fmt.Println("start")
	// var arr [LEN]int
	// fmt.Println(len(arr))
	var x = X{pole: 1}
	x.pole = 2
	sum := 0
	x.arr[2131] = 213
	for _, v := range x.arr {
		sum += v
	}
	fmt.Println(x.arr[2131], sum)
	//fmt.Printf("%v\n", &x)
	// fmt.Printf("type b: %T \n", sl)
}
