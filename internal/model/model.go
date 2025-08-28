package model

type NodeType int

const (
	NodeText = 1 << iota
	NodePlaceholder
	NodeBlock
)

// Chunk represents a piece of text or a placeholder
type Chunk interface {
	Text | Placeholder
}

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

type Context struct {
	CurrentNode   *Node
	ParentContext *Node
}

func NewTextNode(start int, end int, parent *Node) *Node {
	node := &Node{
		Type:   NodeText,
		Start:  start,
		End:    end,
		Parent: parent,
		Value:  &Text{},
	}

	if parent != nil {
		node.Parent = parent
		parent.Children = append(parent.Children, node)
	}
	return node
}
func NewPlaceholderNode(start int, end int, parent *Node) *Node {
	node := &Node{
		Type:  NodePlaceholder,
		Start: start,
		End:   end,
	}

	if parent != nil {
		node.Parent = parent
		parent.Children = append(parent.Children, node)
	}

	return node

}
func NewBlockNode(start int, end int, parent *Node) *Node {
	node := &Node{
		Type:  NodeBlock,
		Start: start,
		End:   end,
	}

	if parent != nil {
		node.Parent = parent
		parent.Children = append(parent.Children, node)
	}

	return node
}

func (n *Node) HasParent() bool {
	return n.Parent != nil
}
