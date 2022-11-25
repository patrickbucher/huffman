package huffman

import "testing"

func TestCoding(t *testing.T) {
	messages := []string{
		"abacabacabaa",
		"This is a test.",
		"Это тоже тест.",
		"The quick brown fox jumps over the lazy dog.",
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
