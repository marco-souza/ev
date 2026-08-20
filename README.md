# Environment Vault CLI

A small cli tool for managing environment variables and secrets across projects.

This project values:

- offline first, sync if needed
- security as a first class citizen
- handwritten code, for control over complexity

## For Individuals

- Source of truth: Share secrets across projects
- Agentic Secrets: Securely execute agent's code without exposing credentials
- Encrypted at rest: No data exposure without master passphrase

## Usage

```sh replace-with: go run main.go
# Initialize project
ev init

# Variables command
ev set [NAME] [VALUE]
ev get [NAME]
ev rm  [NAME]

# Secrets command
ev set -s [NAME] [VALUE]
ev get -s [NAME]
ev rm  -s [NAME]

# Drop vault
ev drop
```
