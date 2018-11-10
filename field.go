package main

type Field struct {
  width, height int
  cells         [][]Cell
}

func FieldSize(field Field) int {
  return field.width * field.height
}

func Neighbours(field Field, cell Cell) []Cell {
  var neigbours []Cell

  for x := cell.x - 1; x <= cell.x+1; x++ {
    if x >= 0 && x < field.width {
      for y := cell.y - 1; y <= cell.y+1; y++ {
        if y >= 0 && y < field.height {
          if x == cell.x && y == cell.y {
            continue
          }
          neigbours = append(neigbours, field.cells[x][y])
        }
      }
    }
  }

  return neigbours
}

func GetCell(field Field, x int, y int) Cell {
  return field.cells[x][y]
}
