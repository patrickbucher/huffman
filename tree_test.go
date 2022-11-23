package huffman

import (
	"fmt"
	"reflect"
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
		if !mapEqual(actual, expected) {
			t.Errorf(`CountFrequency("%s"): expected %v, got %v`,
				text, expected, actual)
		}
	}
}

func TestCreateLeaves(t *testing.T) {
	type testCase struct {
		frequencies map[rune]int
		nodes       []Node
	}
	tests := []testCase{
		testCase{
			frequencies: map[rune]int{'a': 1, 'b': 2, 'c': 3},
			nodes: []Node{
				Node{[]rune{'a'}, 1, nil, nil},
				Node{[]rune{'b'}, 2, nil, nil},
				Node{[]rune{'c'}, 3, nil, nil},
			},
		},
		testCase{
			frequencies: map[rune]int{'x': 7, 'y': 5, 'z': 9},
			nodes: []Node{
				Node{[]rune{'y'}, 5, nil, nil},
				Node{[]rune{'x'}, 7, nil, nil},
				Node{[]rune{'z'}, 9, nil, nil},
			},
		},
	}
	for _, test := range tests {
		expected := test.nodes
		actual := CreateLeaves(test.frequencies)
		fmt.Println(actual, expected)
		if !nodeSliceEqual(actual, expected) {
			t.Errorf(`CreateLeaves(%v): expected %v, got %v`,
				test.frequencies, expected, actual)
		}
	}
}

func TestBuildTree(t *testing.T) {
	type testCase struct {
		nodes []Node
		tree  Node
	}
	tests := []testCase{
		testCase{
			nodes: []Node{},
			tree:  Node{},
		},
		testCase{
			nodes: []Node{
				Node{Symbols: []rune{'a'}, Weight: 1},
				Node{Symbols: []rune{'b'}, Weight: 1},
				Node{Symbols: []rune{'c'}, Weight: 3},
				Node{Symbols: []rune{'d'}, Weight: 4},
				Node{Symbols: []rune{'e'}, Weight: 7},
			},
			tree: Node{
				Symbols: []rune{'a', 'b', 'c', 'd', 'e'},
				Weight:  16,
				// TODO: figure out the tree structure on paper...
				Left:  nil,
				Right: nil,
			},
		},
	}
	for _, test := range tests {
		// TODO: call the function (to be implemented)
		fmt.Println(test)
	}
}

func nodeSliceEqual(a, b []Node) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !reflect.DeepEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func mapEqual[K, V comparable](a, b map[K]V) bool {
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
