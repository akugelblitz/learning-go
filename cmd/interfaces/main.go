package main

import "fmt"

func main(){
  result := add(1.2, 2)
  fmt.Printf("%v\n", result)
  result1 := add(29, 3)
  fmt.Printf("%v\n", result1)
  res := add("hi ", "there")
  fmt.Printf("%v\n", res)
}



func add[T int|float64|string](a, b T) T{
  return a+b

}
