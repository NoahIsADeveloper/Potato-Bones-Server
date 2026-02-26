package utils

import "math"

func Sin(theta float32) float32 {
	return float32(math.Sin(float64(theta)))
}

func Cos(theta float32) float32 {
	return float32(math.Cos(float64(theta)))
}

func Tan(theta float32) float32 {
	return float32(math.Tan(float64(theta)))
}

func Sign(num float32) float32 {
	if num > 0 {
		return 1
	} else if num < 0 {
		return -1
	} else {
		return 0
	}
}

func Abs(num float32) float32 {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}

func Floor(num float32) float32 {
	return float32(math.Floor(float64(num)))
}

func Ceil(num float32) float32 {
	return float32(math.Ceil(float64(num)))
}

func Round(num float32) float32 {
	return float32(math.Round(float64(num)))
}

var Pi float32 = 3.141592653589