package main

import (
	"errors"
	"fmt"
)

func main() {
	newfunc("sup")
}

func newfunc(printvalue string) {
	var quotient, remainder, err = divide(21, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(printvalue, quotient, remainder)
	}
}

func divide(inp int, divisor int) (int, int, error) {
	var err error
	if divisor == 0 {
		err = errors.New("cannot divide by 0")
		return 0, 0, err
	}
	return inp / 5, inp % 5, err
}
