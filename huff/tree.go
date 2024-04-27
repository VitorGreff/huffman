package huff

import (
	"huffman/queue"
	"sort"
)

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

func BuildTree(m map[string]int) Tree {
	q := BuildQueueFromFrequencyMap(m)
	var tmp1, tmp2, tmp3 Tree

	for len(q.Items) > 1 {
		sortTreeQueue(&q)
		tmp1 = q.Dequeue()
		tmp2 = q.Dequeue()
		tmp3 = Tree{Root: InternalNode{Left: tmp1.Root, Right: tmp2.Root, Weight: tmp1.GetWeight() + tmp2.GetWeight()}}
		q.Enqueue(tmp3)
	}
	return tmp3
}

func sortTreeQueue(q *queue.Queue[Tree]) {
	slice := q.Items
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].GetWeight() < slice[j].GetWeight()
	})
	q.Items = slice
}
