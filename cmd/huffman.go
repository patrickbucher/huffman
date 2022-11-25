package main

import (
	"fmt"
	"os"

	huffman "github.com/patrickbucher/huffmann"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [message]\n", os.Args[0])
		os.Exit(1)
	}
	message := os.Args[1]

	compressed, tree := huffman.Encode(message)
	original := huffman.Decode(compressed, tree)

	fmt.Println(message, "->", compressed, "->", original)
	fmt.Printf("encoded message '%s' (%d characters) in %d bits\n",
		message, len(message), len(compressed))
}
