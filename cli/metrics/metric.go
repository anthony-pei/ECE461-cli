package metrics

type Metric interface {
	Calculate(m Module) float64
}

type Module interface {
	GetGitHubUrl() string
}
