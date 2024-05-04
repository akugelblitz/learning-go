package main

import "fmt"

func main(){
  intArr := [...]int32{1,2,3}
  fmt.Println((intArr))

  var intSlice []int = []int{1,2,3}
  fmt.Println(intSlice, len(intSlice), cap(intSlice))

  intSlice = append(intSlice, 7)
  fmt.Println(intSlice, len(intSlice), cap(intSlice))

  var intSlice2 []int = []int{8, 9, 10, 11} 
  intSlice2 = append(intSlice, intSlice2...)


  var myMap map[string]uint = make(map[string]uint)
  fmt.Println(myMap)

  var myMap2 map[string]uint = map[string]uint{"Adam": 23, "jason": 24}

   fmt.Println(myMap2["Adam"])
 fmt.Println(myMap2["Jason"])
}
