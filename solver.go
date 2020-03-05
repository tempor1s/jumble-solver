package solver

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Solver represents the word jumble solver
type Solver struct {
	Dict map[string][]string
}

// NewSolver will create a new solver object and create underlying map of the dictonary
func NewSolver() *Solver {
	solver := &Solver{Dict: make(map[string][]string)}

	solver.generateDict()

	return solver
}

// generateDict will generate a sorted string key with array of strings with those characters map
func (s *Solver) generateDict() {
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

// SolveWord takes a jumbled word and then returns the solved version of it
func (s *Solver) SolveWord(inputWord string) []string {
	sortedWord := sortWord(inputWord)

	if s.Dict[sortedWord] != nil {
		return s.Dict[sortedWord]
	}

	return nil
}

// SolveWords takes a slice of words, and then will solve every one of them and return a slice of the answers
func (s *Solver) SolveWords(inputWords []string) []string {
	var output []string

	for _, word := range inputWords {
		solvedWords := s.SolveWord(word)

		output = append(output, solvedWords...)
	}

	return output
}
