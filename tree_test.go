package huffman

import (
	"testing"
)

func TestCountFrequency(t *testing.T) {
	tests := map[string]map[rune]int{
		"":    {},
		"a":   {'a': 1},
		"abc": {'a': 1, 'b': 1, 'c': 1},
		"aaa": {'a': 3},
		"Dilbert": {
			'D': 1,
			'i': 1,
			'l': 1,
			'b': 1,
			'e': 1,
			'r': 1,
			't': 1,
		},
		"every evening": {
			'e': 4,
			'v': 2,
			'r': 1,
			'y': 1,
			'n': 2,
			'i': 1,
			'g': 1,
			' ': 1,
		},
	}
	for text, expected := range tests {
		actual := CountFrequency(text)
		if !MapEqual(actual, expected) {
			t.Errorf(`CountFrequency("%s"): expected %v, got %v`,
				text, expected, actual)
		}
	}
}

func MapEqual[K, V comparable](a, b map[K]V) bool {
	if len(a) != len(b) {
		return false
	}
	for ka, va := range a {
		vb, ok := b[ka]
		if !ok {
			return false
		}
		if va != vb {
			return false
		}
	}
	return true
}
