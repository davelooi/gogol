package main

import (
  "fmt"
  "math/rand"
)

func main() {
  width := 60
  height := 60

  cells := make([][]Cell, 0)
  for h := 0; h < height; h++ {
    row := make([]Cell, 0)
    for w := 0; w < width; w++ {
      row = append(row, Cell{x: w, y: h, Alive: RandomBool()})
    }
    cells = append(cells, row)
  }
  field := Field{width: width, height: height, cells: cells}

  fmt.Print(Output(field))
}

func RandomBool() bool {
  if rand.Intn(100) < 10 {
    return true
  }
  return false
}
