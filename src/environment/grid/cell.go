// Positions are 0-1 where 0 is far left and 1 is far right

package grid

import "fmt"

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
	r := Vector2{x: dx * 255, y: dy * 255}

	var position Vector2
	var normal float32 = 0
	var best float32 = 2
	var hit bool = false

	// TODO: Better variable names
	// TODO: Fix skipping over collision borders
	for index, pos := range collider {
		if index % 2 != 0 { continue }
		if index > len(collider) - 4 { break }

		startX, startY := pos, collider[index + 1]
		endX, endY := collider[index + 2], collider[index + 3]

		fmt.Printf("%d, %d / %d, %d\n", startX, startY, endX, endY)

		c := Vector2{x: float32(startX), y: float32(startY)}
		d := Vector2{x: float32(endX), y: float32(endY)}

		s := d.Sub(c)
		v := r.Cross(s)

		j := c.Sub(a)

		u := j.Cross(r) / v
		t := j.Cross(s) / v

		fmt.Println(t, best, r)
		if (u > 0 && u <= 1 && t > 0 && t <= 1) {
			hit = true

			if (t < best) {

				best = t

				// should be equal to s.Mul(u).Add(c)
				position = r.Mul(t).Add(a)

				if (s.Dot(a) < 0) {
					normal = atan2(-s.x, s.y)
				} else {
					normal = atan2(s.x, -s.y)
				}
			}
		}
	}

	if (normal < 0) {
		normal += 2 * pi;
	}

	fmt.Println("done", hit)
	return hit, RaycastResult{
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
