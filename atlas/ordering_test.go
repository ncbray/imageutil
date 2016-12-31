package atlas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderByArea(t *testing.T) {
	items := testItems()
	expected := []*Item{
		items[5],
		items[2],
		items[4],
		items[3],
		items[0],
		items[1],
	}
	assert.Equal(t, expected, orderByArea(items))
}

func TestOrderByHeight(t *testing.T) {
	items := testItems()
	expected := []*Item{
		items[5],
		items[2],
		items[4],
		items[3],
		items[0],
		items[1],
	}
	assert.Equal(t, expected, orderByHeight(items))
}
