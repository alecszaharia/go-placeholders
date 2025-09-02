package model

type NodeType int

const (
	NodeText = 1 << iota
	NodePlaceholder
	NodeBlock
	NodeRoot
)

// Node represents a node in the parse tree
type Node struct {
	Type         NodeType    // node type
	Start        int         // start token index
	End          int         // end token index
	ContentStart int         // start token index of the content (for block placeholders)
	ContentEnd   int         // end token index of the content (for block placeholders)
	Value        interface{} // we use this value to store the Chunk when Type is NodeText
	Children     []*Node     // we use this value to store children nodes when Type is NodeBlock or NodePlaceholder
	Parent       *Node       `json:"-"` // pointer to parent node, nil if root
}

func (n *Node) AddChild(s *Node) {
	s.Parent = n
	n.Children = append(n.Children, s)
}

type Attr struct {
	Name  string
	Value string
}

type Text struct {
	Text string
}

type Placeholder struct {
	Name         string
	IsBlock      bool
	Attrs        []*Attr
	ContentStart int // start token index of the content (for block placeholders)
	ContentEnd   int // end token index of the content (for block placeholders)
}

func NewTextNode(start int, end int, parent *Node) *Node {
	node := &Node{
		Type:   NodeText,
		Start:  start,
		End:    end,
		Parent: parent,
		Value:  &Text{},
	}

	return node
}
func NewPlaceholderNode(start int, end int) *Node {
	node := &Node{
		Type:  NodePlaceholder,
		Start: start,
		End:   end,
	}

	return node

}
func NewBlockNode(start int, end int) *Node {
	node := &Node{
		Type:  NodeBlock,
		Start: start,
		End:   end,
	}

	return node
}
func NewRootNode(start int, end int) *Node {
	node := &Node{
		Type:  NodeRoot,
		Start: start,
		End:   end,
	}
	return node
}

func (n *Node) HasParent() bool {
	return n.Parent != nil
}

// Context
// ---------------------------------------------------------------------------
type Context struct {
	CurrentNode   *Node
	ParentContext *Context
}

func (c *Context) IsRoot() bool {
	return c.ParentContext == nil
}

func (c *Context) GetPlaceholderByName(name string) bool {
	panic("implement me: GetPlaceholderByName")
}

func (c *Context) GetPlaceholderByNameAndAttr(name string, attr string, value string) bool {
	panic("implement me: GetPlaceholderByNameAndAttr")
}
