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

  testTable := []struct {
    in  Cell
    out []Cell
  }{
    {cell_00, []Cell{cell_01, cell_11, cell_10}},
    {cell_11, []Cell{cell_00, cell_01, cell_02, cell_10, cell_12, cell_20, cell_21, cell_22}},
    {cell_22, []Cell{cell_12, cell_11, cell_21}},
  }

  for _, tt := range testTable {
    assert.Equal(t, len(tt.out), len(Neighbours(field, tt.in.x, tt.in.y)))
  }
}

func TestNeighboursCount(t *testing.T) {
  cell_00 := Cell{x: 0, y: 0}
  cell_01 := Cell{x: 0, y: 1}
  cell_02 := Cell{x: 0, y: 2}
  cell_10 := Cell{x: 1, y: 0}
  cell_20 := Cell{x: 2, y: 0}
  cell_11 := Cell{x: 1, y: 1}
  cell_12 := Cell{x: 1, y: 2}
  cell_21 := Cell{x: 2, y: 1}
  cell_22 := Cell{x: 2, y: 2}
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

  testTable := []struct {
    in    Cell
    count int
  }{
    {cell_00, 3},
    {cell_01, 5},
    {cell_02, 3},
    {cell_10, 5},
    {cell_11, 8},
    {cell_12, 5},
    {cell_20, 3},
    {cell_21, 5},
    {cell_22, 3},
  }

  for _, tt := range testTable {
    assert.Equal(t, tt.count, NeighboursCount(field, tt.in))
  }
}
