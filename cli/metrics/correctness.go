package metrics

import (
	log "github.com/sirupsen/logrus"
	"math"
)

type CorrectnessMetric struct {
}

func (l CorrectnessMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()
	log.Info("Calculating correctness metric for module:", m.GetGitHubUrl())
	correctness := 1.0 - (0.00001 + float64(m.GetOpenIssuesCount())) / (0.00001 + float64(m.GetStargazerCount()))
	correctness = math.Min(correctness, 1.0)
	correctness = math.Max(correctness, 0.0)
	return correctness
}
