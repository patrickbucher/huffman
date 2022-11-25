package huffman

import (
	"fmt"
	"sort"
	"strings"
)

type Node struct {
	Symbols []rune
	Weight  int
	Left    *Node
	Right   *Node
}

func (n Node) IsLeaf() bool {
	return len(n.Symbols) == 1
}

func (n Node) String() string {
	symbols := ""
	for _, symbol := range n.Symbols {
		symbols += fmt.Sprintf("%c ", symbol)
	}
	symbols = strings.TrimSpace(symbols)
	return fmt.Sprintf("%v (%d) [left: %v, right: %v]",
		symbols, n.Weight, n.Left, n.Right)
}

type Nodes []Node

func (n Nodes) Len() int           { return len(n) }
func (n Nodes) Less(i, j int) bool { return n[i].Weight < n[j].Weight }
func (n Nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func BuildHuffmanTree(message string) *Node {
	freqs := countFrequency(message)
	leaves := createLeaves(freqs)
	return buildTree(leaves)
}

func buildTree(nodes Nodes) *Node {
	n := len(nodes)
	if n == 0 {
		return nil
	} else if n == 1 {
		return &nodes[0]
	} else {
		sort.Sort(nodes)
		smallest := nodes[0]
		nextInSize := nodes[1]
		newNode := combineNodes(smallest, nextInSize)
		remainder := append(nodes[2:], newNode)
		return buildTree(remainder)
	}
}

func combineNodes(left, right Node) Node {
	return Node{
		Symbols: append(left.Symbols, right.Symbols...),
		Weight:  left.Weight + right.Weight,
		Left:    &left,
		Right:   &right,
	}
}

func createLeaves(frequencies map[rune]int) []Node {
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

func countFrequency(text string) map[rune]int {
	count := make(map[rune]int)
	for _, c := range text {
		count[c]++
	}
	return count
}
