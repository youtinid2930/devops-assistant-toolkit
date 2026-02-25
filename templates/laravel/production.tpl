{{- if .Optional.MultiStage }}
# Multi-stage build
FROM {{.BaseImage}} AS builder
WORKDIR /app
COPY {{range .DependencyFiles}}{{.}} {{end}} ./
RUN {{.InstallCmd}}
COPY . .
{{- end}}

{{- if .Optional.MultiStage }}
FROM php:8.2-fpm
WORKDIR /app
COPY --from=builder /app ./
{{- else }}
FROM {{.BaseImage}}
WORKDIR /app
COPY . .
RUN {{.InstallCmd}}
{{- end}}

{{- if .Optional.Security }}
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
{{- end}}

EXPOSE {{range .Ports}}{{.}} {{end}}

CMD ["sh", "-c", "{{.StartCmd}}"]

{{- if .Optional.Healthcheck }}
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:{{index .Ports 0}}/ || exit 1
{{- end}}