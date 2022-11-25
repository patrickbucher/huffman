package main

import (
	"fmt"

	huffman "github.com/patrickbucher/huffmann"
)

func main() {
	message := "abacabacabbabac"

	compressed, tree := huffman.Encode(message)
	original := huffman.Decode(compressed, tree)

	fmt.Println(message, compressed, original)
	fmt.Printf("encoded message '%s' (%d characters) in %d bits\n",
		message, len(message), len(compressed))
}
