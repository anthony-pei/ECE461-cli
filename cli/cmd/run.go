/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/anthony-pei/ECE461/cli/metrics"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Args:  cobra.ExactArgs(1),
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		git_url := args[0]
		fmt.Println(git_url)
		g := metrics.GitHubModule{
			Url: git_url,
		}
		m := metrics.LicenseMetric{}
		m.Calculate(g)
		// TODO: Read line by line file in go
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
