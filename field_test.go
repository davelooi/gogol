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
  cell_00 := Cell{x: 0, y: 0}
  cell_01 := Cell{x: 0, y: 1}
  cell_02 := Cell{x: 0, y: 2}
  cell_10 := Cell{x: 1, y: 0}
  cell_20 := Cell{x: 2, y: 0}
  cell_11 := Cell{x: 1, y: 1}
  cell_12 := Cell{x: 1, y: 2}
  cell_21 := Cell{x: 2, y: 1}
  cell_22 := Cell{x: 2, y: 2}
  cells := [][]Cell{
    {cell_00, cell_01, cell_02},
    {cell_10, cell_11, cell_12},
    {cell_20, cell_21, cell_22},
  }
  field := Field{width: 3, height: 3, cells: cells}

  testTable := []struct {
    in  Cell
    out []Cell
  }{
    {cell_00, []Cell{cell_01, cell_10, cell_11}},
    {cell_11, []Cell{cell_00, cell_01, cell_02, cell_10, cell_12, cell_20, cell_21, cell_22}},
    {cell_22, []Cell{cell_11, cell_12, cell_21}},
  }

  for _, tt := range testTable {
    assert.Equal(t, tt.out, Neighbours(field, tt.in))
  }
}

func TestGetCell(t *testing.T) {
  cell_00 := Cell{x: 0, y: 0, Alive: true}
  cell_01 := Cell{x: 0, y: 1, Alive: false}
  cell_02 := Cell{x: 0, y: 2, Alive: true}
  cell_10 := Cell{x: 1, y: 0, Alive: false}
  cell_20 := Cell{x: 2, y: 0, Alive: true}
  cell_11 := Cell{x: 1, y: 1, Alive: false}
  cell_12 := Cell{x: 1, y: 2, Alive: true}
  cell_21 := Cell{x: 2, y: 1, Alive: true}
  cell_22 := Cell{x: 2, y: 2, Alive: false}
  cells := [][]Cell{
    {cell_00, cell_01, cell_02},
    {cell_10, cell_11, cell_12},
    {cell_20, cell_21, cell_22},
  }
  field := Field{width: 3, height: 3, cells: cells}

  testTable := []struct {
    x   int
    y   int
    out Cell
  }{
    {0, 0, cell_00},
    {0, 1, cell_01},
    {0, 2, cell_02},
    {1, 0, cell_10},
    {1, 1, cell_11},
    {1, 2, cell_12},
    {2, 0, cell_20},
    {2, 1, cell_21},
    {2, 2, cell_22},
  }

  for _, tt := range testTable {
    assert.Equal(t, tt.out, GetCell(field, tt.x, tt.y))
  }
}
