package metrics

import log "github.com/sirupsen/logrus"

type LicenseMetric struct {
}

func (l LicenseMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()\
	log.Info("Calculating license metric for module:", m.GetGitHubUrl())
	return 0.0
}
