package math

import (
	"testing"
)

func TestDegRad(t *testing.T) {
	t.Helper()

	var targetDeg uint16 = 180
	rad := ToRad(targetDeg)
	resultDeg := ToDeg(rad)

	if resultDeg != targetDeg {
		t.Errorf("wanted %d got %d", targetDeg, resultDeg)
	}

	if rad != Pi {
		t.Errorf("wanted %f got %f", Pi, HalfPi)
	}
}

func TestAbs(t *testing.T) {
	t.Helper()

	var target float32 = 1023
	result := Abs(-1023)
	if target != result {
		t.Errorf("wanted %f got %f", target, result)
	}
}

func TestFloor(t *testing.T) {
	t.Helper()

	var target float32 = 120
	result := Floor(120.857)
	if target != result {
		t.Errorf("wanted %f got %f", target, result)
	}
}

func TestCeil(t *testing.T) {
	t.Helper()

	var target float32 = 121
	result := Ceil(120.11231)
	if target != result {
		t.Errorf("wanted %f got %f", target, result)
	}
}

func TestRound(t *testing.T) {
	t.Helper()

	var target float32 = 120
	result := Round(120.49)
	if target != result {
		t.Errorf("wanted %f got %f", target, result)
	}
}

func TestMod(t *testing.T) {
	t.Helper()

	var target float32 = 8
	result := Mod(18, 10)
	if target != result {
		t.Errorf("wanted %f got %f", target, result)
	}
}

