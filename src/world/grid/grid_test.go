package grid

import (
	"potato-bones/src/utils/math"
	"testing"
)

func TestCellRaycast(t *testing.T) {
	t.Helper()

	cell := NewCell(0, 0)
	cell.AddCollisionPoint(0, 0)
	cell.AddCollisionPoint(0, 255)

	success, result := cell.Raycast(127, 127, math.ToRad(180), 255)

	if !success {
		t.Errorf("Cell raycast failed: %v", success)
	}

	if success && (math.Round(result.normal) != math.Round(math.ToRad(0))) {
		t.Errorf("Cell raycast incorrect normal: %v, %f", result.normal, math.ToRad(0))
	}

	if success && (math.Round(result.hit.X) != 0 || math.Round(result.hit.Y) != 127) {
		t.Errorf("Cell raycast incorrect position: %f, %f", result.hit.X, result.hit.Y)
	}

	cell = NewCell(0, 0)
	cell.AddCollisionPoint(0, 0)
	cell.AddCollisionPoint(0, 255)

	success, _ = cell.Raycast(127, 127, math.ToRad(0), 255)

	if success {
		t.Errorf("Cell raycast failed: %v", success)
	}

	cell = NewCell(0, 0)
	cell.AddCollisionPoint(0, 127)
	cell.AddCollisionPoint(255, 127)

	success, result = cell.Raycast(127, 0, math.ToRad(90), 255)

	if !success {
		t.Errorf("Cell raycast failed: %v", success)
	}

	if success && (math.Round(result.normal) != math.Round(math.ToRad(270))) {
		t.Errorf("Cell raycast incorrect normal: %v, %f", result.normal, math.ToRad(270))
	}

	if success && (math.Round(result.hit.X) != 127 || math.Round(result.hit.Y) != 127) {
		t.Errorf("Cell raycast incorrect position: %f, %f", result.hit.X, result.hit.Y)
	}

	cell = NewCell(0, 0)
	cell.AddCollisionPoint(0, 127)
	cell.AddCollisionPoint(255, 127)

	success, result = cell.Raycast(127, 0, math.ToRad(90), 3)

	if success {
		t.Errorf("Cell raycast failed: %v", success)
	}
}

func TestCellRaycastMultiwall(t *testing.T) {
	t.Helper()

	cell := NewCell(0, 0)
	cell.AddCollisionPoint(255, 0)
	cell.AddCollisionPoint(255, 255)
	cell.AddCollisionPoint(200, 255)
	cell.AddCollisionPoint(200, 0)

	success, result := cell.Raycast(127, 127, math.ToRad(0), 255)

	if !success {
		t.Errorf("Cell raycast failed: %v", success)
	}

	if success && (math.Round(result.normal) != math.Round(math.ToRad(180))) {
		t.Errorf("Cell raycast incorrect normal: %v, %f", result.normal, math.ToRad(180))
	}

	if success && (math.Round(result.hit.X) != 200 || math.Round(result.hit.Y) != 127) {
		t.Errorf("Cell raycast incorrect position: %f, %f", result.hit.X, result.hit.Y)
	}
}

func TestGridCreation(t *testing.T) {
	t.Helper()

	grid := NewGrid(100, 150)
	sizeX := len(grid.cells)
	sizeY := len(grid.cells[1])

	if sizeX != 100 || sizeY != 150 {
		t.Errorf("NewGrid(x, y) failed: got (%d, %d), want (100, 150)", sizeX, sizeY)
	}

	cell := grid.GetCell(5, 10)
	cellX, cellY := cell.GetPosition()

	if cellX != 5 || cellY != 10 {
		t.Errorf("GetCell(x, y) failed: got (%d, %d), want (5, 10)", cellX, cellY)
	}

	if cell.HasCollider() {
		t.Errorf("Cell(%d, %d) has collider %v", cellX, cellY, cell.GetCollider())
	}

	cell.AddCollisionPoint(0, 255)
	if !cell.HasCollider() {
		t.Errorf("Cell(%d, %d) missing collider", cellX, cellY)
	}

	testCell := grid.GetCell(5, 10)
	testCellX, testCellY := testCell.GetPosition()
	if testCellX != 5 || testCellY != 10 {
		t.Errorf("GetCell(x, y) failed: got (%d, %d), want (5, 10)", testCell, testCell)
	}

	if !testCell.HasCollider() {
		t.Errorf("Cell(%d, %d) has collider %v", testCellX, testCellY, testCell.GetCollider())
	}
}

func TestRaycastDiagonalHit(t *testing.T) {
	t.Helper()

	grid := NewGrid(100, 100)

	target := grid.GetCell(70, 70)
	target.AddCollisionPoint(0, 255)
	target.AddCollisionPoint(255, 0)

	success, hit := grid.Raycast(50, 50, math.ToRad(45), -1)

	if !success {
		t.Errorf("Diagonal raycast failed to hit")
	}

	if hit.position.X != 70.5 || hit.position.Y != 70.5 {
		t.Errorf("Expected (70.5, 70.5), got (%f, %f)", hit.position.X, hit.position.Y)
	}

	if math.Round(hit.normal * 10) != math.Round(math.ToRad(225) * 10) {
		t.Errorf("Expected (%f), got (%f)", math.ToRad(225), hit.normal)
	}
}

func TestRaycastMaxDistance(t *testing.T) {
	t.Helper()

	grid := NewGrid(100, 100)

	target := grid.GetCell(90, 50)
	target.AddCollisionPoint(0, 255)
	target.AddCollisionPoint(255, 0)

	success, _ := grid.Raycast(50.5, 50.5, 0, 10)

	if success {
		t.Errorf("Raycast hit something past max distance")
	}
}

func TestRaycastOutOfBounds(t *testing.T) {
	t.Helper()

	grid := NewGrid(50, 50)

	success, _ := grid.Raycast(25, 25, math.Pi, -1)

	if success {
		t.Errorf("Raycast should not hit outside grid")
	}
}