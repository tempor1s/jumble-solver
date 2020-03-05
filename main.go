package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Len() int {
	return len(s)
}
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortWord(word string) string {
	lowerWord := strings.ToLower(word)
	runes := []rune(lowerWord)

	sort.Sort(sortRunes(runes))

	return string(runes)
}

// Solver represents the word jumble solver
type Solver struct {
	Dict map[string][]string
}

func NewSolver() *Solver {
	solver := &Solver{Dict: make(map[string][]string)}

	solver.GenerateDict()

	return solver
}

func (s *Solver) GenerateDict() {
	path := "/usr/share/dict/words"

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		// Convert string to lower and convert it into rune format
		normalWord := strings.ToLower(scanner.Text())
		sortedWord := sortWord(normalWord)

		s.Dict[sortedWord] = append(s.Dict[sortedWord], normalWord)
	}
}

func (s *Solver) SolveWord(inputWord string) []string {
	sortedWord := sortWord(inputWord)

	if s.Dict[sortedWord] != nil {
		return s.Dict[sortedWord]
	}

	return nil
}

func (s *Solver) SolveWords(inputWords []string) []string {
	var output []string

	for _, word := range inputWords {
		solvedWords := s.SolveWord(word)

		output = append(output, solvedWords...)
	}

	return output
}

func main() {
	solver := NewSolver()

	fmt.Println(solver.SolveWord("crtolualca"))
}
