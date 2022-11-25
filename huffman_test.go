package huffman

import "testing"

func TestCoding(t *testing.T) {
	messages := []string{
		"abacabacabaa",
		"This is a test.",
		"Это тоже тесть.",
		"The quick brown fox jumps over the lazy dog.",
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
		tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
		veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea
		commodo consequat. Duis aute irure dolor in reprehenderit in voluptate
		velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
		occaecat cupidatat non proident, sunt in culpa qui officia deserunt
		mollit anim id est laborum.`,
	}
	for _, message := range messages {
		encoded, tree := Encode(message)
		decoded := Decode(encoded, tree)
		if decoded != message {
			t.Errorf("encoded message '%s' with tree %v as %v, decoded as '%s'",
				message, tree, encoded, decoded)
		}
	}
}
