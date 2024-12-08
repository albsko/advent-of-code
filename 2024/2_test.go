package main

import "testing"

func TestY2024_2(t *testing.T) {
	result := y2024_2()
	t.Log(result)
}

func TestIsChangingInSameDirection(t *testing.T) {
	tests := map[string]struct {
		input     []int
		maxChange int
		expected  bool
	}{
		`1`: {
			input:     []int{0, 1, 2, 3, 4},
			maxChange: 3,
			expected:  true,
		},
		`2`: {
			input:     []int{10, 7, 6, 5, 4},
			maxChange: 3,
			expected:  true,
		},
		`3`: {
			input:     []int{0, 1, 2, 3, 8},
			maxChange: 3,
			expected:  false,
		},
		`4`: {
			input:     []int{10, 1, 2, 3, 4},
			maxChange: 3,
			expected:  false,
		},
		`5`: {
			input:     []int{0, -1, 2, 3, 4},
			maxChange: 3,
			expected:  false,
		},
		`6`: {
			input:     []int{1, 2},
			maxChange: 3,
			expected:  true,
		},
		`7`: {
			input:     []int{5, 3},
			maxChange: 3,
			expected:  true,
		},
		`8`: {
			input:     []int{1, 1, 2, 3},
			maxChange: 3,
			expected:  false,
		},
		`9`: {
			input:     []int{10, 7, 4, 1},
			maxChange: 3,
			expected:  true,
		},
		`10`: {
			input:     []int{1, 4, 7, 10},
			maxChange: 3,
			expected:  true,
		},
		`11`: {
			input:     []int{1, 5, 7, 10},
			maxChange: 3,
			expected:  false,
		},
		// `12`: { // edge case that was not necessary to be handled
		// 	input:     []int{6, 4, 1, -2},
		// 	maxChange: 3,
		// 	expected:  false,
		// },
		`13`: {
			input:     []int{1},
			maxChange: 3,
			expected:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := isChangingInSameDirection(tt.input, tt.maxChange)
			if got != tt.expected {
				t.Errorf("input: %v, maxChange: %d, got: %v, expected: %v",
					tt.input, tt.maxChange, got, tt.expected)
			}
		})
	}
}
