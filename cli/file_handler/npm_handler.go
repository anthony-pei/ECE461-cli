package file_handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"

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
func ConvertNpmToGitHub(path string) (string, string) {
	parts := strings.Split(path, "/")
	resp, err := http.Get("https://registry.npmjs.org/" + parts[len(parts)-1])
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	var npmResp NPMResponse
	err = json.Unmarshal(body, &npmResp)
	if err != nil {
		log.Error(err)
	}
	git_url := npmResp.Repository.URL
	URL, err := url.Parse(git_url)
	if err != nil {
		log.Errorf("Error translating npm to git %v", err)
	}

	parts = strings.Split(URL.Path, "/")
	if len(parts) != 3 {
		log.Errorf("Malformed path translating npm to git %v", URL.Path)
		return "", ""
	} else {
		return parts[1], strings.ReplaceAll(parts[2], ".git", "")
	}
}
