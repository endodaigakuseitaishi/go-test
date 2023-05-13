package main

import (
	"fmt"
	"go-sample/alib"
)


func IsOne(n int) bool {
	if n == 1 {
    return true
  }
  return false
}

func main() {
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(0))

	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(alib.Average(s))
}
