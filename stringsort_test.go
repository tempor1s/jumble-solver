package solver

import (
	"fmt"
	"testing"
)

var (
	sortCases = []struct {
		UnsortedString string
		SortedString   string
	}{
		{"cba", "abc"},
		{"zzz", "zzz"},
		{"zyx", "xyz"},
		{"zqa", "aqz"},
		{"CBA", "abc"},
	}
)

func TestStringSorting(t *testing.T) {
	t.Run("test string sort correctly", func(t *testing.T) {
		for _, test := range sortCases {
			t.Run(fmt.Sprintf("%q when sorted becomes %q\n", test.UnsortedString, test.SortedString), func(t *testing.T) {
				got := sortWord(test.UnsortedString)

				if got != test.SortedString {
					t.Errorf("got sorted string of %q but wanted sorted string of %q\n", got, test.SortedString)
				}
			})
		}
	})
}
