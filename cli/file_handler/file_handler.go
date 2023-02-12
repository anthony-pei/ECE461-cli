package file_handler

import (
	"bufio"
	"net/url"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/anthony-pei/ECE461/cli/github_util"
)

func GetOwnersNamesFromFile(filename string) []github_util.OwnerName {
	file, err := os.Open(filename)
	if err != nil {
		log.Debug("Error getting name from file")
	}
	defer file.Close()

	var ownerNames []github_util.OwnerName

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		link := scanner.Text()
		URL, err := url.Parse(link)
		if err != nil {
			log.Debugf("Error with URL parse: %v", err) // If error continue reading file
			continue
		}
		owner, name := "", ""
		if URL.Host == "github.com" {
			log.Infof("Parsing github link %v", link)
			parts := strings.Split(URL.Path, "/")
			if len(parts) != 3 {
				log.Infof("URL.path malformed %v", URL.Path)
				continue
			} else {
				owner, name = parts[1], parts[2]
			}
		} else if URL.Host == "www.npmjs.com" {
			log.Infof("Parsing npm link %v", link)
			owner, name = ConvertNpmToGitHub(URL.Path)
		} else {
			log.Debugf("Unkown URL host %v", link)
			continue
		}
		//fmt.Println(owner, name)
		ownerNames = append(ownerNames, github_util.OwnerName{Owner: owner, Name: name, Url: link})
	}
	return ownerNames
}
