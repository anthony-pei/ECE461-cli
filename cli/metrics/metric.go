package metrics

// Score must be between [0, 1]
// limit API calls.
type Metric interface {
	CalculateScore(m Module) float64
}

type Module interface {
	GetGitHubUrl() string
}
