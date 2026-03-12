// Positions are 0-1 where 0 is far left and 1 is far right

package grid

type Cell struct {
	x uint16
	y uint16
	collision []uint8
}

type RaycastResult struct {
	hit Vector2
	normal float32
}

func (cell *Cell) Raycast(
	rayX uint8, rayY uint8, rot float32,
) (bool, RaycastResult) {
	if !cell.HasCollider() { return false, RaycastResult{} }
	collider := cell.GetCollider()

	dx, dy := cos(rot), sin(rot)

	a := Vector2{x: float32(rayX), y: float32(rayY)}
	b := Vector2{x: dx * 255, y: dy * 255}
	r := b.Sub(a)

	var position Vector2
	var best float32 = maxFloat32
	hit := false

	for index, pos := range collider {
		if index % 4 != 0 { continue }

		startX, startY := pos, collider[index + 1]
		endX, endY := collider[index + 2], collider[index + 3]

		c := Vector2{x: float32(startX), y: float32(startY)}
		d := Vector2{x: float32(endX), y: float32(endY)}

		s := d.Sub(c)
		v := r.x * s.y - r.y * s.x

		u := ((c.x - a.x) * r.y - (c.y - a.y) * r.x) / v
		t := ((c.x - a.x) * s.y - (c.y - a.y) * s.x) / v

		// if a is to the left of s then get left side perpendicular vector of s
		// if a is to the right of s then get right side perpendicular vector of s

		if (t < best) {
			position = r.Mul(t).Add(a)
		}

		if (0 <= u && u <= 1 && 0 <= t && t <= 1) {
			hit = true
		}
	}

	return hit, RaycastResult{
		hit: position,
	}
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
