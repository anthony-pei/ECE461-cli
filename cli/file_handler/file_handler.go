package file_handler

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/anthony-pei/ECE461/cli/github_util"
)

func GetOwnersNamesFromFile(filename string) []github_util.OwnerName {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ownerNames []github_util.OwnerName

	scanner := bufio.NewScanner(file)

	// Handle poorly formatted and malicious links
	for scanner.Scan() {
		link := scanner.Text()
		parts := strings.Split(link, "/")
		if parts[2] == "www.npmjs.com" {
			new_link := ConvertNpmToGitHub(link, parts[len(parts)-1])
			parts = strings.Split(new_link, "/")
		}

		ownerNames = append(ownerNames, github_util.OwnerName{Owner: parts[3], Name: parts[4], Url: link})
	}
	return ownerNames
}
