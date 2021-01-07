package integers

import (
	"reflect"
	"testing"
)

func TestMakeRangeWithStep(t *testing.T) {
	tests := map[string]struct {
		min  int
		max  int
		step int
		expected []int
	}{
		"case-1": {
			min: 1,
			max: 7,
			step: 3,
			expected: []int{1,4,7},
		},

		"min > max": {
			min: 1,
			max: -7,
			step: 3,
			expected: []int{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := MakeRangeWithStep(tt.min, tt.max, tt.step); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("MakeRangeWithStep() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMakeRange(t *testing.T) {
	tests := map[string]struct {
		min  int
		max  int
		expected []int
	}{
		"case-1": {
			min: 1,
			max: 7,
			expected: []int{1,2,3,4,5,6,7},
		},

		"min > max": {
			min: 1,
			max: -7,
			expected: []int{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := MakeRange(tt.min, tt.max); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("MakeRange() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := map[string]struct {
		haystack []int
		needle int
		expected int
	}{
		"found-1": {
			haystack: []int{1,2,3,4,5,6,7},
			needle: 2,
			expected: 1,
		},
		"found-2": {
			haystack: []int{1,2,3,4,5,6,7},
			needle: 1,
			expected: 0,
		},
		"empty haystack": {
			haystack: []int{},
			needle: 10,
			expected: -1,
		},
		"not found": {
			haystack: []int{1,2,3,4,5,6,7},
			needle: 10,
			expected: -1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := Find(tt.haystack, tt.needle); got != tt.expected {
				t.Errorf("Find() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := map[string]struct {
		haystack []int
		needle int
		expected bool
	}{
		"found-1": {
			haystack: []int{1,2,3,4,5,6,7},
			needle: 2,
			expected: true,
		},
		"found-2": {
			haystack: []int{1,2,3,4,5,6,7},
			needle: 1,
			expected: true,
		},
		"not found": {
			haystack: []int{1,2,3,4,5,6,7},
			needle: 10,
			expected: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := Contains(tt.haystack, tt.needle); got != tt.expected {
				t.Errorf("Contains() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestRemoveIndex(t *testing.T) {
	tests := map[string]struct {
		array []int
		indexToRemove int
		expected []int
	}{
		"case 1": {
			array: []int{1,2,3,4,5,6,7},
			indexToRemove: 2,
			expected: []int{1,2,4,5,6,7},
		},
		"case 2": {
			array: []int{1,2,3,4,5,6,7},
			indexToRemove: 0,
			expected: []int{2,3,4,5,6,7},
		},
		"case 3": {
			array: []int{1,2,3,4,5,6,7},
			indexToRemove: 6,
			expected: []int{1,2,3,4,5,6},
		},
		"case 4": {
			array: []int{4,9},
			indexToRemove: 1,
			expected: []int{4},
		},
		"case 5": {
			array: []int{4},
			indexToRemove: 0,
			expected: []int{},
		},
		"case 6": {
			array: []int{},
			indexToRemove: 1,
			expected: []int{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := RemoveIndex(tt.array, tt.indexToRemove); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("RemoveIndex() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestRemoveValue(t *testing.T) {
	tests := map[string]struct {
		array []int
		value int
		expected []int
	}{
		"case 1": {
			array: []int{1,2,3,4,5,6,7},
			value: 7,
			expected: []int{1,2,3,4,5,6},
		},
		"case 2": {
			array: []int{1,2,3,4,5,6,7},
			value: 1,
			expected: []int{2,3,4,5,6,7},
		},
		"case 3": {
			array: []int{1,2,3,4,5,6,7},
			value: 6,
			expected: []int{1,2,3,4,5,7},
		},
		"case 4": {
			array: []int{4,9},
			value: 9,
			expected: []int{4},
		},
		"case 5": {
			array: []int{4},
			value: 4,
			expected: []int{},
		},
		"case 6": {
			array: []int{},
			value: 1,
			expected: []int{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := RemoveValue(tt.array, tt.value); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("RemoveValue() = %v, want %v", got, tt.expected)
			}
		})
	}
}