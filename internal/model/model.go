package model

type NodeType int

const (
	NodeText = 1 << iota
	NodePlaceholder
	NodeBlock
)

// Chunk represents a piece of text or a placeholder
type Chunk interface {
	~string
}

// Node represents a node in the parse tree
type Node[T Chunk] struct {
	Type     NodeType  // node type
	Start    int       // start token index
	End      int       // end token index
	Value    T         // we use this value to store the Chunk when Type is NodeText
	Children []Node[T] // we use this value to store children nodes when Type is NodeBlock or NodePlaceholder
}
