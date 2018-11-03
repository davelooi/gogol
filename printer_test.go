package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestAllAlive(t *testing.T) {
  cell_00 := Cell{x: 0, y: 0, Alive: true}
  cell_01 := Cell{x: 0, y: 1, Alive: true}
  cell_02 := Cell{x: 0, y: 2, Alive: true}
  cell_10 := Cell{x: 1, y: 0, Alive: true}
  cell_20 := Cell{x: 2, y: 0, Alive: true}
  cell_11 := Cell{x: 1, y: 1, Alive: true}
  cell_12 := Cell{x: 1, y: 2, Alive: true}
  cell_21 := Cell{x: 2, y: 1, Alive: true}
  cell_22 := Cell{x: 2, y: 2, Alive: true}
  cells := []Cell{
    cell_00,
    cell_01,
    cell_02,
    cell_10,
    cell_20,
    cell_11,
    cell_12,
    cell_21,
    cell_22,
  }
  field := Field{width: 3, height: 3, cells: cells}
  out := "xxx\nxxx\nxxx\n"

  assert.Equal(t, out, Output(field))
}

func TestAllDead(t *testing.T) {
  cell_00 := Cell{x: 0, y: 0, Alive: false}
  cell_01 := Cell{x: 0, y: 1, Alive: false}
  cell_02 := Cell{x: 0, y: 2, Alive: false}
  cell_10 := Cell{x: 1, y: 0, Alive: false}
  cell_20 := Cell{x: 2, y: 0, Alive: false}
  cell_11 := Cell{x: 1, y: 1, Alive: false}
  cell_12 := Cell{x: 1, y: 2, Alive: false}
  cell_21 := Cell{x: 2, y: 1, Alive: false}
  cell_22 := Cell{x: 2, y: 2, Alive: false}
  cells := []Cell{
    cell_00,
    cell_01,
    cell_02,
    cell_10,
    cell_20,
    cell_11,
    cell_12,
    cell_21,
    cell_22,
  }
  field := Field{width: 3, height: 3, cells: cells}
  out := "   \n   \n   \n"

  assert.Equal(t, out, Output(field))
}

func TestHalfAlive(t *testing.T) {
  cell_00 := Cell{x: 0, y: 0, Alive: true}
  cell_01 := Cell{x: 0, y: 1, Alive: false}
  cell_02 := Cell{x: 0, y: 2, Alive: true}
  cell_10 := Cell{x: 1, y: 0, Alive: false}
  cell_11 := Cell{x: 1, y: 1, Alive: false}
  cell_12 := Cell{x: 1, y: 2, Alive: true}
  cell_20 := Cell{x: 2, y: 0, Alive: true}
  cell_21 := Cell{x: 2, y: 1, Alive: true}
  cell_22 := Cell{x: 2, y: 2, Alive: false}
  cells := []Cell{
    cell_00,
    cell_01,
    cell_02,
    cell_10,
    cell_20,
    cell_11,
    cell_12,
    cell_21,
    cell_22,
  }
  field := Field{width: 3, height: 3, cells: cells}
  out := "x x\n  x\nxx \n"

  assert.Equal(t, out, Output(field))
}
