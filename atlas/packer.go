package atlas

type Item struct {
	// Input
	Width  int
	Height int

	// Output
	X int
	Y int

	id int
}

func (i *Item) area() int {
	return i.Width * i.Height
}

type simpleSizePolicy struct {
	width  int
	height int
}

func (p *simpleSizePolicy) Next() (int, int) {
	if p.width == 0 {
		p.width = 1
		p.height = 1
	}
	if p.width <= p.height {
		p.width *= 2
	} else {
		p.height *= 2
	}
	return p.width, p.height
}

func (p *simpleSizePolicy) MinimumArea(area int) (int, int) {
	if p.width == 0 {
		p.width = 1
		p.height = 1
	}
	for p.width*p.height < area {
		p.Next()
	}
	return p.width, p.height
}

func tryPackSimple(items []*Item, padding int, boxW int, boxH int) bool {
	x := padding
	y := padding
	height := 0
	for _, item := range items {
		w := item.Width + padding
		h := item.Height + padding
		if x+w > boxW {
			x = padding
			y += height
			height = 0
		}
		item.X = x
		item.Y = y

		x += w
		if h > height {
			height = h
		}

		if x > boxW || y+height > boxH {
			return false
		}
	}
	return true
}

func PackSimple(items []*Item, padding int) (int, int) {
	// TODO area zero items?

	policy := &simpleSizePolicy{}
	// TODO padding?
	area := 0
	for i := 0; i < len(items); i++ {
		area += items[i].area()
	}
	boxW, boxH := policy.MinimumArea(area)

	items = orderByHeight(items)
	for {
		if tryPackSimple(items, padding, boxW, boxH) {
			return boxW, boxH
		}
		boxW, boxH = policy.Next()
	}
}

type treeNode struct {
	x      int
	y      int
	width  int
	height int
	left   *treeNode
	right  *treeNode
	filled bool
}

func (n *treeNode) add(w int, h int) (int, int, bool) {
	if n.left != nil {
		// The node has already been split, so recurse.
		x, y, ok := n.left.add(w, h)
		if ok {
			return x, y, ok
		}
		return n.right.add(w, h)
	} else if n.filled {
		// The node is completely filled.
		return 0, 0, false
	} else if w > n.width || h > n.height {
		// The node is too small.
		return 0, 0, false
	} else if w == n.width && h == n.height {
		// The node is an exact fit.
		n.filled = true
		return n.x, n.y, true
	} else {
		// The node is too big, so split it.
		// TODO if the width or the height is an exact match, just clip a bit off
		// instead of doing two splits.
		dw := n.width - w
		dh := n.height - h
		if dw > dh {
			n.left = &treeNode{x: n.x, y: n.y, width: w, height: n.height}
			n.right = &treeNode{x: n.x + w, y: n.y, width: dw, height: n.height}
		} else {
			n.left = &treeNode{x: n.x, y: n.y, width: n.width, height: h}
			n.right = &treeNode{x: n.x, y: n.y + h, width: n.width, height: dh}
		}
		// The left node should always fit this item.
		return n.left.add(w, h)
	}
}

func tryPackTree(items []*Item, padding int, boxW int, boxH int) bool {
	// Clip the upper and left margins for padding.
	root := &treeNode{
		x:      padding,
		y:      padding,
		width:  boxW - padding,
		height: boxH - padding,
	}
	for _, item := range items {
		w := item.Width + padding
		h := item.Height + padding
		x, y, ok := root.add(w, h)
		if !ok {
			return false
		}
		item.X = x
		item.Y = y
	}
	return true
}

func PackTree(items []*Item, padding int) (int, int) {
	// TODO area zero items?

	policy := &simpleSizePolicy{}
	// TODO padding?
	area := 0
	for i := 0; i < len(items); i++ {
		area += items[i].area()
	}
	boxW, boxH := policy.MinimumArea(area)

	items = orderByArea(items)
	for {
		if tryPackTree(items, padding, boxW, boxH) {
			return boxW, boxH
		}
		boxW, boxH = policy.Next()
	}
}
