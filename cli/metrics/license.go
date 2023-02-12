package metrics

import "log"

type LicenseMetric struct {
	licenseKey string
}

func (l LicenseMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()\
	log.Println("Calculating license metric for module:", m.GetGitHubUrl())
	score := 0.0
	key := m.GetLicense()
	if len(key) == 0 {
		return score
	}
	approvedLicenses := [8]string{"mit", "cc0-1.0", "bsd-2-clause", "bsd-3-clause", "lgpl-2.1", "lgpl-2.1-or-later", "unlicense", "bsl-1.0"}
	for _, license := range approvedLicenses {
		if license == key {
			score = 1
			break
		}
	}

	return score
}
