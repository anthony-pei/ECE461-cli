package metrics

import (
	"math"

	log "github.com/sirupsen/logrus"
)

// Responsiveness: 1 - .01 * (open time (days) of last 10 closed issues)
type ResponsiveMaintainerMetric struct {
}

func (l ResponsiveMaintainerMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()
	issues := m.GetLast10ClosedIssues()

	score := 1.0
	log.Debugf("Received %v closed issues", len(issues))
	for i, issue := range issues {
		log.Debugf("Issue %v : CreatedAt: %v : Closed At %v", i, issue.CreatedAt, issue.ClosedAt)
		score -= 0.01 * (issue.ClosedAt.Sub(issue.CreatedAt)).Hours() / 24
	}
	score = math.Max(0, score)
	log.Info("Calculating license metric for module:", m.GetGitHubUrl())
	return score
}
