package grid

type Cell struct {
	x uint16
	y uint16
	collision []int8
}

func (cell *Cell) Raycast(
	rayX, rayY float64,
	dirX, dirY float64,
) bool {
	return cell.HasCollider()
}

func (cell *Cell) HasCollider() bool {
	return len(cell.collision) > 0
}

func (cell *Cell) GetCollision() []int8 {
	return cell.collision
}

func (cell *Cell) GetPosition() (uint16, uint16) {
	return cell.x, cell.y
}

func NewCell() *Cell {
	cell := &Cell{}

	return cell
}
