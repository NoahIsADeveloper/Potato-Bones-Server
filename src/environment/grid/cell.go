// Positions are 0-1 where 0 is far left and 1 is far right

package grid

import "fmt"

type Cell struct {
	x uint16
	y uint16
	collision []uint8
}

type RaycastResult struct {
	hitX uint8
	hitY uint8
	normal float32
}

func (cell *Cell) Raycast(
	rayX uint8, rayY uint8,
	rot float32, dist uint8,
) (bool, RaycastResult) {
	if !cell.HasCollider() { return false, RaycastResult{} }
	collider := cell.GetCollider()

	for index, pos := range collider {
		if index % 4 != 0 { continue }

		startX, startY := pos / 255, collider[index + 1] / 255
		endX, endY := collider[index + 2] / 255, collider[index + 3] / 255

		fmt.Printf("Line 1: (%d, %d)\nLine 2: (%d, %d)\n", startX, startY, endX, endY)
	}

	return true, RaycastResult{}
}

func (cell *Cell) HasCollider() bool {
	return len(cell.collision) > 0
}

func (cell *Cell) GetCollider() []uint8 {
	return cell.collision
}

func (cell *Cell) AddCollisionPoint(x uint8, y uint8) {
	cell.collision = append(cell.collision, x, y)
}

func (cell *Cell) GetPosition() (uint16, uint16) {
	return cell.x, cell.y
}

func NewCell(x uint16, y uint16) *Cell {
	cell := &Cell{
		x: x,
		y: y,
	}

	return cell
}
