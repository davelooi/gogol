package main

import (
  "fmt"
  "math/rand"
  "os"
  "strconv"
)

func main() {
  if len(os.Args) < 4 {
    PrintHelp()
    return
  }

  width := ParseUnlessError(os.Args[1])
  height := ParseUnlessError(os.Args[2])
  seed := ParseUnlessError(os.Args[3])

  if width == -1 || height == -1 || seed == -1 {
    return
  }

  cells := make([][]Cell, 0)
  for h := 0; h < height; h++ {
    row := make([]Cell, 0)
    for w := 0; w < width; w++ {
      row = append(row, Cell{x: w, y: h, Alive: RandomBool(seed)})
    }
    cells = append(cells, row)
  }
  field := Field{width: width, height: height, cells: cells}

  fmt.Print(Output(field))
}

func RandomBool(seed int) bool {
  if rand.Intn(100) < seed {
    return true
  }
  return false
}

func ParseUnlessError(args string) int {
  arg, err := strconv.Atoi(args)
  if err != nil {
    fmt.Printf("%s\n", err)
    return -1
  }
  return arg
}

func PrintHelp() {
  helpStatement := `
  Usage:
    ./gogol [width] [height] [seed]
  `
  fmt.Print(helpStatement)
}
