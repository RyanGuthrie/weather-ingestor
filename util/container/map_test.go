package container

import (
	"reflect"
	"testing"
)

func TestValues(t *testing.T) {
	type testCase struct {
		name     string
		input    map[string]int
		expected []int
	}

	testCases := []testCase{
		{
			name:     "emptyMap",
			input:    map[string]int{},
			expected: []int{},
		},
		{
			name:     "singleEntry",
			input:    map[string]int{"one": 1},
			expected: []int{1},
		},
		{
			name:     "multipleEntry",
			input:    map[string]int{"one": 1, "two": 2, "three": 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Values(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}