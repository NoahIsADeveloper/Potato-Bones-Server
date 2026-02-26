package grid

type Cell struct {
	x uint16
	y uint16
	collision []uint8
}

//TODO: Use "Cross Product Method" or whatever to do precise raycasts.
func (cell *Cell) Raycast(
	rayX, rayY float64,
	dirX, dirY float64,
) bool {
	return cell.HasCollider()
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
