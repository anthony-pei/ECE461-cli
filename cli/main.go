/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/anthony-pei/ECE461/cli/cmd"
	log "github.com/sirupsen/logrus"
	"os"
)



func main() {
	lf,e :=  os.LookupEnv("LOG_FILE")
	if e {
		log.Fatal(e)
	}
	file, err := os.OpenFile(lf, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel)
	cmd.Execute()
}
