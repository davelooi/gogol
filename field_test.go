package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestFieldSize(t *testing.T) {
  field := Field{width: 2, height: 3}
  assert.Equal(t, 6, FieldSize(field))
}

func TestCellNeighbours(t *testing.T) {
  cells := []Cell{
    Cell{Alive: false, x: 0, y: 0},
    Cell{Alive: false, x: 0, y: 1},
    Cell{Alive: false, x: 1, y: 0},
    Cell{Alive: false, x: 1, y: 1},
  }
  field := Field{width: 2, height: 2, cells: cells}
  assert.Equal(t, 6, NeighboursCount(field, cells[0]))
}
