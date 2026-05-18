package grid

import (
	"potato-bones/src/utils/math"
)

type Cell struct {
	x uint16
	y uint16
	collision []uint8
}

type CellRaycastResult struct {
	hit math.Vector2
	normal float32
}

func (cell *Cell) Raycast(
	rayX uint8, rayY uint8,
	rot float32, dist float32,
) (bool, CellRaycastResult) {
	if !cell.HasCollider() { return false, CellRaycastResult{} }
	collider := cell.GetCollider()

	directionX, directionY := math.Cos(rot), math.Sin(rot)

	rayStart := math.NewVector2(float32(rayX) - directionX * math.Epsilon, float32(rayY) - directionY * math.Epsilon)
	rayDelta := math.NewVector2(directionX, directionY).Mul(dist)

	var position math.Vector2
	var normal float32 = 0
	var best float32 = 2
	var hit bool = false

	for index, pos := range collider {
		if index % 2 != 0 { continue }
		if index > len(collider) - 4 { break }

		startX, startY := pos, collider[index + 1]
		endX, endY := collider[index + 2], collider[index + 3]

		segmentStart := math.NewVector2(float32(startX), float32(startY))
		segmentEnd := math.NewVector2(float32(endX), float32(endY))

		segmentDelta := segmentEnd.Sub(segmentStart)

		scalarFactor := rayDelta.Cross(segmentDelta)
		if math.Abs(scalarFactor) < math.Epsilon { continue } // Parallel
		rayStartToSegmentStart := segmentStart.Sub(rayStart)

		segmentScalar := rayStartToSegmentStart.Cross(rayDelta) / scalarFactor
		rayScalar := rayStartToSegmentStart.Cross(segmentDelta) / scalarFactor

		if (segmentScalar > 0 && segmentScalar <= 1 && rayScalar > 0 && rayScalar <= 1) {
			hit = true

			if (rayScalar < best) {
				best = rayScalar

				// also equal to rayDelta.Mul(rayScalar).Add(rayStart)
				position = segmentDelta.Mul(segmentScalar).Add(segmentStart)

				normalVector := math.NewVector2(-segmentDelta.Y, segmentDelta.X)

				if normalVector.Dot(rayDelta) > 0 {
					normalVector = normalVector.Inverse()
				}

				normal = math.Atan2(normalVector.Y, normalVector.X)
			}
		}
	}

	if (normal < 0) {
		normal += 2 * math.Pi;
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
