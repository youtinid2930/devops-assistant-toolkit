package docker

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// DockerfileGenerator handles generating Dockerfiles from templates
type DockerfileGenerator struct {
	Stacks    map[string]map[string]interface{}
	Config    DockerConfig
	Template  string
	Output    string
	StackName string
}

// NewGenerator creates a new generator instance
func NewGenerator(stacks map[string]map[string]interface{}, templatePath, outputPath string) *DockerfileGenerator {
	return &DockerfileGenerator{
		Stacks:   stacks,
		Template: templatePath,
		Output:   outputPath,
	}
}

// SelectStack prompts the user to select a stack and sets the Config
func (g *DockerfileGenerator) SelectStack() error {
	// List stacks
	println("Available stacks:")
	for s := range g.Stacks {
		println("-", s)
	}

	stack := Prompt("Choose stack (default: node)", "node")
	stackCfg, ok := g.Stacks[stack]
	if !ok {
		return os.ErrNotExist
	}
	g.StackName = stack

	// Auto-detect dependency files
	depFiles := AutoDetectDeps(stackCfg)
	if len(depFiles) == 0 {
		// fallback: use all listed files
		for _, f := range stackCfg["dependency_files"].([]interface{}) {
			depFiles = append(depFiles, f.(string))
		}
	}

	// Optional blocks
	security := Prompt("Include security best practices? [Y/n]", "Y") == "Y"
	multiStage := Prompt("Include multi-stage build? [Y/n]", "N") == "Y"
	healthcheck := Prompt("Include healthcheck? [Y/n]", "N") == "Y"

	// Set config
	g.Config = DockerConfig{
		BaseImage:       stackCfg["base"].(string),
		DependencyFiles: strings.Join(depFiles, " "),
		InstallCmd:      stackCfg["install_cmd"].(string),
		StartCmd:        stackCfg["start_cmd"].(string),
		Port:            strings.TrimSpace(strings.Split(fmt.Sprintf("%v", stackCfg["ports"].([]interface{})[0]), " ")[0]),
		Blocks: OptionalBlocks{
			Security:    security,
			MultiStage:  multiStage,
			HealthCheck: healthcheck,
		},
		StackName: stack,
	}

	return nil
}

// Generate creates the Dockerfile from template
func (g *DockerfileGenerator) Generate() error {
	tpl, err := template.ParseFiles(g.Template)
	if err != nil {
		return err
	}

	f, err := os.Create(g.Output)
	if err != nil {
		return err
	}
	defer f.Close()

	return tpl.Execute(f, g.Config)
}