{{- if .Optional.MultiStage }}
# Multi-stage build
FROM {{.BaseImage}} AS builder
WORKDIR /app
COPY {{range .DependencyFiles}}{{.}} {{end}} ./
RUN {{.InstallCmd}}
COPY . .
RUN go build -o app
{{- end}}

{{- if .Optional.MultiStage }}
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/app .
{{- else }}
FROM {{.BaseImage}}
WORKDIR /app
COPY . .
RUN {{.InstallCmd}}
RUN go build -o app
{{- end}}

{{- if .Optional.Security }}
# Security: run as non-root
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
{{- end}}

EXPOSE {{range .Ports}}{{.}} {{end}}

CMD ["./app"]

{{- if .Optional.Healthcheck }}
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:{{index .Ports 0}}/ || exit 1
{{- end}}