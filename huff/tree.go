package huff

type Node interface {
	IsLeaf() bool
	GetWeight() int
}

type LeafNode struct {
	Element string
	Weight  int
}

func (node LeafNode) IsLeaf() bool   { return true }
func (node LeafNode) GetWeight() int { return node.Weight }

type InternalNode struct {
	Weight int
	Left   Node
	Right  Node
}

func (node InternalNode) IsLeaf() bool   { return false }
func (node InternalNode) GetWeight() int { return node.Weight }

type Tree struct {
	Root Node
}

func (node Tree) GetWeight() int { return node.Root.GetWeight() }
