package docker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Prompt asks the user for input with a default value
func Prompt(msg, defaultVal string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg + ": ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return defaultVal
	}
	return input
}

// AutoDetectDeps returns dependency files from stack config that exist in the folder
func AutoDetectDeps(stackCfg map[string]interface{}) []string {
	depFiles := []string{}
	if files, ok := stackCfg["dependency_files"].([]interface{}); ok {
		for _, f := range files {
			file := fmt.Sprintf("%v", f)
			if _, err := os.Stat(file); err == nil {
				depFiles = append(depFiles, file)
			}
		}
	}
	return depFiles
}