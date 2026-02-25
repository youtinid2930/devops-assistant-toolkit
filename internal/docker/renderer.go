package docker

import (
	"fmt"
	"sort"
	"strings"
)

func RenderDockerfile(df Dockerfile) string {
	var builder strings.Builder

	for _, stage := range df.Stages {

		// FROM (with optional stage name)
		if stage.Name != "" {
			builder.WriteString(fmt.Sprintf("FROM %s AS %s\n", stage.BaseImage, stage.Name))
		} else {
			builder.WriteString(fmt.Sprintf("FROM %s\n", stage.BaseImage))
		}

		// ---- ARG (sorted for deterministic output)
		if len(stage.Args) > 0 {
			keys := sortedKeys(stage.Args)
			for _, k := range keys {
				builder.WriteString(fmt.Sprintf("ARG %s=%s\n", k, stage.Args[k]))
			}
		}

		// ---- ENV (sorted for deterministic output)
		if len(stage.Envs) > 0 {
			keys := sortedKeys(stage.Envs)
			for _, k := range keys {
				builder.WriteString(fmt.Sprintf("ENV %s=%s\n", k, stage.Envs[k]))
			}
		}

		// ---- WORKDIR
		if stage.WorkDir != "" {
			builder.WriteString(fmt.Sprintf("WORKDIR %s\n", stage.WorkDir))
		}

		// ---- COPY
		for _, c := range stage.Copies {
			builder.WriteString(fmt.Sprintf("COPY %s %s\n", c.Source, c.Destination))
		}

		// ---- RUN (grouped best practice)
		if len(stage.Runs) > 0 {
			builder.WriteString("RUN ")
			builder.WriteString(strings.Join(stage.Runs, " && \\\n    "))
			builder.WriteString("\n")
		}

		// ---- EXPOSE
		for _, port := range stage.Exposes {
			builder.WriteString(fmt.Sprintf("EXPOSE %d\n", port))
		}

		// ---- USER
		if stage.User != "" {
			builder.WriteString(fmt.Sprintf("USER %s\n", stage.User))
		}

		// ---- ENTRYPOINT
		if stage.Entrypoint != "" {
			builder.WriteString(fmt.Sprintf("ENTRYPOINT %s\n", stage.Entrypoint))
		}

		// ---- CMD
		if stage.Cmd != "" {
			builder.WriteString(fmt.Sprintf("CMD %s\n", stage.Cmd))
		}

		builder.WriteString("\n")
	}

	return builder.String()
}

// helper for deterministic map ordering
func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}