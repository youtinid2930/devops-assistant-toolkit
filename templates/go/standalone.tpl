FROM {{.BaseImage}}
WORKDIR /app

COPY {{range .DependencyFiles}}{{.}} {{end}} ./
RUN {{.InstallCmd}}
COPY . .
RUN go build -o app

{{- if .Optional.Security }}
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
{{- end}}

EXPOSE {{range .Ports}}{{.}} {{end}}

CMD ["./app"]

{{- if .Optional.Healthcheck }}
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:{{index .Ports 0}}/ || exit 1
{{- end}}