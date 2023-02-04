/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/anthony-pei/ECE461/cli/metrics"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

type ownerName struct {
	Owner string
	Name  string
	Url   string
}

func calcNetScore(modules []metrics.Module) {
	correctnessMetric := metrics.CorrectnessMetric{}
	licenseMetric := metrics.LicenseMetric{}
	for _, module := range modules {
		fmt.Println(module.GetName())
		correcntess := correctnessMetric.CalculateScore(module)
		license := licenseMetric.CalculateScore(module)
		fmt.Println(correcntess)
		fmt.Println(license)
	}

}

func getOwnersNamesFromFile(filename string) []ownerName {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ownerNames []ownerName

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		link := scanner.Text()
		parts := strings.Split(link, "/")
		if parts[2] != "www.npmjs.com" {
			ownerNames = append(ownerNames, ownerName{Owner: parts[3], Name: parts[4], Url: link})
		}
	}
	return ownerNames
}

// TODO: Move into github_module.go Need to make it global
func getGithubModules(ownerNames []ownerName) []metrics.Module {
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
		module := metrics.GitHubModule{Repo: repos, ContributorStats: contributorStats, Url: on.Url}
		res = append(res, module)
	}
	return res
}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		names := getOwnersNamesFromFile(args[0])
		modules := getGithubModules(names)
		calcNetScore(modules)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
