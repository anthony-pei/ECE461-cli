package metrics

import (
	"encoding/json"
	"fmt"
	"log"
)

// If more metric scores are needed, add score into this struct, and call appropriate methods in CalculateScore()
type NetScoreMetric struct {
	URL            string  `json:"url"`
	Correctness    float64 `json:"correctness"`
	License        float64 `json:"license"`
	Busfactor      float64 `json:"busfactor"`
	Rampup         float64 `json:"ramp_up"`
	Responsiveness float64 `json:"responsive_maintainer"`
	Netscore       float64 `json:"net_score"`
}

var correctnessMetric CorrectnessMetric
var licenseMetric LicenseMetric
var busfactorMetric BusFactorMetric
var rampUpMetric RampUpMetric
var responsivnessMetric ResponsiveMaintainerMetric

func (l *NetScoreMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()

	//fmt.Println(m.GetName())
	l.Correctness = correctnessMetric.CalculateScore(m)
	l.License = licenseMetric.CalculateScore(m)
	l.Busfactor = busfactorMetric.CalculateScore(m)
	l.Rampup = rampUpMetric.CalculateScore(m)
	l.Responsiveness = responsivnessMetric.CalculateScore(m)

	l.Netscore = 0.4*l.Responsiveness + 0.1*l.Rampup + 0.2*l.Busfactor + 0.2*l.License + 0.1*l.Correctness
	log.Println("Calculating Netscore metric for module:", m.GetGitHubUrl())
	return l.Netscore
}

func (l NetScoreMetric) ToNDJson() string {
	b, err := json.Marshal(l)
	fmt.Println(l.Netscore)
	if err != nil {
		log.Fatal("Error with NDJson conversion")
	}
	return string(b)
}
