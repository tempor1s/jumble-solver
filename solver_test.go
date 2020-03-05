package solver

import (
	"fmt"
	"testing"
)

var (
	solverCases = []struct {
		UnsolvedWord string
		SolvedWord   string
	}{
		{"crtolualca", "calculator"},
		{"lsodi", "soldi"},
		{"amot", "atmo"},
		{"cesinec", "science"},
		{"mabe", "beam"},
		{"iudqli", "liquid"},
	}
)

// TODO: Test array, because this is very fragile atm.
func TestSingleSolve(t *testing.T) {
	t.Run("find the correct word from a jumble of letters", func(t *testing.T) {
		solver := NewSolver()
		for _, test := range solverCases {
			t.Run(fmt.Sprintf("%q when sovled becomes %q\n", test.UnsolvedWord, test.SolvedWord), func(t *testing.T) {
				got := solver.SolveWord(test.UnsolvedWord)

				if got[0] != test.SolvedWord {
					t.Errorf("got solved word of %q but wanted solved word of %q\n", got[0], test.SolvedWord)
				}
			})
		}
	})
}

func BenchmarkSolve(b *testing.B) {
	solver := NewSolver()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		solver.SolveWord("crtolualca")
	}
}
