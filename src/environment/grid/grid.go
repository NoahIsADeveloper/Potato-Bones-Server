package grid

import (
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

func (grid *Grid) Raycast(
	posX float32, posY float32,
	rot float32, dist float32,
) (bool, *Cell) {
	grid.mutex.RLock(); defer grid.mutex.RUnlock()

	px, py := posX, posY
	dx, dy := cos(rot), sin(rot)

	if abs(dx) < epsilon { dx = epsilon }
	if abs(dy) < epsilon { dy = epsilon }

	sx := abs(1 / dx)
	sy := abs(1 / dy)

	stepX := sign(dx)
	stepY := sign(dy)

	var tx, ty float32
	if dx > 0 {
		tx = (ceil(px) - px) * sx
	} else {
		tx = (px - floor(px)) * sx
	}

	if dy > 0 {
		ty = (ceil(py) - py) * sy
	} else {
		ty = (py - floor(py)) * sy
	}

	var totalDist float32 = 0

	for px >= 0 && py >= 0 && px < float32(grid.width) && py < float32(grid.height) {
		cellX := uint16(floor(px))
		cellY := uint16(floor(py))

		if cellX < grid.width && cellY < grid.height {
			cell := &grid.cells[cellX][cellY]
			if cell.HasCollider() {
				return true, cell
			}
		}

		if dist > 0 && totalDist >= dist {
			break
		}

		if tx < ty {
			px += stepX
			totalDist += sx
			tx += sx
		} else {
			py += stepY
			totalDist += sy
			ty += sy
		}
	}

	return false, NewCell(0, 0)
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
