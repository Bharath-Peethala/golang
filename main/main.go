package main

import (
	"errors"
	"fmt"
	"golang/basics"
	"unicode/utf8"
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
	basics.Generics()
	fmt.Println(Reverse("bharath map"))
	fmt.Println(store)
}

func Reverse(s string) (string,error){
	fmt.Printf("input: %q\n", s)
	r := []rune(s)
	 fmt.Printf("runes: %q\n", r)
	if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
	for i,j:=0,len(r)-1;i < len(r)/2;i,j = i+1, j-1{
		r[i],r[j] = r[j],r[i]
	}
	return string(r),nil
}