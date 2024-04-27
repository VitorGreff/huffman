package huff

import "huffman/queue"

func BuildQueueFromFrequencyMap(m map[string]int) queue.Queue[Tree] {
	var q queue.Queue[Tree]
	for k, v := range m {
		q.Enqueue(Tree{Root: LeafNode{Element: k, Weight: v}})
	}
	return q
}
