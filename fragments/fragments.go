package fragments

type (
	NodePath []int
	Range    struct {
		BeginFragment         int
		BeginFragmentPosition int
		EndFragment           int
		EndFragmentPosition   int
	}
	node struct {
		fragmentsRange *Range
		children       []*node
	}
	Context struct {
		node      *node
		fragments *Fragments
	}
	Fragments struct {
		nodes     []*node
		fragments []interface{}
	}
)

func New(fragments []interface{}) *Fragments {
	return &Fragments{
		nodes:     []*node{},
		fragments: fragments,
	}
}
func newRange() *Range {
	return &Range{
		BeginFragment:         -1,
		BeginFragmentPosition: -1,
		EndFragment:           -1,
		EndFragmentPosition:   -1,
	}
}
func (fs *Fragments) InContext(fn func(*Context)) {
	n := &node{
		fragmentsRange: newRange(),
	}
	fn(
		&Context{
			fragments: fs,
			node:      n,
		})
	fs.nodes = append(fs.nodes, n)
}
func (fs *Fragments) Fragments() []interface{} {
	return fs.fragments
}
func (fs *Fragments) Range(path NodePath) *Range {
	// no range
	if len(fs.fragments) == 0 {
		return nil
	}
	// full range
	if len(path) == 0 {
		beginning := fs.nodes[0].fragmentsRange
		edgeChild := fs.nodes[len(fs.nodes)-1]
		for len(edgeChild.children) > 0 {
			edgeChild = edgeChild.children[len(edgeChild.children)-1]
		}
		ending := edgeChild.fragmentsRange
		return &Range{
			BeginFragment:         beginning.BeginFragment,
			BeginFragmentPosition: beginning.BeginFragmentPosition,
			EndFragment:           ending.EndFragment,
			EndFragmentPosition:   ending.EndFragmentPosition,
		}
	}
	// nested node range
	root := fs.nodes[path[0]]
	for _, idx := range path {
		root = root.children[idx]
	}
	beginning := root.fragmentsRange
	for len(root.children) > 0 {
		root = root.children[len(root.children)-1]
	}
	ending := root.fragmentsRange
	return &Range{
		BeginFragment:         beginning.BeginFragment,
		BeginFragmentPosition: beginning.BeginFragmentPosition,
		EndFragment:           ending.BeginFragment,
		EndFragmentPosition:   ending.BeginFragmentPosition,
	}
}
func (c *Context) InContext(fn func(*Context)) {
	n := &node{
		fragmentsRange: newRange(),
	}
	fn(
		&Context{
			fragments: c.fragments,
			node:      n,
		})
	c.node.children = append(c.node.children, n)
}
func (c *Context) Append(rawNewFragment interface{}) {
	switch newFragment := rawNewFragment.(type) {
	case string:
		lastFragmentIdx := len(c.fragments.fragments) - 1
		// adding new string fragment
		if lastFragmentIdx < 0 {
			c.node.fragmentsRange.BeginFragment = 0
			c.node.fragmentsRange.BeginFragmentPosition = 0
			c.node.fragmentsRange.EndFragment = 0
			c.node.fragmentsRange.EndFragmentPosition = len(newFragment) - 1
			c.fragments.fragments = append(c.fragments.fragments, rawNewFragment)
			return
		}
		switch lastFragment := c.fragments.fragments[lastFragmentIdx].(type) {
		// joining string fragments
		case string:
			updFragment := lastFragment + newFragment
			c.fragments.fragments[lastFragmentIdx] = updFragment
			if c.node.fragmentsRange.BeginFragment < 0 {
				c.node.fragmentsRange.BeginFragment = lastFragmentIdx
				c.node.fragmentsRange.BeginFragmentPosition = len(lastFragment) - 1
			}
			c.node.fragmentsRange.EndFragment = lastFragmentIdx
			c.node.fragmentsRange.EndFragmentPosition = len(updFragment) - 1
			return
		// adding new string fragment
		default:
			c.fragments.fragments = append(c.fragments.fragments, rawNewFragment)
			lastFragmentIdx = len(c.fragments.fragments) - 1
			if c.node.fragmentsRange.BeginFragment < 0 {
				c.node.fragmentsRange.BeginFragment = lastFragmentIdx
				c.node.fragmentsRange.BeginFragmentPosition = 0
			}
			c.node.fragmentsRange.EndFragment = lastFragmentIdx
			c.node.fragmentsRange.EndFragmentPosition = len(newFragment) - 1
			return
		}
	// adding new injection fragment
	default:
		c.fragments.fragments = append(c.fragments.fragments, rawNewFragment)
		lastFragmentIdx := len(c.fragments.fragments) - 1
		if c.node.fragmentsRange.BeginFragment < 0 {
			c.node.fragmentsRange.BeginFragment = lastFragmentIdx
		}
		c.node.fragmentsRange.EndFragment = lastFragmentIdx
		c.node.fragmentsRange.EndFragmentPosition = -1
		return
	}
}
