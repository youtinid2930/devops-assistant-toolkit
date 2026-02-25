package docker

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// TemplateData combines stack config + optional blocks
type TemplateData struct {
	BaseImage       string
	DependencyFiles []string
	InstallCmd      string
	StartCmd        string
	Ports           []int
	Optional        struct {
		MultiStage  bool
		Security    bool
		Healthcheck bool
	}
}

// GenerateDockerfile renders a Dockerfile using a template and optional blocks
func GenerateDockerfile(stackName string, cfg StackConfig, templateName string, optionalBlocks []string) error {
	// Prepare template data
	data := TemplateData{
		BaseImage:       cfg.BaseImage,
		DependencyFiles: cfg.DependencyFiles,
		InstallCmd:      cfg.InstallCmd,
		StartCmd:        cfg.StartCmd,
		Ports:           cfg.Ports,
	}

	for _, block := range optionalBlocks {
		switch block {
		case "multi_stage":
			data.Optional.MultiStage = true
		case "security":
			data.Optional.Security = true
		case "healthcheck":
			data.Optional.Healthcheck = true
		}
	}

	// Load template file for the selected stack
	templatePath := filepath.Join("templates", stackName, templateName)
	tplBytes, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to load template %q: %w", templatePath, err)
	}

	tpl, err := template.New(templateName).Parse(string(tplBytes))
	if err != nil {
		return fmt.Errorf("failed to parse template %q: %w", templatePath, err)
	}

	// Render template
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to execute template %q: %w", templatePath, err)
	}

	// Write the generated Dockerfile
	if err := os.WriteFile("Dockerfile", buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write Dockerfile: %w", err)
	}

	return nil
}