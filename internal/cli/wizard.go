package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"devops-assistant-toolkit/internal/docker"
)

// ------------------------------
// Prompt helper
// ------------------------------
func Prompt(message string, defaultVal string) string {
	var input string
	fmt.Printf("%s [%s]: ", message, defaultVal)
	fmt.Scanln(&input)
	input = strings.TrimSpace(input)
	if input == "" {
		return defaultVal
	}
	return input
}

// ------------------------------
// Numbered selection helper
// ------------------------------
func SelectFromList(prompt string, options []string, defaultIndex int) int {
	for i, opt := range options {
		fmt.Printf("%d) %s\n", i+1, opt)
	}
	fmt.Printf("%s [%d]: ", prompt, defaultIndex+1)

	var input string
	fmt.Scanln(&input)

	if input == "" {
		return defaultIndex
	}

	num, err := strconv.Atoi(input)
	if err != nil || num < 1 || num > len(options) {
		fmt.Println("Invalid choice, using default.")
		return defaultIndex
	}
	return num - 1
}

// ------------------------------
// SelectStack using numbered options
// ------------------------------
func SelectStack(stacks docker.Stacks) (string, docker.StackConfig, error) {
	if len(stacks) == 0 {
		return "", docker.StackConfig{}, fmt.Errorf("no stacks available")
	}

	stacksList := []string{}
	for name := range stacks {
		stacksList = append(stacksList, name)
	}
	sort.Strings(stacksList)

	index := SelectFromList("Choose stack", stacksList, 0)
	stackName := stacksList[index]
	stackConfig := stacks[stackName]

	return stackName, stackConfig, nil
}

// ------------------------------
// ListTemplates helper
// ------------------------------
func ListTemplates(stackName string) ([]string, error) {
	dir := filepath.Join("templates", stackName)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read template folder for stack %s: %w", stackName, err)
	}

	templates := []string{}
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".tpl") {
			templates = append(templates, e.Name())
		}
	}

	if len(templates) == 0 {
		return nil, fmt.Errorf("no templates found for stack %s", stackName)
	}

	sort.Strings(templates)
	return templates, nil
}

// ------------------------------
// SelectTemplate using numbered options
// ------------------------------
func SelectTemplate(stackName string) (string, error) {
	templates, err := ListTemplates(stackName)
	if err != nil {
		return "", err
	}

	index := SelectFromList("Choose template", templates, 0)
	return templates[index], nil
}

// ------------------------------
// SelectOptionalBlocks using numbered multi-select
// ------------------------------
func SelectOptionalBlocks(config docker.StackConfig) []string {
	if len(config.OptionalBlocks) == 0 {
		return nil
	}

	fmt.Println("\nOptional blocks available (comma-separated numbers):")
	for i, block := range config.OptionalBlocks {
		fmt.Printf("%d) %s\n", i+1, block)
	}

	fmt.Print("Choose optional blocks (e.g., 1,3) or leave empty: ")
	var input string
	fmt.Scanln(&input)

	selectedBlocks := []string{}
	if input != "" {
		choices := strings.Split(input, ",")
		for _, c := range choices {
			num, err := strconv.Atoi(strings.TrimSpace(c))
			if err == nil && num >= 1 && num <= len(config.OptionalBlocks) {
				selectedBlocks = append(selectedBlocks, config.OptionalBlocks[num-1])
			}
		}
	}

	return selectedBlocks
}

// ------------------------------
// RunDockerWizard
// ------------------------------
func RunDockerWizard() error {
	stacks, err := docker.LoadStacks(filepath.Join("config", "stacks.json"))
	if err != nil {
		return err
	}

	stackName, stackConfig, err := SelectStack(stacks)
	if err != nil {
		return err
	}

	template, err := SelectTemplate(stackName)
	if err != nil {
		return err
	}

	optionalBlocks := SelectOptionalBlocks(stackConfig)

	fmt.Println("\nSummary:")
	fmt.Println("Stack:", stackName)
	fmt.Println("Template:", template)
	fmt.Println("Optional blocks:", optionalBlocks)

	// Generate Dockerfile using template
	if err := docker.GenerateDockerfile(stackName, stackConfig, template, optionalBlocks); err != nil {
		return fmt.Errorf("failed to generate Dockerfile: %w", err)
	}

	fmt.Println("\nDockerfile generated successfully!")
	return nil
}