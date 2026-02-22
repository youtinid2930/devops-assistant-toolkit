package docker

// OptionalBlocks represents optional Dockerfile sections
type OptionalBlocks struct {
	MultiStage  bool
	Security    bool
	HealthCheck bool
}

// DockerConfig represents a full Dockerfile configuration
type DockerConfig struct {
	BaseImage       string
	DependencyFiles string
	InstallCmd      string
	StartCmd        string
	Port            string
	Blocks          OptionalBlocks
	StackName       string
}