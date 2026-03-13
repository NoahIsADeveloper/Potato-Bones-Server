package grid

import "testing"

func TestCellRaycast(t *testing.T) {
	t.Helper()

	cell := NewCell(0, 0)
	cell.AddCollisionPoint(0, 0)
	cell.AddCollisionPoint(0, 255)

	success, result := cell.Raycast(127, 127, toRad(180), 255)

	if !success {
		t.Errorf("Cell raycast failed: %v", success)
	}

	if success && (round(result.normal) != round(toRad(0))) {
		t.Errorf("Cell raycast incorrect normal: %v, %f", result.normal, toRad(0))
	}

	if success && (round(result.hit.x) != 0 || round(result.hit.y) != 127) {
		t.Errorf("Cell raycast incorrect position: %f, %f", result.hit.x, result.hit.y)
	}

	cell = NewCell(0, 0)
	cell.AddCollisionPoint(0, 0)
	cell.AddCollisionPoint(0, 255)

	success, _ = cell.Raycast(127, 127, toRad(0), 255)

	if success {
		t.Errorf("Cell raycast failed: %v", success)
	}

	cell = NewCell(0, 0)
	cell.AddCollisionPoint(0, 127)
	cell.AddCollisionPoint(255, 127)

	success, result = cell.Raycast(127, 0, toRad(90), 255)

	if !success {
		t.Errorf("Cell raycast failed: %v", success)
	}

	if success && (round(result.normal) != round(toRad(270))) {
		t.Errorf("Cell raycast incorrect normal: %v, %f", result.normal, toRad(270))
	}

	if success && (round(result.hit.x) != 127 || round(result.hit.y) != 127) {
		t.Errorf("Cell raycast incorrect position: %f, %f", result.hit.x, result.hit.y)
	}

	cell = NewCell(0, 0)
	cell.AddCollisionPoint(0, 127)
	cell.AddCollisionPoint(255, 127)

	success, result = cell.Raycast(127, 0, toRad(90), 3)

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

	success, result := cell.Raycast(127, 127, toRad(0), 255)

	if !success {
		t.Errorf("Cell raycast failed: %v", success)
	}

	if success && (round(result.normal) != round(toRad(180))) {
		t.Errorf("Cell raycast incorrect normal: %v, %f", result.normal, toRad(180))
	}

	if success && (round(result.hit.x) != 200 || round(result.hit.y) != 127) {
		t.Errorf("Cell raycast incorrect position: %f, %f", result.hit.x, result.hit.y)
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

	success, hit := grid.Raycast(50, 50, toRad(45), -1)

	if !success {
		t.Errorf("Diagonal raycast failed to hit")
	}

	if hit.position.x != 70.5 || hit.position.y != 70.5 {
		t.Errorf("Expected (70.5, 70.5), got (%f, %f)", hit.position.x, hit.position.y)
	}

	if round(hit.normal * 10) != round(toRad(225) * 10) {
		t.Errorf("Expected (%f), got (%f)", toRad(225), hit.normal)
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

	success, _ := grid.Raycast(25, 25, pi, -1)

	if success {
		t.Errorf("Raycast should not hit outside grid")
	}
}