# Production Node Dockerfile

{{- if .Optional.MultiStage }}
# Build stage
FROM {{ .BaseImage }} AS builder
WORKDIR /app

{{- range .DependencyFiles }}
COPY {{ . }} .
{{- end }}

RUN {{ .InstallCmd }}

COPY . .

# Final stage
FROM {{ .BaseImage }}
WORKDIR /app
COPY --from=builder /app .

{{- else }}
FROM {{ .BaseImage }}
WORKDIR /app

{{- range .DependencyFiles }}
COPY {{ . }} .
{{- end }}

RUN {{ .InstallCmd }}

COPY . .
{{- end }}

# Expose ports
{{- range .Ports }}
EXPOSE {{ . }}
{{- end }}

CMD ["sh", "-c", "{{ .StartCmd }}"]