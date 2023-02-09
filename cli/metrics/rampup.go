package metrics

import (
	"fmt"
)

type RampUpMetric struct {
}

func (l RampUpMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()

	dir := "temp"
	m.Clone(dir)
	// Analyze code
	// Clean dir
	fmt.Println("Calculating ramp up metric for module:", m.GetGitHubUrl())
	return 0.0
}
