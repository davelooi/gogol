package rules

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestLivingCondition(t *testing.T) {
  m := map[int]CellState{
    0: UnderPopulation,
    1: UnderPopulation,
    2: Living,
    3: Living,
    4: OverPopulation,
    5: OverPopulation,
    6: OverPopulation,
    7: OverPopulation,
    8: OverPopulation,
  }

  for k, v := range m {
    c := Cell{
      NumberOfNeighbours: k,
      Position:           "somewhere",
      Alive:              true,
    }
    assert.Equal(t, v, LivingCondition(c))
  }
}

func TestDeadCondition(t *testing.T) {
  m := map[int]CellState{
    0: Dead,
    1: Dead,
    2: Dead,
    3: Reproduction,
    4: Dead,
    5: Dead,
    6: Dead,
    7: Dead,
    8: Dead,
  }

  for k, v := range m {
    c := Cell{
      NumberOfNeighbours: k,
      Position:           "somewhere",
      Alive:              false,
    }
    assert.Equal(t, v, DeadCondition(c))
  }
}
