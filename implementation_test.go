package lab2

import (
	"fmt"
	"testing"
)

func TestCalculatePrefix(t *testing.T) {
	tests := []struct {
		expression string
		expected   int
		expectErr  bool
	}{
		{
			expression: "* + 9 3 - 4 2", // (9 + 3) * (4 - 2) = 24
			expected:   24,
			expectErr:  false,
		},
		{
			expression: "+ * - 15 5 2 / 20 4", // (15 - 5) * 2 + 20 / 4 = 10 * 2 + 5 = 25
			expected:   25,
			expectErr:  false,
		},
		{
			expression: "* 3 4", // 3 * 4 = 12
			expected:   12,
			expectErr:  false,
		},
		{
			expression: "- / 20 5 + 3 7", // 20 / 5 - (3 + 7) = 4 - 10 = -6
			expected:   -6,
			expectErr:  false,
		},
		{
			expression: "^ 2 3", // 2^3 = 8
			expected:   8,
			expectErr:  false,
		},
		{
			expression: "+ * 3 4 - 10 2", // 3 * 4 + (10 - 2) = 20
			expected:   20,
			expectErr:  false,
		},
		{
			expression: "+ 3", // Некоректний вираз
			expected:   0,
			expectErr:  true,
		},
		{
			expression: "/ 10 0", // Поділ на нуль
			expected:   0,
			expectErr:  true,
		},
		{
			expression: "abc", // Некоректний вираз
			expected:   0,
			expectErr:  true,
		},
		{
			expression: " ", // Пустий вираз
			expected:   0,
			expectErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.expression, func(t *testing.T) {
			result, err := CalculatePrefix(test.expression)
			if test.expectErr && err == nil {
				t.Errorf("expected error for expression %s, got nil", test.expression)
			} else if !test.expectErr && err != nil {
				t.Errorf("unexpected error for expression %s: %v", test.expression, err)
			} else if result != test.expected {
				t.Errorf("for expression %s: expected %d, got %d", test.expression, test.expected, result)
			}
		})
	}
}

func ExampleCalculatePrefix() {
	result, err := CalculatePrefix("+ 3 5")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
