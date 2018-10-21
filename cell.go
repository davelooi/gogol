package main

type Cell struct {
  x, y  int // position of a cell
  Alive bool
}

func IsAlive(c Cell) bool {
  return c.Alive
}

func NextState(c Cell, neighbours int) Cell {
  if IsAlive(c) {
    if neighbours == 2 || neighbours == 3 {
      return Cell{x: c.x, y: c.y, Alive: true}
    } else {
      return Cell{x: c.x, y: c.y, Alive: false}
    }
  } else {
    if neighbours == 3 {
      return Cell{x: c.x, y: c.y, Alive: true}
    } else {
      return Cell{x: c.x, y: c.y, Alive: false}
    }
  }
}
