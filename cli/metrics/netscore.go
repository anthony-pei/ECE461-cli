package metrics

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// If more metric scores are needed, add score into this struct, and call appropriate methods in CalculateScore()
type NetScoreMetric struct {
	URL            string  `json:"URL"`
	Correctness    float64 `json:"CORRECTNESS_SCORE"`
	License        float64 `json:"LICENSE_SCORE"`
	Busfactor      float64 `json:"BUS_FACTOR_SCORE"`
	Rampup         float64 `json:"RAMP_UP_SCORE"`
	Responsiveness float64 `json:"RESPONSIVE_MAINTAINER_SCORE"`
	Netscore       float64 `json:"NET_SCORE"`
}

var correctnessMetric Metric = CorrectnessMetric{}
var licenseMetric Metric = LicenseMetric{}
var busfactorMetric Metric = BusFactorMetric{}
var rampUpMetric Metric = RampUpMetric{}
var responsivnessMetric Metric = ResponsiveMaintainerMetric{}

func (l *NetScoreMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()

	//fmt.Println(m.GetName())
	l.Correctness = correctnessMetric.CalculateScore(m)
	l.License = licenseMetric.CalculateScore(m)
	l.Busfactor = busfactorMetric.CalculateScore(m)
	l.Rampup = rampUpMetric.CalculateScore(m)
	l.Responsiveness = responsivnessMetric.CalculateScore(m)

	l.Netscore = 0.4*l.Responsiveness + 0.1*l.Rampup + 0.2*l.Busfactor + 0.2*l.License + 0.1*l.Correctness
	log.Info("Calculating Netscore metric for module:", m.GetGitHubUrl())
	return l.Netscore
}

func (l NetScoreMetric) ToNDJson() string {
	b, err := json.Marshal(l)
	if err != nil {
		log.Debug("Error with NDJson conversion")
	}
	return string(b)
}
