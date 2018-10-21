package main

type Field struct {
  width, height int
  cells         []Cell
}

func FieldSize(field Field) int {
  return field.width * field.height
}

func NeighboursCount(field Field, cell Cell) int {
  return 8
}
