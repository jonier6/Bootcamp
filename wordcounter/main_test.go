package main

import (
	"strings"
	"testing"
)

func TestWordScenarios(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Single sentence with words",
			input:    "hello world from go\n",
			expected: 4,
		},
		{
			name:     "Multiple sentences with words per sentence",
			input:    "hello world\nthis is go\n",
			expected: 5,
		},
		{
			name:     "Single word",
			input:    "golang\n",
			expected: 1,
		},
		{
			name:     "Single composed word read-only",
			input:    "read-only\n",
			expected: 1,
		},
		{
			name:     "Multiple break lines",
			input:    "\n\n\n",
			expected: 0,
		},
		{
			name:     "Exit uppercase stops processing",
			input:    "hello world\nEXIT\nthis should not count\n",
			expected: 2,
		},
		{
			name:     "exit lowercase stops processing",
			input:    "hello world\nexit\nthis should not count\n",
			expected: 2,
		},
		{
			name:     "Exit capitalized stops processing",
			input:    "hello world\nExit\nthis should not count\n",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			reader := strings.NewReader(tt.input)
			lines := GetInput(reader)

			result := ProcesarEntrada(lines, false)

			if result != tt.expected {
				t.Errorf("Word counting failed in %s: expected %d, got %d",
					tt.name, tt.expected, result)
			}
		})
	}
}

func TestLineScenarios(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Single line",
			input:    "hello world\n",
			expected: 1,
		},
		{
			name:     "Multiple lines without break line in between",
			input:    "line one\nline two\nline three\n",
			expected: 3,
		},
		{
			name:     "Multiple lines with break lines in between",
			input:    "line one\n\nline two\n\n\n",
			expected: 5,
		},
		{
			name:     "Exit at start of input",
			input:    "EXIT\nline one\nline two\n",
			expected: 0,
		},
		{
			name:     "Exit in middle of input",
			input:    "line one\nexit\nline two\n",
			expected: 1,
		},
		{
			name:     "Exit at end of input",
			input:    "line one\nline two\nExit\n",
			expected: 2,
		},
		{
			name:     "Exit word inside line does NOT stop",
			input:    "line with exit inside\nanother line\nEXIT\n",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			reader := strings.NewReader(tt.input)
			lines := GetInput(reader)

			result := ProcesarEntrada(lines, true)

			if result != tt.expected {
				t.Errorf("Line counting failed in %s: expected %d, got %d",
					tt.name, tt.expected, result)
			}
		})
	}
}