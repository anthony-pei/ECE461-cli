package github_util

import (
	"context"
	log "github.com/sirupsen/logrus"
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

// TODO: Handle recieving errors from github API, no need to panic move on to next OwnerName (log issue)
func GetGithubModules(ownerNames []OwnerName) []metrics.Module {
	res := []metrics.Module{}
	ctx := context.Background()
	token, has_token := os.LookupEnv("GITHUB_TOKEN")
	if !has_token {
		log.Debug("GITHUB_TOKEN variable not in environment, please set it in enviroment variables")
	}
	if len(token) == 0 {
		log.Debug("GITHUB_TOKEN variable is present, but not set to a value")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	for _, on := range ownerNames {
		repos, _, err := client.Repositories.Get(ctx, on.Owner, on.Name)
		if err != nil {
			log.Debug(err)
		}

		opt := &github.ListContributorsOptions{
			ListOptions: github.ListOptions{PerPage: 30},
		}
		var allContributors []*github.Contributor
		for {
			contributors, resp, err := client.Repositories.ListContributors(ctx, on.Owner, on.Name, opt)
			if err != nil {
				log.Debug(err)
			}
			allContributors = append(allContributors, contributors...)
			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}

		// Can create error with contributor stats and umarshalling, not consistent
		module := GitHubModule{Repo: repos, Contributors: allContributors, Url: on.Url}
		res = append(res, module)

	}
	return res
}
