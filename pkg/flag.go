package main

import (
	"flag"
	"fmt"
)

func Flags() {
	var firstName = flag.String("first-name", "Muhamad Andre", "put your first name")
	var age = flag.Int("age", 19, "The age of the person")

	// use parse()
	flag.Parse()

	fmt.Printf("my name is %s\n dan umur saya adalah %v", *firstName, *age)
	// output
	/**
	  * my name is Muhamad Andre
	    dan umur saya adalah 19
	*/
}

func FlagsVar() {
	var username string
	var age int

	flag.StringVar(&username, "my username", "muhamadandre", "example username")
	flag.IntVar(&age, "my age", 20, "example age")
	flag.Parse()
	fmt.Printf("my name is %s\n dan umur saya adalah %v", username, age)
}
