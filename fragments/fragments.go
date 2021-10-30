package fragments

type (
	NodePath []int
	Node     struct {
		beginFragment         int
		beginFragmentPosition int
		endFragment           int
		endFragmentPosition   int
		children              []*Node
	}
	Context struct {
		fragments *Fragments
		node      *Node
	}
	Fragments struct {
		nodesMapping []NodePath
		fragments    []interface{}
	}
)

func New() *Fragments {
	return &Fragments{
		nodesMapping: []NodePath{},
		fragments:    []interface{}{},
	}
}
func (fs *Fragments) Context() *Context {
	return &Context{
		fragments: fs,
		node:      &Node{},
	}
}
func (fs *Fragments) Fragments() []interface{} {
	return fs.fragments
}

func (c *Context) Context() *Context {
	return &Context{
		fragments: c.fragments,
		node:      &Node{},
	}
}
func (c *Context) Append(rawNewFragment interface{}) {
	newFragment, ok := rawNewFragment.(string)
	if !ok {
		c.fragments.fragments = append(c.fragments.fragments, rawNewFragment)
		return
	}
	lastFragmentIdx := len(c.fragments.fragments) - 1
	if len(c.fragments.fragments) == 0 {
		c.fragments.fragments = append(c.fragments.fragments, rawNewFragment)
		return
	}
	rawLastFragment := c.fragments.fragments[lastFragmentIdx]
	lastFragment, ok := rawLastFragment.(string)
	if !ok {
		c.fragments.fragments = append(c.fragments.fragments, rawNewFragment)
		return
	}
	c.fragments.fragments[lastFragmentIdx] = lastFragment + newFragment
}
