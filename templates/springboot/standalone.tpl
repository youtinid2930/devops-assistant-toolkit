FROM {{.BaseImage}}
WORKDIR /app

COPY {{range .DependencyFiles}}{{.}} {{end}} ./
RUN {{.InstallCmd}}
COPY . .
RUN mvn package -DskipTests

{{- if .Optional.Security }}
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
{{- end}}

EXPOSE {{range .Ports}}{{.}} {{end}}

CMD ["java", "-jar", "target/app.jar"]

{{- if .Optional.Healthcheck }}
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:{{index .Ports 0}}/actuator/health || exit 1
{{- end}}