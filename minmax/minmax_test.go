package main

import(
	"testing"
	"reflect"
)

func TestMinMax(t *testing.T) {

	tests := []struct {
		name     string
		min      float64
		max      float64
		values   []float64
		expected []float64
	}{
		{
			name:     "Multiple values within range",
			min:      2,
			max:      6,
			values:   []float64{1, 2, 3, 4, 7},
			expected: []float64{2, 3, 4},
		},
		{
			name:     "Single value within range",
			min:      0,
			max:      10,
			values:   []float64{5},
			expected: []float64{5},
		},
		{
			name:     "Min greater than max",
			min:      10,
			max:      5,
			values:   []float64{1, 2, 3},
			expected: []float64{},
		},
		{
			name:     "No values within range",
			min:      10,
			max:      20,
			values:   []float64{1, 2, 3},
			expected: []float64{},
		},
		{
			name:     "Both min and max negative",
			min:      -10,
			max:      -2,
			values:   []float64{-15, -8, -3},
			expected: []float64{-8, -3},
		},
		{
			name:     "Min and max equal",
			min:      5,
			max:      5,
			values:   []float64{5, 6, 4},
			expected: []float64{5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minmax(tt.min, tt.max, tt.values...)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf(
					"%s failed: expected %v, got %v",
					tt.name,
					tt.expected,
					result,
				)
			}
		})
	}
}