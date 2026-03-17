package math

import "math"

var Pi float32 = 3.14159265358979323846
var TwoPi float32 = 2 * Pi
var HalfPi float32 = Pi / 2
var Epsilon float32 = 0.00001

var MaxFloat32 float32 = math.MaxFloat32

func Abs(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}

func Atan2(y float32, x float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}

func Sign(x float32) float32 {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func NormalizeAngle(x float32) float32 {
	for x > Pi {
		x -= TwoPi
	}
	for x < -Pi {
		x += TwoPi
	}
	return x
}

func Sin(x float32) float32 {
	x = NormalizeAngle(x)

	return float32(math.Sin(float64(x)))
}

func Cos(x float32) float32 {
	x = NormalizeAngle(x)

	return float32(math.Cos(float64(x)))
}

func Tan(x float32) float32 {
	c := Cos(x)

	if Abs(c) < 0.00001 {
		return Sign(Sin(x)) * 1e9
	}

	return Sin(x) / c
}

func Floor(x float32) float32 {
	i := int(x)

	if float32(i) > x {
		return float32(i - 1)
	}

	return float32(i)
}

func Ceil(x float32) float32 {
	i := int(x)

	if float32(i) < x {
		return float32(i + 1)
	}

	return float32(i)
}

func Round(x float32) float32 {
	if x >= 0 {
		return Floor(x + 0.5)
	}
	return Ceil(x - 0.5)
}

func ToRad(rot uint16) float32 {
	rot = rot % 360

	n := float32(rot) / 360
	n *= TwoPi

	return n
}

func ToDeg(rot float32) uint16 {
	if rot < 0 { rot = -rot }
	rot = Mod(rot, TwoPi)

	n := rot / TwoPi
	n *= 360

	return uint16(n)
}

func Mod(a float32, b float32) float32 {
	for a >= b {
		a = a - b
	}

	return a
}