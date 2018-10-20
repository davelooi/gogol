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

  tables := []struct {
    neighbours int
    expected   Cell
  }{
    {0, Cell{Alive: false}},
    {1, Cell{Alive: false}},
    {2, Cell{Alive: false}},
    {3, Cell{Alive: true}},
    {4, Cell{Alive: false}},
    {5, Cell{Alive: false}},
    {6, Cell{Alive: false}},
    {7, Cell{Alive: false}},
    {8, Cell{Alive: false}},
  }

  for _, table := range tables {
    assert.Equal(t, table.expected, NextState(cell, table.neighbours))
  }
}

func TestLivingCellNextState(t *testing.T) {
  cell := Cell{Alive: true}
  tables := []struct {
    neighbours int
    expected   Cell
  }{
    {0, Cell{Alive: false}},
    {1, Cell{Alive: false}},
    {2, Cell{Alive: true}},
    {3, Cell{Alive: true}},
    {4, Cell{Alive: false}},
    {5, Cell{Alive: false}},
    {6, Cell{Alive: false}},
    {7, Cell{Alive: false}},
    {8, Cell{Alive: false}},
  }

  for _, table := range tables {
    assert.Equal(t, table.expected, NextState(cell, table.neighbours))
  }
}
