package atlas

import (
	"sort"
)

type byHeight struct {
	Items []*Item
}

func (a byHeight) Len() int {
	return len(a.Items)
}

func (a byHeight) Swap(i, j int) {
	a.Items[i], a.Items[j] = a.Items[j], a.Items[i]
}

func (a byHeight) Less(i, j int) bool {
	if a.Items[i].Height != a.Items[j].Height {
		return a.Items[i].Height > a.Items[j].Height
	} else if a.Items[i].Width != a.Items[j].Width {
		return a.Items[i].Width > a.Items[j].Width
	} else {
		// Stable sort
		return a.Items[i].id < a.Items[j].id
	}
}

func orderByHeight(items []*Item) []*Item {
	for i, item := range items {
		item.id = i
	}
	out := make([]*Item, len(items))
	copy(out, items)
	sort.Sort(byHeight{Items: out})
	return out
}

type byArea struct {
	Items []*Item
}

func (a byArea) Len() int {
	return len(a.Items)
}

func (a byArea) Swap(i, j int) {
	a.Items[i], a.Items[j] = a.Items[j], a.Items[i]
}

func (a byArea) Less(i, j int) bool {
	areaI := a.Items[i].area()
	areaJ := a.Items[j].area()
	if areaI != areaJ {
		return areaI > areaJ
	} else if a.Items[i].Height != a.Items[j].Height {
		return a.Items[i].Height > a.Items[j].Height
	} else {
		// Stable sort
		return a.Items[i].id < a.Items[j].id
	}
}

func orderByArea(items []*Item) []*Item {
	for i, item := range items {
		item.id = i
	}
	out := make([]*Item, len(items))
	copy(out, items)
	sort.Sort(byArea{Items: out})
	return out
}
