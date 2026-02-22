{{- if .Blocks.MultiStage}}

{{- if eq .StackName "node" }}
# ---- Node multi-stage ----
# Build stage
FROM {{.BaseImage}} AS builder
WORKDIR /app
COPY {{.DependencyFiles}} ./
RUN {{.InstallCmd}}
COPY . .

# Final stage
FROM node:20-alpine
WORKDIR /app
COPY --from=builder /app ./
EXPOSE {{.Port}}

{{- if .Blocks.Security}}
# Security block: non-root user
RUN adduser -D appuser
USER appuser
{{- end}}

{{- if .Blocks.HealthCheck}}
# Healthcheck block
HEALTHCHECK CMD curl --fail http://localhost:{{.Port}} || exit 1
{{- end}}

CMD {{.StartCmd}}

{{- else if eq .StackName "go" }}
# ---- Go multi-stage ----
# Build stage
FROM {{.BaseImage}} AS builder
WORKDIR /app
COPY {{.DependencyFiles}} ./
RUN {{.InstallCmd}}
COPY . .
RUN go build -o app

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE {{.Port}}

{{- if .Blocks.Security}}
# Security block: non-root user
RUN adduser -D appuser
USER appuser
{{- end}}

{{- if .Blocks.HealthCheck}}
# Healthcheck block
HEALTHCHECK CMD curl --fail http://localhost:{{.Port}} || exit 1
{{- end}}

CMD ["./app"]
{{- end}}

{{- else }}

# ---- Single-stage build ----
# Base image
FROM {{.BaseImage}}

# Set working directory
WORKDIR /app

# Copy dependency files first for caching
COPY {{.DependencyFiles}} ./

# Install dependencies
RUN {{.InstallCmd}}

# Copy the rest of the project
COPY . .

# Expose port
EXPOSE {{.Port}}

{{- if .Blocks.Security}}
# Security block: non-root user
RUN adduser -D appuser
USER appuser
{{- end}}

{{- if .Blocks.HealthCheck}}
# Healthcheck block
HEALTHCHECK CMD curl --fail http://localhost:{{.Port}} || exit 1
{{- end}}

# Start command
CMD {{.StartCmd}}

{{- end}}