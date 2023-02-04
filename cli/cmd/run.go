/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/anthony-pei/ECE461/cli/github_util"
	"github.com/anthony-pei/ECE461/cli/metrics"
	"github.com/spf13/cobra"
)

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

func getOwnersNamesFromFile(filename string) []github_util.OwnerName {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ownerNames []github_util.OwnerName

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		link := scanner.Text()
		parts := strings.Split(link, "/")
		if parts[2] != "www.npmjs.com" {
			ownerNames = append(ownerNames, github_util.OwnerName{Owner: parts[3], Name: parts[4], Url: link})
		}
	}
	return ownerNames
}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		names := getOwnersNamesFromFile(args[0]) // Not error checking file name
		modules := github_util.GetGithubModules(names)
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
