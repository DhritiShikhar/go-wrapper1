// NOT COMPLETE --> Needs work

// Wrapping arbitrary function in go

package main

import (
	"fmt"
	"reflect"
)

type (
	// the function which is to be wrapped
	func1 func(int, int) int

	// the function which is going to wrap
	wrapper func(func1, int, int) int
)

// the actual wrapper
func genericWrapper(in []reflect.Value) []reflect.Value {

	// do Something

	return in[0].Call(in[1:])
}

// creates a wrapper function
func createWrapper(w interface{}) {
	w1 := reflect.ValueOf(w)

	fmt.Println(w1)
}

func main() {
	var wrap1 wrapper

	fmt.Println(wrap1)
	createWrapper(wrap1)
}
