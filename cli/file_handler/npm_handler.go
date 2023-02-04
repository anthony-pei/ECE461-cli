package file_handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type NPMResponse struct {
	Repository struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"repository"`
}

// function for npm links
// if npmjs link, find GitHub repo. ADD message for npm modules with no GitHub repo
func ConvertNpmToGitHub(url string, packageName string) string {
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
	url = url[:len(url)-4] // removes .git at the end
	return url
}
