package main

type Field struct {
  width, height int
  cells         []Cell
}

func FieldSize(field Field) int {
  return field.width * field.height
}

func NeighboursCount(field Field, cell Cell) int {
  return len(Neighbours(field, cell))
}

func Neighbours(field Field, cell Cell) []Cell {
  var neigbours []Cell
  x := cell.x
  y := cell.y
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

func FindCell(field Field, x int, y int) Cell {
  for i := range field.cells {
    if field.cells[i].x == x && field.cells[i].y == y {
      return field.cells[i]
    }
  }
  // SHOULD NOT HAPPEN
  return field.cells[0]
}
