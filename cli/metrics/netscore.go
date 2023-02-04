package metrics

import "fmt"

// If more metric scores are needed, add score into this struct, and call appropriate methods in CalculateScore()
type NetScoreMetric struct {
	correctness    float64
	license        float64
	busfactor      float64
	rampup         float64
	responsiveness float64
	netscore       float64
}

var correctnessMetric CorrectnessMetric
var licenseMetric LicenseMetric
var busfactorMetric BusFactorMetric
var rampUpMetric RampUpMetric
var responsivnessMetric ResponsiveMaintainerMetric

func (l NetScoreMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()

	fmt.Println(m.GetName())
	l.correctness = correctnessMetric.CalculateScore(m)
	l.license = licenseMetric.CalculateScore(m)
	l.busfactor = busfactorMetric.CalculateScore(m)
	l.rampup = rampUpMetric.CalculateScore(m)
	l.responsiveness = responsivnessMetric.CalculateScore(m)

	l.netscore = 0.4*l.responsiveness + 0.1*l.rampup + 0.2*l.busfactor + 0.2*l.license + 0.1*l.correctness
	fmt.Println("Calculating Netscore metric for module:", m.GetGitHubUrl())
	return l.netscore
}
