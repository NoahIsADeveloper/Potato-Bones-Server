package grid

import (
	"fmt"
	"potato-bones/src/globals"
	"potato-bones/src/utils/files"
	"potato-bones/src/utils/math"
	"strings"
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

type GridRaycastResult struct {
	cell *Cell
	normal float32
	position math.Vector2
}

func (grid *Grid) GetCell(x uint16, y uint16) *Cell {
	grid.mutex.RLock(); defer grid.mutex.RUnlock()
	return &grid.cells[x][y]
}

func (grid *Grid) Raycast(
	posX float32, posY float32,
	rot float32, dist float32,
) (bool, GridRaycastResult) {
	grid.mutex.RLock(); defer grid.mutex.RUnlock()

	px, py := posX, posY
	dx, dy := math.Cos(rot), math.Sin(rot)

	if math.Abs(dx) < math.Epsilon { dx = math.Epsilon }
	if math.Abs(dy) < math.Epsilon { dy = math.Epsilon }

	sx := math.Abs(1 / dx)
	sy := math.Abs(1 / dy)

	stepX := math.Sign(dx)
	stepY := math.Sign(dy)

	var tx, ty float32
	if dx > 0 {
		tx = (math.Ceil(px) - px) * sx
	} else {
		tx = (px - math.Floor(px)) * sx
	}

	if dy > 0 {
		ty = (math.Ceil(py) - py) * sy
	} else {
		ty = (py - math.Floor(py)) * sy
	}

	var totalDist float32 = 0

	for px >= 0 && py >= 0 && px < float32(grid.width) && py < float32(grid.height) {
		cellX := uint16(math.Floor(px))
		cellY := uint16(math.Floor(py))

		if cellX < grid.width && cellY < grid.height {
			cell := &grid.cells[cellX][cellY]

			var localX, localY uint8
			if tx < ty {
				if dx > 0 {
					localX = 0
				} else {
					localX = 255
				}
				localY = uint8(math.Mod(py, 1) * 255)
			} else {
				localX = uint8(math.Mod(px, 1) * 255)
				if dy > 0 {
					localY = 0
				} else {
					localY = 255
				}
			}

			var cellDist float32
			if dist < 0 {
				cellDist = 10000
			} else {
				remainingDist := dist - totalDist
				if remainingDist <= 0 {
					break
				}
				cellDist = remainingDist * 255
			}
			success, raycastResult := cell.Raycast(localX, localY, rot, cellDist)
			if success {
				return true, GridRaycastResult{
					cell: cell,
					normal: raycastResult.normal,
					position: math.NewVector2(
						raycastResult.hit.X / 255 + float32(cellX),
						raycastResult.hit.Y / 255 + float32(cellY),
					),
				}
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

	return false, GridRaycastResult{}
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

func LoadGridFromMap(name string) (*Grid, error) {
	mapPath := *globals.MapPath + name + ".pbmap"
	if !files.FileExists(mapPath) {
		return nil, fmt.Errorf("map file does not exist: %s", mapPath)
	}

	data, err := files.ReadFile(mapPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read map file: %w", err)
	}

	mapData := strings.Split(string(data), "\n")

	// PASS 1: Get dependencies
	for index, line := range mapData {
		if !strings.HasSuffix(line, ";") && !strings.HasSuffix(line, ":") {
			if index == 1 {
				if line != globals.Version {
					var errString = fmt.Sprintf("map version mismatch: got %s, expected %s", line, globals.Version)

					if *globals.ErrorOnVersionMismatch {
						return nil, fmt.Errorf("%s", errString)
					} else {
						fmt.Printf("%s", errString)
					}
				}
			}
			if index == 2 { continue }
			if line == "" { continue }
			if strings.HasSuffix(line, ":") { continue }
			if strings.HasSuffix(line, ";") { continue }

			dependencyPath := *globals.TilesetPath + line + ".pbtile"
			if !files.FileExists(dependencyPath) {
				return nil, fmt.Errorf("tileset %s does not exist: %s", line, dependencyPath)
			}
			tilesetData, err := files.ReadFile(dependencyPath)
			if err != nil {
				return nil, fmt.Errorf("failed to read tileset file: %w", err)
			}
			tilesetLines := strings.Split(string(tilesetData), "\n")

			// parse tileset
		}
	}

	grid := NewGrid(100, 100)

	return grid, nil
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
