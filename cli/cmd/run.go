/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/anthony-pei/ECE461/cli/file_handler"
	"github.com/anthony-pei/ECE461/cli/github_util"
	"github.com/anthony-pei/ECE461/cli/metrics"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		names := file_handler.GetOwnersNamesFromFile(args[0]) // Not error checking file name
		modules := github_util.GetGithubModules(names)
		for _, module := range modules {
			netscoreModule := metrics.NetScoreMetric{URL: module.GetGitHubUrl()}
			netscoreModule.CalculateScore(module)
			fmt.Println(netscoreModule.ToNDJson())
		}
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
