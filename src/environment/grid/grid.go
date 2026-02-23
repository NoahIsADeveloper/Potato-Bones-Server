package grid

import (
	"math"
	"sync"
)

type Grid struct {
	cells [][]Cell
	rotation float32
	mutex sync.RWMutex
}

func (grid *Grid) GetCell(x uint16, y uint16) Cell {
	grid.mutex.RLock(); defer grid.mutex.RUnlock()
	return grid.cells[x][y]
}

func (grid *Grid) Raycast(
	posX float64, posY float64,
	dirX float64, dirY float64,
) Cell {
	grid.mutex.RLock()
	defer grid.mutex.RUnlock()

	mapX := int(posX)
	mapY := int(posY)

	if dirX == 0 {
		dirX = 1e-6
	}
	if dirY == 0 {
		dirY = 1e-6
	}

	deltaDistX := math.Abs(1 / dirX)
	deltaDistY := math.Abs(1 / dirY)

	var stepX, stepY int
	var sideDistX, sideDistY float64

	if dirX < 0 {
		stepX = -1
		sideDistX = (posX - float64(mapX)) * deltaDistX
	} else {
		stepX = 1
		sideDistX = (float64(mapX+1) - posX) * deltaDistX
	}

	if dirY < 0 {
		stepY = -1
		sideDistY = (posY - float64(mapY)) * deltaDistY
	} else {
		stepY = 1
		sideDistY = (float64(mapY+1) - posY) * deltaDistY
	}

	sizeX := len(grid.cells)
	sizeY := len(grid.cells[0])
	var side int // 0 = x, 1 = y

	for {
		if sideDistX < sideDistY {
			mapX += stepX
			sideDistX += deltaDistX
			side = 0
		} else {
			mapY += stepY
			sideDistY += deltaDistY
			side = 1
		}

		if mapX < 0 || mapX >= sizeX || mapY < 0 || mapY >= sizeY {
			return Cell{}
		}

		cell := grid.cells[mapX][mapY]

		var hitX, hitY float64
		if side == 0 {
			hitX = float64(mapX)
			hitY = posY + dirY*(sideDistX-deltaDistX)
		} else {
			hitX = posX + dirX*(sideDistY-deltaDistY)
			hitY = float64(mapY)
		}

		if cell.Raycast(hitX, hitY, dirX, dirY) {
			return cell
		}
	}
}

func generateCells(sizeX uint16, sizeY uint16) [][]Cell {
	cells := make([][]Cell, sizeX)

	var x uint16
	for x = range sizeX {
		cells[x] = make([]Cell, sizeY)
        var y uint16
		for y = range sizeY {
			cells[x][y] = *NewCell()
		}
    }

	return cells
}

func NewGrid(sizeX uint16, sizeY uint16) *Grid {
	grid := &Grid{
		cells: generateCells(sizeX, sizeY),
	}

	return grid
}
