package file_handler

import (
	"bufio"
	"log"
	"net/url"
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

	for scanner.Scan() {
		link := scanner.Text()
		URL, err := url.Parse(link)
		if err != nil {
			log.Printf("Error with URL parse: %v\n", err) // If error continue reading file
			continue
		}
		owner, name := "", ""
		if URL.Host == "github.com" {
			log.Printf("Parsing github link %v\n", link)
			parts := strings.Split(URL.Path, "/")
			if len(parts) != 3 {
				log.Printf("URL.path malformed %v\n", URL.Path)
				continue
			} else {
				owner, name = parts[1], parts[2]
			}
		} else if URL.Host == "www.npmjs.com" {
			log.Printf("Parsing npm link %v\n", link)
			owner, name = ConvertNpmToGitHub(URL.Path)
		} else {
			log.Printf("Unkown URL host %v\n", link)
			continue
		}
		//fmt.Println(owner, name)
		ownerNames = append(ownerNames, github_util.OwnerName{Owner: owner, Name: name, Url: link})
	}
	return ownerNames
}
