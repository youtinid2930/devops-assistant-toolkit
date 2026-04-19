# DevOps Assistant Toolkit

A CLI tool written in Go to speed up common DevOps setup tasks. Built with Cobra, it provides commands for generating Dockerfiles and scaffolding project structure for multiple stacks.

**Solo project** — built entirely by me to explore Go and CLI tooling beyond typical web development.

---

## Tech Stack

- **Language:** Go
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra)
- **Features:** Dockerfile generation, multi-stack templates, interactive wizard

---

## Commands

```
devops
└── docker
    └── init    # Initialize Docker setup for a project (interactive wizard)
```

### Usage

```bash
# Show help
./devops --help

# Initialize Docker setup interactively
./devops docker init
```

The `init` command launches an interactive wizard that asks about your project stack and generates the appropriate Dockerfile and configuration.

---

## Project Structure

```
├── main.go                   # Entry point — calls cli.Execute()
├── internal/
│   ├── cli/
│   │   ├── root.go           # Root command (devops)
│   │   ├── docker.go         # docker subcommand
│   │   ├── init.go           # init subcommand
│   │   └── wizard.go         # Interactive setup wizard
│   └── docker/               # Dockerfile generation logic
├── config/                   # Configuration files
├── templates/                # Stack-specific Dockerfile templates (Node, Go, etc.)
├── go.mod
└── go.sum
```

---

## Getting Started

### Prerequisites

- Go 1.21+

### Build

```bash
git clone https://github.com/youtinid2930/devops-assistant-toolkit.git
cd devops-assistant-toolkit
go build -o devops .
```

### Run

```bash
./devops --help
./devops docker init
```

---

## Why This Project?

Built specifically to practice Go and understand how CLI tools work at a systems level — moving beyond web frameworks into lower-level tooling. It also solves a real annoyance: writing Dockerfiles from scratch every time you start a new project.
