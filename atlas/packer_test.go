package atlas

import (
	"testing"
)

func checkInt(actual int, expected int, t *testing.T) {
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}

func testItems() []*Item {
	return []*Item{
		{Width: 10, Height: 10},
		{Width: 10, Height: 10},
		{Width: 20, Height: 20},
		{Width: 20, Height: 10},
		{Width: 30, Height: 10},
		{Width: 30, Height: 20},
	}
}

func TestPackSimple(t *testing.T) {
	items := testItems()

	boxW, boxH := PackSimple(items, 0)

	checkInt(boxW, 64, t)
	checkInt(boxH, 64, t)

	checkInt(items[0].X, 50, t)
	checkInt(items[0].Y, 20, t)

	checkInt(items[1].X, 0, t)
	checkInt(items[1].Y, 30, t)

	checkInt(items[2].X, 30, t)
	checkInt(items[2].Y, 0, t)

	checkInt(items[3].X, 30, t)
	checkInt(items[3].Y, 20, t)

	checkInt(items[4].X, 0, t)
	checkInt(items[4].Y, 20, t)

	checkInt(items[5].X, 0, t)
	checkInt(items[5].Y, 0, t)
}
