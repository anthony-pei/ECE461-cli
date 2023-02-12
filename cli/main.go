/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/anthony-pei/ECE461/cli/cmd"
	log "github.com/sirupsen/logrus"
)



func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel)
	cmd.Execute()
}
