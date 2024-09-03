package basics

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

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
