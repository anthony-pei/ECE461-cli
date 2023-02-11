package metrics

import (
	"log"
	"math"
)

// Responsiveness: 1 - .01 * (open time (days) of last 10 closed issues)
type ResponsiveMaintainerMetric struct {
}

func (l ResponsiveMaintainerMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()
	issues := m.GetLast10ClosedIssues()

	score := 1.0
	for _, issue := range issues {
		score -= 0.01 * (issue.ClosedAt.Sub(issue.CreatedAt)).Hours() / 24
	}
	score = math.Max(0, score)
	log.Println("Calculating license metric for module:", m.GetGitHubUrl())

	return score
}
