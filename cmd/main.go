package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"

	"github.com/arsharaj/project-logger/config"
	"github.com/arsharaj/project-logger/parser"
	"github.com/arsharaj/project-logger/tailer"
)

func main() {
	fmt.Println("project-logger application starting...")

	// Load config
	configPath, _ := filepath.Abs(".")
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	fmt.Printf("configurations loaded: %+v\n", cfg)

	// Start tailing each log file
	for _, file := range cfg.LogFiles {
		go func(filePath string) {
			err := tailer.TailFile(filePath, func(line string) {
				entry := parser.ParseSyslogLine(line, filePath)

				jsonEntry, _ := json.Marshal(entry)
				fmt.Println(string(jsonEntry))
			})

			if err != nil {
				log.Printf("Error tailing %s: %v", filePath, err)
			}
		}(file)
	}

	// Keep the main goroutine alive
	select {}
}
