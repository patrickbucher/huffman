package huffman

import (
	"sort"
)

// Node is a tree node, i.e. either a leaf with a single symbol, or a node
// further above the structure, containing all the symbols of the nodes below,
// referenced as left and right nodes. The weight is either the sum of all the
// leafs below, or the leaf's weight.
type Node struct {
	Symbols []rune
	Weight  int
	Left    *Node
	Right   *Node
}

type Nodes []Node

func (n Nodes) Len() int           { return len(n) }
func (n Nodes) Less(i, j int) bool { return n[i].Weight < n[j].Weight }
func (n Nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

// IsLeaf returns true iff the node represents exactly one symbol.
func (n Node) IsLeaf() bool {
	return len(n.Symbols) == 1
}

// CountFrequency creates a map from the given text with the number of
// occurrences for every rune.
func CountFrequency(text string) map[rune]int {
	count := make(map[rune]int)
	for _, c := range text {
		count[c]++
	}
	return count
}

// CreateLeaves converts the rune frequencies to a slice of nodes, each
// containing a single symbol, and the frequency as the weight. The nodes are
// sorted in ascending order by their weight.
func CreateLeaves(frequencies map[rune]int) []Node {
	leaves := make(Nodes, 0)
	for symbol, count := range frequencies {
		node := Node{
			Symbols: []rune{symbol},
			Weight:  count,
			Left:    nil,
			Right:   nil,
		}
		leaves = append(leaves, node)
	}
	sort.Sort(leaves)
	return leaves
}
