package metrics

type metric interface {
	calculate(m module) float64
}

type module interface {
	get_github_url() string
}
