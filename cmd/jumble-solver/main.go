package main

import (
	"fmt"

	solver "github.com/tempor1s/jumble-solver"
)

func main() {
	solver := solver.NewSolver()

	fmt.Println(solver.SolveWord("crtolualca"))
}
