package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hellworld")
	principal := 10000.0
	rate := 5.5
	time := 10.0

	result := principal * math.Pow((100.0+rate)/100, time)
	fmt.Println(result)
}
