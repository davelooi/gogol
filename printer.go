package main

import "bytes"

func Output(field Field) string {
  var b bytes.Buffer

  for h := 0; h < field.height; h++ {
    for w := 0; w < field.width; w++ {
      b.WriteString(FmtCell(FindCell(field, w, h)))

    }
    b.WriteString("\n")
  }
  return b.String()
}

func FmtCell(cell Cell) string {
  if cell.Alive == true {
    return "x"
  }
  return " "
}
