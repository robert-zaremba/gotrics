package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func test(m interface{}) {
	jsonStream := `{"Name": "Andy", "Age": 22}`

	mt := reflect.TypeOf(m)
	fmt.Println("type of m:", mt, mt.Kind())
	if mt.Kind() != reflect.Ptr {
		fmt.Println("m not a pointer, pointer type of m is:", reflect.TypeOf(&m))
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	var err error
	for {
		if mt.Kind() == reflect.Ptr {
			err = dec.Decode(m)
		} else {
			err = dec.Decode(&m)
		}
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reflect.TypeOf(m))
		fmt.Printf("%v\n", m)
	}
}

func main() {
	var x interface{}
	var y interface{} = Person{}

	// use pointer to better introspect a type
	fmt.Println(reflect.TypeOf(x))  // nil?! that's not right either!
	fmt.Println(reflect.TypeOf(y))  // note that y's type is really interface{}!
	fmt.Println(reflect.TypeOf(&x)) // ahh thats better
	fmt.Println(reflect.TypeOf(&y)) // ahh thats better

	println("-------------------")
	test(new(Person))
	println("-------------------")
	test(Person{}) // Decode doesn't recognize Person if we convert interface to pointer.
	// also take a look about using factory method to not always decode in the same place:
	//   http://play.golang.org/p/qVaOGHib8r
}
