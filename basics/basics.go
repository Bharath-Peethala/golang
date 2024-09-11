package basics

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Number interface{
	int64 | float64
}

func HelloWorld() string {
	return fmt.Sprintf("Hello World!")
}

func DivideTwo(a int64, b int64) (float64, error) {
	if b == 0 {
		return 0, errors.New("denominator is 0")
	}
	return float64(a) / float64(b), nil
}

func GenerateRandomNumber() int64 {
	return rand.Int63n(100)
}

func StoreResults(store map[string]string) map[string]string {
	for numsStr, _ := range store {
		nums := strings.Split(numsStr, ",")
		numerator, _ := strconv.ParseInt(nums[0], 10, 64)
		denominator, _ := strconv.ParseInt(nums[1], 10, 64)
		result, _ := DivideTwo(numerator, denominator)
		store[numsStr] = fmt.Sprintf("%.2f", result)
	}
	return store
}

func Generics(){
	ints := map[string]int64{
		"num1":1,
		"num2":2,
	}
	floats := map[string]float64{
		"num1":1.2,
		"num2":2.3,
	}
	
	numsArr := map[int]float64{
		1:1.2,
		2:2.3,
	}

	fmt.Println(SumIntOrFloats(ints))
	fmt.Println(SumIntOrFloats(floats))
	fmt.Println(SumIntOrFloats(numsArr))
}

func SumIntOrFloats[K comparable, V Number](numsMap map[K]V) V{
	var sum V
	for _, num := range numsMap{
		sum += num
	}
	return sum
}