package grid

import (
	"potato-bones/src/utils"
	"sync"
)

type Grid struct {
	cells [][]Cell
	cellSize uint8
	rotation float32
	mutex sync.RWMutex
	width uint16
	height uint16
}

func (grid *Grid) GetCell(x uint16, y uint16) *Cell {
	grid.mutex.RLock(); defer grid.mutex.RUnlock()
	return &grid.cells[x][y]
}

// If dist is -1 then assume raycast is infinite
func (grid *Grid) Raycast(
	posX float32, posY float32,
	rot float32, dist float32,
) (bool, *Cell) {
	grid.mutex.RLock(); defer grid.mutex.RUnlock()

	px, py := posX, posY
	dx, dy := utils.Cos(rot), utils.Sin(rot)
	sx, sy := utils.Abs(1 / dx), utils.Abs(1 / dy)

	steps := 0

	for px >= 0 && py >= 0 && px < float32(grid.width - 1) && py < float32(grid.height - 1) {
		if sx < sy {
			px += dx
		} else {
			py += dy
		}

		cell := grid.GetCell(uint16(px), uint16(py))
		if cell.HasCollider() {
			return true, cell
		}

		steps++
	}

	return false, NewCell(0, 0)
}

func Sign(num float32) float32 {
	if num > 0 {
		return 1
	} else if num < 0 {
		return -1
	} else {
		return 0
	}
}

func generateCells(sizeX uint16, sizeY uint16) [][]Cell {
	cells := make([][]Cell, sizeX)

	var x uint16
	for x = range sizeX {
		cells[x] = make([]Cell, sizeY)
        var y uint16
		for y = range sizeY {
			cells[x][y] = *NewCell(x, y)
		}
    }

	return cells
}

func NewGrid(sizeX uint16, sizeY uint16) *Grid {
	grid := &Grid{
		cells: generateCells(sizeX, sizeY),
		cellSize: 50,
		width: sizeX,
		height: sizeY,
	}

	return grid
}
