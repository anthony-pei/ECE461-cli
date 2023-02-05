package metrics

import "fmt"

// Bus Factor: 1 - (1 / (# of contributors))
type BusFactorMetric struct {
	numContributors float64
}

func (l BusFactorMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()\
	fmt.Println("Calculating busFactor metric for module:", m.GetGitHubUrl())
	return 0.0
}
