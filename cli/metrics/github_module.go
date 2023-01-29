package metrics

type GitHubModule struct {
	Url string
}

func (g GitHubModule) GetGitHubUrl() string {
	return g.Url
}
