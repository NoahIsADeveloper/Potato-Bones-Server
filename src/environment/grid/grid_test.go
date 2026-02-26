package grid

import (
	"potato-bones/src/utils"
	"testing"
)

func TestGridCreation(t *testing.T) {
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

	if testCell.HasCollider() {
		t.Errorf("Cell(%d, %d) has collider %v", testCellX, testCellY, testCell.GetCollider())
	}

	testCell.AddCollisionPoint(0, 255)
	if !testCell.HasCollider() {
		t.Errorf("Cell(%d, %d) missing collider", testCellX, testCellY)
	}
}

func TestGridRaycast(t *testing.T) {
	grid := NewGrid(100, 100)

	targetCell := grid.GetCell(50, 55)
	targetCell.AddCollisionPoint(0, 255)

	if !targetCell.HasCollider() {
		t.Errorf("TargetCell (%d, %d) missing collider", targetCell.x, targetCell.y)
	}

	if targetCell.x != 50 || targetCell.y != 55 {
		t.Errorf("TargetCell (50, 55) failed: got (%d, %d)", targetCell.x, targetCell.y)
	}

	success, hitCell := grid.Raycast(50.5, 50.5, utils.Pi / 2, 10)

	if !success {
		t.Errorf("Raycast failed: %v", success)
	}

	if success && (hitCell.x != targetCell.x || hitCell.y != targetCell.y) {
		t.Errorf("Wanted (%d, %d), got (%d, %d)", targetCell.x, targetCell.y, hitCell.x, hitCell.y)
	}
}
