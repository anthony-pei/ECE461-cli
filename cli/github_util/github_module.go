package github_util

import (
	"log"
	"os/exec"

	"github.com/google/go-github/github"
)

type GitHubModule struct {
	Url          string
	Repo         *github.Repository
	Contributors []*github.Contributor
}

func (g GitHubModule) GetGitHubUrl() string {
	return g.Url
}

func (g GitHubModule) GetLicense() string {
	license := g.Repo.GetLicense()
	return *license.Name
}

func (g GitHubModule) GetOpenIssuesCount() int {
	return *g.Repo.OpenIssuesCount
}

func (g GitHubModule) GetStargazerCount() int {
	return *g.Repo.StargazersCount
}

func (g GitHubModule) GetContributorCount() int {
	return len(g.Contributors)
}

func (g GitHubModule) GetName() string {
	return g.Repo.GetFullName()
}

func (g GitHubModule) Clone(dir string) {
	cmd := exec.Command("git", "clone", *g.Repo.CloneURL, dir)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err) // Maybe no need to be Fatal?
	}
}
