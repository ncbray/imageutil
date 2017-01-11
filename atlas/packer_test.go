package atlas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

	boxW, boxH := PackSimple(SortItems(items, AREA_ORDER), 0)

	assert.Equal(t, 64, boxW)
	assert.Equal(t, 64, boxH)

	assert.Equal(t, 50, items[0].X)
	assert.Equal(t, 20, items[0].Y)

	assert.Equal(t, 0, items[1].X)
	assert.Equal(t, 30, items[1].Y)

	assert.Equal(t, 30, items[2].X)
	assert.Equal(t, 0, items[2].Y)

	assert.Equal(t, 30, items[3].X)
	assert.Equal(t, 20, items[3].Y)

	assert.Equal(t, 0, items[4].X)
	assert.Equal(t, 20, items[4].Y)

	assert.Equal(t, 0, items[5].X)
	assert.Equal(t, 0, items[5].Y)
}
