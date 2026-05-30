package counter

import (
	"bytes"
	"strings"
	"testing"
)

func TestCountWords(t *testing.T) {
	testsCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "simple five words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "empty",
			input: "",
			wants: 0,
		},
		{
			name:  "single space",
			input: " ",
			wants: 0,
		},
		{
			name:  "two space",
			input: "  ",
			wants: 0,
		},
	}
	for _, value := range testsCases {
		t.Run(value.name, func(t *testing.T) {
			result := CountWords(bytes.NewReader([]byte(value.input)))
			if result != value.wants {
				t.Log("expect result of", value.wants, "got result of:", result)
				t.Fail()
			}
		})
	}
}

func TestGetCounts(t *testing.T) {
	testsCases := []struct {
		name  string
		input string
		wants Counts
	}{
		{
			name:  "simple five words",
			input: "one two three four five",
			wants: Counts{Lines: 1, Words: 5, Bytes: 24},
		},
		{
			name:  "empty",
			input: "",
			wants: Counts{Lines: 0, Words: 0, Bytes: 0},
		},
	}
	for _, value := range testsCases {
		t.Run(value.name, func(t *testing.T) {
			result := GetCounts(strings.NewReader(value.input))
			if value.wants != result {
				t.Log("expect result of", value.wants, "got result of:", result)
				t.Fail()
			}

		})
	}
}
