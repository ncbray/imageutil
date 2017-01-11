package atlas

import (
	"sort"
)

type SortOrder int

const (
	WIDTH_ORDER SortOrder = iota
	HEIGHT_ORDER
	LARGEST_ORDER
	PERIMETER_ORDER
	AREA_ORDER
)

type byMetric struct {
	Items []*Item
}

func (a byMetric) Len() int {
	return len(a.Items)
}

func (a byMetric) Swap(i, j int) {
	a.Items[i], a.Items[j] = a.Items[j], a.Items[i]
}

func (a byMetric) Less(i, j int) bool {
	itemI := a.Items[i]
	itemJ := a.Items[j]
	if itemI.metric != itemJ.metric {
		return itemI.metric > itemJ.metric
	} else {
		// Stablize sort on equal metrics
		return itemI.id < itemJ.id
	}
}

func sortMetric(item *Item, order SortOrder) int {
	switch order {
	case WIDTH_ORDER:
		return item.Width
	case HEIGHT_ORDER:
		return item.Height
	case LARGEST_ORDER:
		if item.Width > item.Height {
			return item.Width
		} else {
			return item.Height
		}
	case PERIMETER_ORDER:
		return item.Width + item.Height
	case AREA_ORDER:
		return item.Width * item.Height
	default:
		panic(order)
	}
}

func SortItems(items []*Item, order SortOrder) []*Item {
	for i, item := range items {
		item.id = i
		item.metric = sortMetric(item, order)
	}
	out := make([]*Item, len(items))
	copy(out, items)
	sort.Sort(byMetric{Items: out})
	return out
}
