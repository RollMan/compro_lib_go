package main

import (
  "fmt"
  "./compro_lib"
)

func main(){
  compro_lib.Init()
  var words []string
  for i := 0; i < 4; i++ {
    w, _ := compro_lib.NextWord()
    words = append(words, w)
  }
  fmt.Printf("%s\n", words)

  var integers []int
  for i := 0; i < 3; i++ {
    v, _ := compro_lib.NextInt()
    integers = append(integers, v)
  }
  fmt.Printf("%v\n", integers)

  var doubles []float64
  for i := 0; i < 3; i++ {
    v, _ := compro_lib.NextDouble()
    doubles = append(doubles, v)
  }
  fmt.Printf("%v\n", doubles)

}
