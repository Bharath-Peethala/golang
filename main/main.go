package main

import (
	"fmt"
	"golang/basics"
)

func main() {
	store := make(map[string]string)
	i := 0
	for i < 5 {
		numerator := basics.GenerateRandomNumber()
		denominator := basics.GenerateRandomNumber()
		store[fmt.Sprintf("%v,%v", numerator, denominator)] = ""
		i++
	}
	store = basics.StoreResults(store)
	fmt.Println(store)
}
