package main

type Field struct {
  width, height int
  cells         []Cell
}

func FieldSize(field Field) int {
  return field.width * field.height
}

func NeighboursCount(field Field, cell Cell) int {
  return len(Neighbours(field, cell.x, cell.y))
}

func Neighbours(field Field, x int, y int) []Cell {
  var neigbours []Cell
  for _, cell := range field.cells {
    if cell.x == x+1 && cell.y == y {
      neigbours = append(neigbours, cell)
    } else if cell.x == x+1 && cell.y == y+1 {
      neigbours = append(neigbours, cell)
    } else if cell.x == x+1 && cell.y == y-1 {
      neigbours = append(neigbours, cell)
    } else if cell.x == x && cell.y == y+1 {
      neigbours = append(neigbours, cell)
    } else if cell.x == x && cell.y == y-1 {
      neigbours = append(neigbours, cell)
    } else if cell.x == x-1 && cell.y == y {
      neigbours = append(neigbours, cell)
    } else if cell.x == x-1 && cell.y == y+1 {
      neigbours = append(neigbours, cell)
    } else if cell.x == x-1 && cell.y == y-1 {
      neigbours = append(neigbours, cell)
    }
  }
  return neigbours
}
