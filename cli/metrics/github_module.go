package metrics

type github_module struct {
	url string
}

func (g github_module) get_url() string {
	return g.url
}
