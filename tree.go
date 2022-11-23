package huffman

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
