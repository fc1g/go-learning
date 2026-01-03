package pkg

import (
	"strings"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello world  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("words are not the same length")
			return
		}

		for i, actualLen := 0, len(actual); i < actualLen; i++ {
			word := actual[i]
			expectedWord := c.expected[i]

			if strings.Compare(word, expectedWord) != 0 {
				t.Errorf("words are not the same")
				return
			}
		}
	}
}
