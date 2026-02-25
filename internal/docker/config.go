package docker

import (
	"encoding/json"
	"os"
)

type StackConfig struct {
	BaseImage       string   `json:"base_image"`
	DependencyFiles []string `json:"dependency_files"`
	InstallCmd      string   `json:"install_cmd"`
	StartCmd        string   `json:"start_cmd"`
	Ports           []int    `json:"ports"`
	Templates       []string `json:"templates"`
	OptionalBlocks  []string `json:"optional_blocks"`
}

type Stacks map[string]StackConfig

func LoadStacks(path string) (Stacks, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var stacks Stacks
	err = json.Unmarshal(data, &stacks)
	if err != nil {
		return nil, err
	}

	return stacks, nil
}