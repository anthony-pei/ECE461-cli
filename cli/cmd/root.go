/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"
	"encoding/json"
	"os"
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
	"net/http"
	"github.com/spf13/cobra"
	"github.com/joho/godotenv"
)
type NPMResponse struct {
	Repository struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"repository"`
}

type GitHubResponse struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"owner"`
	HTMLURL          string      `json:"html_url"`
	Description      interface{} `json:"description"`
	Fork             bool        `json:"fork"`
	URL              string      `json:"url"`
	ForksURL         string      `json:"forks_url"`
	KeysURL          string      `json:"keys_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	TeamsURL         string      `json:"teams_url"`
	HooksURL         string      `json:"hooks_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	EventsURL        string      `json:"events_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BranchesURL      string      `json:"branches_url"`
	TagsURL          string      `json:"tags_url"`
	BlobsURL         string      `json:"blobs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	TreesURL         string      `json:"trees_url"`
	StatusesURL      string      `json:"statuses_url"`
	LanguagesURL     string      `json:"languages_url"`
	StargazersURL    string      `json:"stargazers_url"`
	ContributorsURL  string      `json:"contributors_url"`
	SubscribersURL   string      `json:"subscribers_url"`
	SubscriptionURL  string      `json:"subscription_url"`
	CommitsURL       string      `json:"commits_url"`
	GitCommitsURL    string      `json:"git_commits_url"`
	CommentsURL      string      `json:"comments_url"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	ContentsURL      string      `json:"contents_url"`
	CompareURL       string      `json:"compare_url"`
	MergesURL        string      `json:"merges_url"`
	ArchiveURL       string      `json:"archive_url"`
	DownloadsURL     string      `json:"downloads_url"`
	IssuesURL        string      `json:"issues_url"`
	PullsURL         string      `json:"pulls_url"`
	MilestonesURL    string      `json:"milestones_url"`
	NotificationsURL string      `json:"notifications_url"`
	LabelsURL        string      `json:"labels_url"`
	ReleasesURL      string      `json:"releases_url"`
	DeploymentsURL   string      `json:"deployments_url"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	PushedAt         time.Time   `json:"pushed_at"`
	GitURL           string      `json:"git_url"`
	SSHURL           string      `json:"ssh_url"`
	CloneURL         string      `json:"clone_url"`
	SvnURL           string      `json:"svn_url"`
	Homepage         interface{} `json:"homepage"`
	Size             int         `json:"size"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Language         string      `json:"language"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasDownloads     bool        `json:"has_downloads"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	HasDiscussions   bool        `json:"has_discussions"`
	ForksCount       int         `json:"forks_count"`
	MirrorURL        interface{} `json:"mirror_url"`
	Archived         bool        `json:"archived"`
	Disabled         bool        `json:"disabled"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	License          struct {
		Key    string `json:"key"`
		Name   string `json:"name"`
		SpdxID string `json:"spdx_id"`
		URL    string `json:"url"`
		NodeID string `json:"node_id"`
	} `json:"license"`
	AllowForking             bool          `json:"allow_forking"`
	IsTemplate               bool          `json:"is_template"`
	WebCommitSignoffRequired bool          `json:"web_commit_signoff_required"`
	Topics                   []interface{} `json:"topics"`
	Visibility               string        `json:"visibility"`
	Forks                    int           `json:"forks"`
	OpenIssues               int           `json:"open_issues"`
	Watchers                 int           `json:"watchers"`
	DefaultBranch            string        `json:"default_branch"`
	Permissions              struct {
		Admin    bool `json:"admin"`
		Maintain bool `json:"maintain"`
		Push     bool `json:"push"`
		Triage   bool `json:"triage"`
		Pull     bool `json:"pull"`
	} `json:"permissions"`
	TempCloneToken            string `json:"temp_clone_token"`
	AllowSquashMerge          bool   `json:"allow_squash_merge"`
	AllowMergeCommit          bool   `json:"allow_merge_commit"`
	AllowRebaseMerge          bool   `json:"allow_rebase_merge"`
	AllowAutoMerge            bool   `json:"allow_auto_merge"`
	DeleteBranchOnMerge       bool   `json:"delete_branch_on_merge"`
	AllowUpdateBranch         bool   `json:"allow_update_branch"`
	UseSquashPrTitleAsDefault bool   `json:"use_squash_pr_title_as_default"`
	SquashMergeCommitMessage  string `json:"squash_merge_commit_message"`
	SquashMergeCommitTitle    string `json:"squash_merge_commit_title"`
	MergeCommitMessage        string `json:"merge_commit_message"`
	MergeCommitTitle          string `json:"merge_commit_title"`
	SecurityAndAnalysis       struct {
		SecretScanning struct {
			Status string `json:"status"`
		} `json:"secret_scanning"`
		SecretScanningPushProtection struct {
			Status string `json:"status"`
		} `json:"secret_scanning_push_protection"`
	} `json:"security_and_analysis"`
	NetworkCount     int `json:"network_count"`
	SubscribersCount int `json:"subscribers_count"`
}

type Output struct {
	URL string `json:"url"`
	NetScore float64 `json:"net_score"`
	RampUp float64 `json:"ramp_up"`
	Correctness float64 `json:"correctness"`
	BusFactor float64 `json:"bus_factor"`
	ResponsiveMaintainer float64 `json:"responsive_maintainer"`
	License float64 `json:"license"`
}

// get links for file
// input: string of URL_file
// output array of url strings
func getLinksFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var links []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		links = append(links, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return links, nil
}

// function for npm links

// function for Github repos


// function for metrics

// function for ndjson



// MOVE THIS TO UTILITY
func toNDJson(url string, ns float64, ru float64, c float64, bf float64, rm float64, l float64)(string, error) {
	j := Output{URL: url, NetScore: ns, RampUp: ru, Correctness: c, BusFactor: bf, ResponsiveMaintainer: rm, License: l}
	b, err := json.Marshal(j)
	if err != nil {
		log.Fatal("Error with NDJson conversion")
	}
	return string(b), nil
}

func getEnvVar(key string) string {
	err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	return os.Getenv(key)
}

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
		links, err := getLinksFromFile(args[0])
		if err != nil {
			log.Fatal("error opening input file")
			os.Exit(1)
		}

		for _, link := range links {
			parts := strings.Split(link, "/")
			url := link	// https://www.npmjs.com/package/browserify

			// if npmjs link, find GitHub repo. ADD message for npm modules with no GitHub repo
			if parts[2] == "www.npmjs.com" {
				packageName := parts[len(parts)-1]
				resp, err := http.Get("https://registry.npmjs.org/" + packageName)
				if err != nil {
					log.Fatal(err)
				}
				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}

				var npmResp NPMResponse
				err = json.Unmarshal(body, &npmResp)
				if err != nil {
					log.Fatal(err)
				}
				url = npmResp.Repository.URL
				url = "https://" + strings.Split(url, "//")[1]
				url = url[:len(url)-4] // removes .git
			} 

			github_parts := strings.Split(url, "/")
			github_api_url := "https://api.github.com/repos/" + github_parts[len(github_parts)-2] + "/" + github_parts[len(github_parts)-1]

		
			// process GitHub repo
			client := &http.Client{}
			req, err := http.NewRequest("GET", github_api_url, nil)
			if err != nil {
				log.Fatal(err)
			}
			req.Header.Set("Accept", "application/vnd.github+json")
			req.Header.Set("Authorization", "Bearer " + getEnvVar("PERSONAL_TOKEN"))
			req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			
			bodyText, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			var response GitHubResponse 
			json.Unmarshal(bodyText, &response)

			
				// COMPUTE METRICS
			correctScore := 1.0 - (0.000001 + (float64(response.OpenIssuesCount) / float64(response.StargazersCount)))
			j, err :=  toNDJson(link, 1.0, 1.0, correctScore, 1.0, 1.0, 1.0)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			_, err = fmt.Fprintln(os.Stdout, j)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
	
		os.Exit(0)
		// READ from a config file managed by config.go and calculate net score.
		// Call netscore function (weights[], scores[])
		// {"url":,"NetScore":,"RampUp":,"Correctness":,"BusFactor","ResponsiveMaintainer","License":}, 
		// product NDJSON output, build a struct for it
		// each score is between 0 and 1
		// exit 0 on success



		// check out https://api.github.com/repos/anthony-pei/ECE461 for more info
		// Ramp up: (# of lines of comments / # of lines of total code) / .2 
		// check contents_url, comments_url

		// Correctness 1 - (# of issues / # of stars)
		// fmt.Printf("star count: %v\n", response.StargazersCount) // check issues_url
		// fmt.Printf("open issues count: %v\n", response.OpenIssuesCount)
		// os.Exit(0)

			// Bus Factor: 1 - (1 / (# of contributors)) 
		// response.contributors_url + count 

		// Responsiveness: 1 - .01 * (open time (days) of last 10 closed issues)
		// check issues_url

		// License
		// response.license.name or .key, create a table of valid licenses
		
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
