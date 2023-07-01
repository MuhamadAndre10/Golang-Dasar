package main

import (
	"fmt"
	"os"
)

func OS() {
	args := os.Args
	hostName, _ := os.Hostname()

	fmt.Println(args)
	fmt.Println(hostName)

	// get arguments
	// fmt.Println(args[1])
	// fmt.Println(args[2])
	//! error jika tidak ada argument.
}
