package metrics

import "fmt"

type RampUpMetric struct {
}

func (l RampUpMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()\
	fmt.Println("Calculating license metric for module:", m.GetGitHubUrl())
	return 0.0
}
