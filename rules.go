package rules

type Cell struct {
  Position string // placeholder
  Alive    bool
}

func IsAlive(c Cell) bool {
  return c.Alive
}

func NextState(c Cell, neighbours int) Cell {
  if IsAlive(c) {
    if neighbours == 2 || neighbours == 3 {
      return Cell{Position: c.Position, Alive: true}
    } else {
      return Cell{Position: c.Position, Alive: false}
    }
  } else {
    if neighbours == 3 {
      return Cell{Position: c.Position, Alive: true}
    } else {
      return Cell{Position: c.Position, Alive: false}
    }
  }
}
