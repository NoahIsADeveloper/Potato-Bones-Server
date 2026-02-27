package grid

import (
	"testing"
)

func TestDegRad(t *testing.T) {
	t.Helper()

	var targetDeg uint16 = 180
	rad := toRad(targetDeg)
	resultDeg := toDeg(rad)

	if resultDeg != targetDeg {
		t.Errorf("wanted %d got %d", targetDeg, resultDeg)
	}

	if rad != pi {
		t.Errorf("wanted %f got %f", pi, halfPi)
	}
}

func TestAbs(t *testing.T) {
	t.Helper()

	var target float32 = 1023
	result := abs(-1023)
	if target != result {
		t.Errorf("wanted %f got %f", target, result)
	}
}

func TestFloor(t *testing.T) {
	t.Helper()

	var target float32 = 120
	result := floor(120.857)
	if target != result {
		t.Errorf("wanted %f got %f", target, result)
	}
}

func TestCeil(t *testing.T) {
	t.Helper()

	var target float32 = 121
	result := ceil(120.11231)
	if target != result {
		t.Errorf("wanted %f got %f", target, result)
	}
}

func TestRound(t *testing.T) {
	t.Helper()

	var target float32 = 120
	result := round(120.49)
	if target != result {
		t.Errorf("wanted %f got %f", target, result)
	}
}

func TestMod(t *testing.T) {
	t.Helper()

	var target float32 = 8
	result := mod(18, 10)
	if target != result {
		t.Errorf("wanted %f got %f", target, result)
	}
}

