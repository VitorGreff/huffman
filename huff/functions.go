package huff

import (
	"huffman/queue"
	"log"
	"sort"
)

func BuildQueueFromFrequencyMap(m map[string]int) queue.Queue[Tree] {
	var q queue.Queue[Tree]
	for k, v := range m {
		q.Enqueue(Tree{Root: LeafNode{Element: k, Weight: v}})
	}
	return q
}

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

func DecodePrefixCode(t Tree, prefix string) string {
	var decodedString string
	node := t.Root

	for _, char := range prefix {

		switch string(char) {
		case "1":
			node = node.(InternalNode).Right
		case "0":
			node = node.(InternalNode).Left
		}
		if node.IsLeaf() {
			decodedString += node.(LeafNode).Element
			node = t.Root
		}
	}
	return decodedString
}

func EncodeText(t Tree, text string) string {
	var code string
	m := BuildPrefixTable(t, "", map[string]string{})

	for _, char := range text {
		code += m[string(char)]
	}
	return code
}

func PrintTree(t Tree) {
	if t.Root == nil {
		return
	}
	switch node := t.Root.(type) {
	case LeafNode:
		log.Printf("Leaf node -> Character: %s Weight: %d\n", t.Root.(LeafNode).Element, t.Root.GetWeight())
		return
	case InternalNode:
		log.Printf("Internal node -> Weight: %d\n", t.Root.GetWeight())
		PrintTree(Tree{node.Left})
		PrintTree(Tree{node.Right})

	}
}

func BuildPrefixTable(t Tree, prefix string, m map[string]string) map[string]string {
	switch t.Root.(type) {
	case LeafNode:
		m[t.Root.(LeafNode).Element] = prefix
	case InternalNode:
		BuildPrefixTable(Tree{Root: t.Root.(InternalNode).Left}, prefix+"0", m)
		BuildPrefixTable(Tree{Root: t.Root.(InternalNode).Right}, prefix+"1", m)
	}
	return m
}

func sortTreeQueue(q *queue.Queue[Tree]) {
	slice := q.Items
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].GetWeight() < slice[j].GetWeight()
	})
	q.Items = slice
}
