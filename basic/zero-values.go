package main

import "fmt"

func zeroValue() {
	var i int
	var f float64
	var b bool
	var s string

	/**
	 * General format.
	 *
	 */
	fmt.Printf("%v, %v, %v, %q\n", i, f, b, s)
}
