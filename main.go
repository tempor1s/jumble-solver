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
		lowerString := strings.ToLower(scanner.Text())
		runes := []rune(lowerString)
		// Sort the runes
		sort.Sort(sortRunes(runes))
		s.Dict[string(runes)] = append(s.Dict[string(runes)], lowerString)
	}
}

func (s *Solver) SolveWord(inputWord string) []string {
	runes := []rune(strings.ToLower(inputWord))
	sort.Sort(sortRunes(runes))

	if s.Dict[string(runes)] != nil {
		return s.Dict[string(runes)]
	}

	return nil
}

func main() {
	solver := NewSolver()

	fmt.Println(solver.SolveWord("dgo"))
}
