package main

import (
	"fmt"
	"golang/basics"
	"log"
)

func main() {
	result, err := basics.DivideTwo(5, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result of DivideTwo: %v \n", result)
}
