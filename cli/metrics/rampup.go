package metrics

import (
	"math"
	"github.com/anthony-pei/ECE461/cli/code_analysis"
	log "github.com/sirupsen/logrus"
)

type RampUpMetric struct {
}

var analyzeDirFunction = code_analysis.AnalyzeCodeBase

func (l RampUpMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()
	log.Info("Calculating ramp up metric for module:", m.GetGitHubUrl())
	dir := "temp"
	code_analysis.CleanDir(dir, false)
	m.Clone(dir)

	_, code, comments, _ := analyzeDirFunction(dir) // Returns total lines, lines of code, comments, and blank lines. Not using first and last at the moment

	log.Infof("Code Lines: %v, Comment Lines: %v", code, comments)
	code_analysis.CleanDir(dir, true)

	if code == 0 {
		return 0.0
	}
	score := float64(comments) / float64(code) / 0.2
	score = math.Min(score, 1)
	score = math.Max(score, 0)
	return score
}
