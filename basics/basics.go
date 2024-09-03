package basics

import (
	"errors"
	"fmt"
)

func HelloWorld() string {
	return fmt.Sprintf("Hello World!")
}

func DivideTwo(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("denomiator is 0")
	}
	return a / b, nil
}
