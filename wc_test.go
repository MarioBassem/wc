package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBytes(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected uint32
	}{
		"empty": {
			input:    "",
			expected: 0,
		},
		"notEmpty": {
			input:    "abcde",
			expected: 5,
		},
		"multiline": {
			input:    "hamada\nyel3ab",
			expected: 13,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			want := tt.expected
			got, err := countBytes(r)
			assert.NoError(t, err)

			assert.Equal(t, want, got)
		})
	}
}

func TestCountLines(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected uint32
	}{
		"empty": {
			input:    "",
			expected: 0,
		},
		"oneLine": {
			input:    "hamada\n",
			expected: 1,
		},
		"multiline": {
			input:    "h\n\nmada",
			expected: 2,
		},
		"endsWithNewLine": {
			input:    "word1\n",
			expected: 1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			want := tt.expected
			got, err := countLines(r)
			assert.NoError(t, err)

			assert.Equal(t, want, got)
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected uint32
	}{
		"empty": {
			input:    "",
			expected: 0,
		},
		"multiword": {
			input:    "multiword sentences are awesome",
			expected: 4,
		},
		"endsWithSpace": {
			input:    "this		sentence ends with a space ",
			expected: 6,
		},
		"hasSpecialCharacters": {
			input:    "- * • word",
			expected: 1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			want := tt.expected
			got, err := countWords(r)
			assert.NoError(t, err)

			assert.Equal(t, want, got)
		})
	}
}

func TestCountCharacters(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected uint32
	}{
		"empty": {
			input:    "",
			expected: 0,
		},
		"notEmpty": {
			input:    "僤凘墈",
			expected: 3,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			want := tt.expected
			got, err := countChars(r)
			assert.NoError(t, err)

			assert.Equal(t, want, got)
		})
	}
}
