package main

import (
	"encoding/json"
	"fmt"
	"os"
	"devops-assistant-toolkit/internal/docker"
)

func main() {
	// Load stacks.json
	stacks, err := loadStacksConfig("config/stacks.json")
	if err != nil {
		fmt.Println("Error loading stacks:", err)
		return
	}

	gen := docker.NewGenerator(stacks, "templates/Dockerfile.tpl", "Dockerfile")

	if err := gen.SelectStack(); err != nil {
		fmt.Println("Error selecting stack:", err)
		return
	}

	if err := gen.Generate(); err != nil {
		fmt.Println("Error generating Dockerfile:", err)
		return
	}

	fmt.Println("Dockerfile generated successfully!")
}

// loadStacksConfig loads JSON stack configuration
func loadStacksConfig(path string) (map[string]map[string]interface{}, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var stacks map[string]map[string]interface{}
	if err := json.Unmarshal(file, &stacks); err != nil {
		return nil, err
	}

	return stacks, nil
}