package metrics

import "fmt"

type CorrectnessMetric struct {
}

func (l CorrectnessMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()\
	fmt.Println("Calculating correctness metric for module:", m.GetGitHubUrl())
	correctness := 1.0 - (0.00001 + float64(m.GetOpenIssuesCount())/float64(m.GetStargazerCount()))
	return correctness
}
