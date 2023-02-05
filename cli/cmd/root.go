/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/anthony-pei/ECE461/cli/file_handler"
	"github.com/anthony-pei/ECE461/cli/github_util"
	"github.com/anthony-pei/ECE461/cli/metrics"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: `Run takes GitHub and npmjs URLs as input and outputs NDJSON containing the URL, NetScore, and various metrics`,
	Long: `
	
Run is a CLI program written in Go. Its core function is to evaluate GitHub repositories and npm modules.
Execute ./run <url_string> and the program will output a NDJSON containing the repository URL, a NetScore, 
and various metrics such as RampUp, Correctness, BusFactor, ResponsiveMaintainer, LicenseScore, etc. The 
metrics and weights can be configured through ./run config.`,
	Args: cobra.ExactArgs(1),
	// test with go run main.go -- https://api.github.com/repos/anthony-pei/ECE461
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Example = `  cli [url_string]`
}
