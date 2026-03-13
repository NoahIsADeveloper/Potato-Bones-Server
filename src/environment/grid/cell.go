package grid

type Cell struct {
	x uint16
	y uint16
	collision []uint8
}

type CellRaycastResult struct {
	hit Vector2
	normal float32
}

func (cell *Cell) Raycast(
	rayX uint8, rayY uint8,
	rot float32, dist float32,
) (bool, CellRaycastResult) {
	if !cell.HasCollider() { return false, CellRaycastResult{} }
	collider := cell.GetCollider()

	dx, dy := cos(rot), sin(rot)

	rayStart := Vector2{x: float32(rayX) - dx * epsilon, y: float32(rayY) - dy * epsilon}
	rayDirection := Vector2{x: dx, y: dy}.Mul(dist)

	var position Vector2
	var normal float32 = 0
	var best float32 = 2
	var hit bool = false

	for index, pos := range collider {
		if index % 2 != 0 { continue }
		if index > len(collider) - 4 { break }

		startX, startY := pos, collider[index + 1]
		endX, endY := collider[index + 2], collider[index + 3]

		segmentStart := Vector2{x: float32(startX), y: float32(startY)}
		segmentEnd := Vector2{x: float32(endX), y: float32(endY)}

		segmentDirection := segmentEnd.Sub(segmentStart)

		v := rayDirection.Cross(segmentDirection)
		if abs(v) < epsilon { continue }
		j := segmentStart.Sub(rayStart)

		segmentScalar := j.Cross(rayDirection) / v
		rayScalar := j.Cross(segmentDirection) / v

		if (segmentScalar > 0 && segmentScalar <= 1 && rayScalar > 0 && rayScalar <= 1) {
			hit = true

			if (rayScalar < best) {
				best = rayScalar

				// position = rayDirection.Mul(rayScalar).Add(rayStart)
				position = segmentDirection.Mul(segmentScalar).Add(segmentStart)

				// TODO: simplify
				nx := -segmentDirection.y
				ny := segmentDirection.x

				if nx * rayDirection.x + ny * rayDirection.y > 0 {
					nx = -nx
					ny = -ny
				}

				normal = atan2(ny, nx)
			}
		}
	}

	if (normal < 0) {
		normal += 2 * pi;
	}

	return hit, CellRaycastResult{
		hit: position,
		normal: normal,
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
