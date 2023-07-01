package main

import (
	"fmt"
	"log"
	"runtime"
)

func Switch() {
	log.Println("Go runs on")
	x := runtime.NumCPU()
	fmt.Println(runtime.GOARCH, x)
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Printf("OS %s", os)
	case "linux":
		fmt.Printf("OS %s", os)
	default:
		fmt.Printf("OS %s", os)
	}
}
