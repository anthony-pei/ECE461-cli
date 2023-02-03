package metrics

import "fmt"

// Responsiveness: 1 - .01 * (open time (days) of last 10 closed issues)
type ResponsiveMaintainerMetric struct {
}

func (l ResponsiveMaintainerMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()\
	fmt.Println("Calculating license metric for module:", m.GetGitHubUrl())
	return 0.0
}
