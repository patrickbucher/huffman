package huffman

import "fmt"

// Bits represents a sequence of bits.
type Bits []byte

// Encode applies Huffman coding to the given message, and returns the encoded
// bits, as well as the tree built up for this process.
func Encode(message string) (Bits, *Node) {
	tree := buildHuffmanTree(message)
	if tree == nil || (tree.Left == nil && tree.Right == nil) {
		panic("nothing to encode")
	}
	bits := make(Bits, 0)
	for _, c := range message {
		bits = append(bits, encode(c, tree)...)
	}
	return bits, tree
}

func encode(symbol rune, tree *Node) Bits {
	bits := make(Bits, 0)
	if tree.IsLeaf() {
		if tree.Symbols[0] == symbol {
			return bits
		}
		panic(fmt.Sprintf("%c not found in %v", symbol, tree))
	} else {
		if contains(tree.Left.Symbols, symbol) {
			bits = append(bits, 0)
			return append(bits, encode(symbol, tree.Left)...)
		} else if contains(tree.Right.Symbols, symbol) {
			bits = append(bits, 1)
			return append(bits, encode(symbol, tree.Right)...)
		} else {
			panic(fmt.Sprintf("%c not found in %v", symbol, tree))
		}
	}
}

func contains(symbols []rune, symbol rune) bool {
	for _, c := range symbols {
		if c == symbol {
			return true
		}
	}
	return false
}

// Decode decodes the given bit sequence using the given tree, and returns the
// decoded string.
func Decode(bits Bits, tree *Node) string {
	message := ""
	var symbol rune
	for len(bits) > 0 {
		symbol, bits = decode(bits, tree)
		message = fmt.Sprintf("%s%c", message, symbol)
	}
	return message
}

func decode(bits Bits, tree *Node) (rune, Bits) {
	if tree.IsLeaf() {
		return tree.Symbols[0], bits
	} else if len(bits) > 0 {
		bit := bits[0]
		bits := bits[1:]
		if bit == 0 {
			return decode(bits, tree.Left)
		} else if bit == 1 {
			return decode(bits, tree.Right)
		} else {
			panic(fmt.Sprintf("%d is not a valid bit", bit))
		}
	} else {
		panic("bits consumed prematurely")
	}
}
