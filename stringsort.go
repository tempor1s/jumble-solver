package solver

import (
	"sort"
	"strings"
)

// sortRunes is needed so that we can implement sort interface on a mock rune object
type sortRunes []rune

// Less is needed to implement sort interface
func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

// Len is needed to implement sort interface
func (s sortRunes) Len() int {
	return len(s)
}

// Swap is needed to implement sort interface
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// SortWord will sort the given word based off of alphabetical order
func sortWord(word string) string {
	lowerWord := strings.ToLower(word)
	runes := []rune(lowerWord)

	sort.Sort(sortRunes(runes))

	return string(runes)
}
