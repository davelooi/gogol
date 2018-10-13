package rules

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestCellAlive(t *testing.T) {
  livingCell := Cell{Alive: true}
  assert.Equal(t, true, IsAlive(livingCell))

  deadCell := Cell{Alive: false}
  assert.Equal(t, false, IsAlive(deadCell))
}

func TestDeadCellNextState(t *testing.T) {
  cell := Cell{Alive: false}
  m := map[int]Cell{
    0: {Alive: false},
    1: {Alive: false},
    2: {Alive: false},
    3: {Alive: true},
    4: {Alive: false},
    5: {Alive: false},
    6: {Alive: false},
    7: {Alive: false},
    8: {Alive: false},
  }

  for neighbours, expected := range m {
    assert.Equal(t, expected, NextState(cell, neighbours))
  }
}

func TestLivingCellNextState(t *testing.T) {
  cell := Cell{Alive: true}
  m := map[int]Cell{
    0: {Alive: false},
    1: {Alive: false},
    2: {Alive: true},
    3: {Alive: true},
    4: {Alive: false},
    5: {Alive: false},
    6: {Alive: false},
    7: {Alive: false},
    8: {Alive: false},
  }

  for neighbours, expected := range m {
    assert.Equal(t, expected, NextState(cell, neighbours))
  }
}
