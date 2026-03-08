package grid

import (
	"testing"
)

func TestCellRaycast(t *testing.T) {
	t.Helper()

	cell := NewCell(0, 0)
	cell.AddCollisionPoint(0, 0)
	cell.AddCollisionPoint(0, 255)

	success, result := cell.Raycast(127, 127, toRad(90), 10)

	if !success {
		t.Errorf("Cell raycast failed: %v", success)
	}

	if success && (result.normal != toRad(270)) {
		t.Errorf("Cell raycast incorrect normal: %v", result.normal)
	}

	if success && (result.hitX != 127 || result.hitY != 255) {
		t.Errorf("Cell raycast incorrect position: %d, %d", result.hitX, result.hitY)
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

func TestGridRaycast(t *testing.T) {
	t.Helper()

	grid := NewGrid(100, 100)

	targetCell := grid.GetCell(50, 55)
	targetCell.AddCollisionPoint(0, 255)

	if !targetCell.HasCollider() {
		t.Errorf("TargetCell (%d, %d) missing collider", targetCell.x, targetCell.y)
	}

	if targetCell.x != 50 || targetCell.y != 55 {
		t.Errorf("TargetCell (50, 55) failed: got (%d, %d)", targetCell.x, targetCell.y)
	}

	success, hitCell := grid.Raycast(50.5, 50.5, toRad(90), -1)

	if !success {
		t.Errorf("Raycast failed: %v", success)
	}

	if success && (hitCell.x != targetCell.x || hitCell.y != targetCell.y) {
		t.Errorf("Wanted (%d, %d), got (%d, %d)", targetCell.x, targetCell.y, hitCell.x, hitCell.y)
	}
}

func TestRaycastMiss(t *testing.T) {
	t.Helper()

	grid := NewGrid(50, 50)

	grid.GetCell(11, 11).AddCollisionPoint(0, 255)
	grid.GetCell(11, 9).AddCollisionPoint(0, 255)

	grid.GetCell(13, 11).AddCollisionPoint(0, 255)
	grid.GetCell(13, 9).AddCollisionPoint(0, 255)

	success, _ := grid.Raycast(10.5, 10.5, 0, -1)

	if success {
		t.Errorf("Raycast should miss but reported hit")
	}
}

func TestRaycastDiagonalHit(t *testing.T) {
	t.Helper()

	grid := NewGrid(100, 100)

	target := grid.GetCell(70, 70)
	target.AddCollisionPoint(0, 255)

	success, hit := grid.Raycast(50, 50, toRad(45), -1)

	if !success {
		t.Errorf("Diagonal raycast failed to hit")
	}

	if hit.x != 70 || hit.y != 70 {
		t.Errorf("Expected (70,70), got (%d,%d)", hit.x, hit.y)
	}
}

func TestRaycastMaxDistance(t *testing.T) {
	t.Helper()

	grid := NewGrid(100, 100)

	target := grid.GetCell(90, 50)
	target.AddCollisionPoint(0, 255)

	success, _ := grid.Raycast(50.5, 50.5, 0, 10)

	if success {
		t.Errorf("Raycast hit something past max distance")
	}
}

func TestRaycastStartInsideCollider(t *testing.T) {
	t.Helper()

	grid := NewGrid(50, 50)

	cell := grid.GetCell(20, 20)
	cell.AddCollisionPoint(0, 255)

	success, hit := grid.Raycast(20.5, 20.5, 0, 20)

	if !success {
		t.Errorf("Raycast starting inside collider should hit")
	}

	if hit.x != 20 || hit.y != 20 {
		t.Errorf("Expected (20,20), got (%d,%d)", hit.x, hit.y)
	}
}

func TestRaycastMultipleObstacles(t *testing.T) {
	t.Helper()

	grid := NewGrid(100, 100)

	near := grid.GetCell(60, 50)
	far := grid.GetCell(80, 50)

	near.AddCollisionPoint(0, 255)
	far.AddCollisionPoint(0, 255)

	success, hit := grid.Raycast(50.5, 50.5, 0, -1)

	if !success {
		t.Errorf("Raycast failed with multiple obstacles")
	}

	if success && (hit.x != 60 || hit.y != 50) {
		t.Errorf("Expected first hit (60,50), got (%d,%d)", hit.x, hit.y)
	}
}

func TestRaycastOutOfBounds(t *testing.T) {
	t.Helper()

	grid := NewGrid(50, 50)

	success, _ := grid.Raycast(25, 25, pi, -1)

	if success {
		t.Errorf("Raycast should not hit outside grid")
	}
}