package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Check command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <project_name>")
		return
	}

	projectName := os.Args[1]

	// Create project folder
	err := os.Mkdir(projectName, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Error creating project folder:", err)
		return
	}

	// List template to copy
	templates := []string{"Dockerfile", "docker-compose.yml"}


	for _, tpl := range templates {
		copyTemplate("templates/"+tpl, projectName+"/"+tpl, projectName)
	}

	fmt.Println("Project", projectName, "initialized successfuly!")
}


// copyTemplate reads a template file, replaces placeholder, and write to destination
func copyTemplate (srcpath, destPath, projectName string) {
	data, err := ioutil.ReadFile(srcpath)

	if err != nil {
		fmt.Println("Error reading template:", err)
		return
	}

	content := strings.ReplaceAll(string(data), "{{PROJECT_NAME}}", projectName)

	err = ioutil.WriteFile(destPath, []byte(content), 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Created:", destPath)
}