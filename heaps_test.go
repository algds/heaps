package heaps

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	input    []interface{}
	expected [][]interface{}
}{
	{
		[]interface{}{1, 2, 3, 4},
		[][]interface{}{
			{1, 2, 3, 4},
			{2, 1, 3, 4},
			{3, 1, 2, 4},
			{1, 3, 2, 4},
			{2, 3, 1, 4},
			{3, 2, 1, 4},
			{4, 2, 1, 3},
			{2, 4, 1, 3},
			{1, 4, 2, 3},
			{4, 1, 2, 3},
			{2, 1, 4, 3},
			{1, 2, 4, 3},
			{1, 3, 4, 2},
			{3, 1, 4, 2},
			{4, 1, 3, 2},
			{1, 4, 3, 2},
			{3, 4, 1, 2},
			{4, 3, 1, 2},
			{4, 3, 2, 1},
			{3, 4, 2, 1},
			{2, 4, 3, 1},
			{4, 2, 3, 1},
			{3, 2, 4, 1},
			{2, 3, 4, 1},
		},
	},
	{
		[]interface{}{4, 3, 2, 1},
		[][]interface{}{
			{4, 3, 2, 1},
			{3, 4, 2, 1},
			{2, 4, 3, 1},
			{4, 2, 3, 1},
			{3, 2, 4, 1},
			{2, 3, 4, 1},
			{1, 3, 4, 2},
			{3, 1, 4, 2},
			{4, 1, 3, 2},
			{1, 4, 3, 2},
			{3, 4, 1, 2},
			{4, 3, 1, 2},
			{4, 2, 1, 3},
			{2, 4, 1, 3},
			{1, 4, 2, 3},
			{4, 1, 2, 3},
			{2, 1, 4, 3},
			{1, 2, 4, 3},
			{1, 2, 3, 4},
			{2, 1, 3, 4},
			{3, 1, 2, 4},
			{1, 3, 2, 4},
			{2, 3, 1, 4},
			{3, 2, 1, 4},
		},
	},
}

func TestPermutations(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res := Permutations(tc.input); !reflect.DeepEqual(res, tc.expected) {
			t.Errorf("Expected (%v), got (%v)", tc.expected, res)
		}
	}
}

func BenchmarkPermutations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Permutations(tc.input)
		}
	}
}
