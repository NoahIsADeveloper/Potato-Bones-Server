package grid

import "math"

var pi float32 = 3.14159265358979323846
var twoPi float32 = 2 * pi
var halfPi float32 = pi / 2
var epsilon float32 = 0.00001

var maxFloat32 float32 = math.MaxFloat32

type Vector2 struct {
	x float32
	y float32
}

func (a Vector2) Sub(b Vector2) (Vector2) {
	return Vector2{
		x: a.x - b.x,
		y: a.y - b.y,
	}
}

func (a Vector2) Add(b Vector2) (Vector2) {
	return Vector2{
		x: a.x + b.x,
		y: a.y + b.y,
	}
}

func (a Vector2) Mul(v float32) (Vector2) {
	return Vector2{
		x: a.x * v,
		y: a.y * v,
	}
}

func (a Vector2) Div(v float32) (Vector2) {
	return Vector2{
		x: a.x / v,
		y: a.y / v,
	}
}

func (a Vector2) Dot(b Vector2) (float32) {
	return a.x * b.x + a.y * b.y
}

func abs(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}

func atan2(y float32, x float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}

func sign(x float32) float32 {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func normalizeAngle(x float32) float32 {
	for x > pi {
		x -= twoPi
	}
	for x < -pi {
		x += twoPi
	}
	return x
}

func sin(x float32) float32 {
	x = normalizeAngle(x)

	x2 := x * x
	x3 := x2 * x
	x5 := x3 * x2
	x7 := x5 * x2

	return x - x3/6 + x5/120 - x7/5040
}

func cos(x float32) float32 {
	x = normalizeAngle(x)

	x2 := x * x
	x4 := x2 * x2
	x6 := x4 * x2

	return 1 - x2/2 + x4/24 - x6/720
}

func tan(x float32) float32 {
	c := cos(x)

	if abs(c) < 0.00001 {
		return sign(sin(x)) * 1e9
	}

	return sin(x) / c
}

// rounding
func floor(x float32) float32 {
	i := int(x)

	if float32(i) > x {
		return float32(i - 1)
	}

	return float32(i)
}

func ceil(x float32) float32 {
	i := int(x)

	if float32(i) < x {
		return float32(i + 1)
	}

	return float32(i)
}

func round(x float32) float32 {
	if x >= 0 {
		return floor(x + 0.5)
	}
	return ceil(x - 0.5)
}

func toRad(rot uint16) float32 {
	rot = rot % 360

	n := float32(rot) / 360
	n *= twoPi

	return n
}

func toDeg(rot float32) uint16 {
	if rot < 0 { rot = -rot }
	rot = mod(rot, twoPi)

	n := rot / twoPi
	n *= 360

	return uint16(n)
}

func mod(a float32, b float32) float32 {
	for a >= b {
		a = a - b
	}

	return a
}