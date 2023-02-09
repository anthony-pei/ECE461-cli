package metrics

import (
	"fmt"
	"log"
	"os"
)

type RampUpMetric struct {
}

func (l RampUpMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()

	dir := "temp"
	os.RemoveAll(dir)

	m.Clone(dir)
	// Analyze code
	err := os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Calculating ramp up metric for module:", m.GetGitHubUrl())
	return 0.0
}
