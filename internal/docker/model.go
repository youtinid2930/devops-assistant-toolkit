package docker

import (
	"fmt"
	"strings"
)

//
// Dockerfile
// ------------------------------------
// Represents the full Dockerfile (can contain multiple stages)
//
type Dockerfile struct {
	Stages []Stage
}

//
// Stage
// ------------------------------------
// Represents one build stage
//
type Stage struct {
	Name      string
	BaseImage string

	Args     map[string]string
	Envs     map[string]string
	WorkDir  string
	Copies   []Copy
	Runs     []string
	Exposes  []int
	User     string
	Entrypoint string
	Cmd      string
}

//
// Copy instruction
//
type Copy struct {
	Source      string
	Destination string
}

//
// NewStage
// ------------------------------------
// Creates a new stage with initialized maps
//
func NewStage(name, baseImage string) Stage {
	return Stage{
		Name:      name,
		BaseImage: baseImage,
		Args:      make(map[string]string),
		Envs:      make(map[string]string),
	}
}

//
// AddARG
//
func (s *Stage) AddARG(key, value string) {
	s.Args[key] = value
}

//
// AddENV
//
func (s *Stage) AddENV(key, value string) {
	s.Envs[key] = value
}

//
// AddRUN
//
func (s *Stage) AddRUN(cmd string) {
	s.Runs = append(s.Runs, cmd)
}

//
// AddCOPY
//
func (s *Stage) AddCOPY(src, dest string) {
	s.Copies = append(s.Copies, Copy{
		Source:      src,
		Destination: dest,
	})
}

//
// AddEXPOSE
//
func (s *Stage) AddEXPOSE(port int) {
	s.Exposes = append(s.Exposes, port)
}

//
// Render
// ------------------------------------
// Converts Dockerfile struct into string
//
func (d *Dockerfile) Render() string {
	var builder strings.Builder

	for _, stage := range d.Stages {

		// FROM
		if stage.Name != "" {
			builder.WriteString(fmt.Sprintf("FROM %s AS %s\n", stage.BaseImage, stage.Name))
		} else {
			builder.WriteString(fmt.Sprintf("FROM %s\n", stage.BaseImage))
		}

		// ARG
		for k, v := range stage.Args {
			builder.WriteString(fmt.Sprintf("ARG %s=%s\n", k, v))
		}

		// ENV
		for k, v := range stage.Envs {
			builder.WriteString(fmt.Sprintf("ENV %s=%s\n", k, v))
		}

		// WORKDIR
		if stage.WorkDir != "" {
			builder.WriteString(fmt.Sprintf("WORKDIR %s\n", stage.WorkDir))
		}

		// COPY
		for _, c := range stage.Copies {
			builder.WriteString(fmt.Sprintf("COPY %s %s\n", c.Source, c.Destination))
		}

		// RUN (grouped for best practice)
		if len(stage.Runs) > 0 {
			builder.WriteString("RUN ")
			builder.WriteString(strings.Join(stage.Runs, " && \\\n    "))
			builder.WriteString("\n")
		}

		// EXPOSE
		for _, port := range stage.Exposes {
			builder.WriteString(fmt.Sprintf("EXPOSE %d\n", port))
		}

		// USER
		if stage.User != "" {
			builder.WriteString(fmt.Sprintf("USER %s\n", stage.User))
		}

		// ENTRYPOINT
		if stage.Entrypoint != "" {
			builder.WriteString(fmt.Sprintf("ENTRYPOINT %s\n", stage.Entrypoint))
		}

		// CMD
		if stage.Cmd != "" {
			builder.WriteString(fmt.Sprintf("CMD %s\n", stage.Cmd))
		}

		builder.WriteString("\n")
	}

	return builder.String()
}