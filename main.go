package main

import (
  "fmt"
  "math/rand"
)

func main() {
  width := 99
  height := 99
  var cells = []Cell{}

  for w := 0; w < width; w++ {
    for h := 0; h < height; h++ {
      cells = append(cells, Cell{x: w, y: h, Alive: RandomBool()})
    }
  }

  field := Field{width: width, height: width, cells: cells}

  fmt.Print(Output(field))
}

func RandomBool() bool {
  if rand.Intn(100) < 10 {
    return true
  }
  return false
}
