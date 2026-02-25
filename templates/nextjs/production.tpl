{{- if .Optional.MultiStage }}
FROM {{.BaseImage}} AS builder
WORKDIR /app
COPY {{range .DependencyFiles}}{{.}} {{end}} ./
RUN {{.InstallCmd}}
COPY . .
RUN npm run build
{{- end}}

{{- if .Optional.MultiStage }}
FROM {{.BaseImage}}
WORKDIR /app
COPY --from=builder /app ./
{{- else }}
FROM {{.BaseImage}}
WORKDIR /app
COPY . .
RUN {{.InstallCmd}}
RUN npm run build
{{- end}}

{{- if .Optional.Security }}
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
{{- end}}

EXPOSE {{range .Ports}}{{.}} {{end}}

CMD ["npm", "start"]

{{- if .Optional.Healthcheck }}
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:{{index .Ports 0}}/ || exit 1
{{- end}}