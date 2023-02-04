package github_util

import (
	"context"
	"log"
	"os"

	"github.com/anthony-pei/ECE461/cli/metrics"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type OwnerName struct {
	Owner string
	Name  string
	Url   string
}

// TODO: Move into github_module.go Need to make it global
func GetGithubModules(ownerNames []OwnerName) []metrics.Module {
	res := []metrics.Module{}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	for _, on := range ownerNames {
		repos, _, err := client.Repositories.Get(ctx, on.Owner, on.Name)
		if err != nil {
			log.Panic(err)
		}
		contributorStats, _, err := client.Repositories.ListContributorsStats(ctx, on.Owner, on.Name)
		if err != nil {
			log.Panic(err)
		}
		module := GitHubModule{Repo: repos, ContributorStats: contributorStats, Url: on.Url}
		res = append(res, module)
	}
	return res
}
