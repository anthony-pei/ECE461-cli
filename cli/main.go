/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/anthony-pei/ECE461/cli/cmd"
	log "github.com/sirupsen/logrus"
)

func logging_init() {
	lf, exists := os.LookupEnv("LOG_FILE")
	if !exists || len(lf) == 0 {
		log.SetOutput(io.Discard)
		return
	}
	level, exists := os.LookupEnv("LOG_LEVEL")
	if !exists {
		level = "0"
	}
	switch level {
	default:
		log.SetOutput(io.Discard)
		return
	case "1":
		log.SetLevel(log.InfoLevel)
	case "2":
		log.SetLevel(log.DebugLevel)
	}

	file, err := os.OpenFile(lf, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening LOG_FILE", err)
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func main() {
	logging_init()
	cmd.Execute()
}
