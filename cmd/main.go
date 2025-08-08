package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/arsharaj/project-logger/config"
)

func main() {
	fmt.Println("project-logger application starting...")

	configPath, _ := filepath.Abs(".")
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	fmt.Printf("configurations loaded: %+v\n", cfg)
}
