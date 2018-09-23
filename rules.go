package rules

type CellState int

const (
  UnderPopulation CellState = iota
  Living
  OverPopulation
  Reproduction
  Dead
)

type Cell struct {
  NumberOfNeighbours int
  Position           string // placeholder
  Alive              bool
}

// Given a cell, determines what is it's current state
func LivingCondition(c Cell) CellState {
  neighbours := c.NumberOfNeighbours
  switch {
  case 0 <= neighbours && neighbours <= 1:
    return UnderPopulation
  case 2 <= neighbours && neighbours <= 3:
    return Living
  case 4 <= neighbours:
    return OverPopulation
  }
  return Living
}

func DeadCondition(c Cell) CellState {
  if c.NumberOfNeighbours == 3 {
    return Reproduction
  }
  return Dead
}
