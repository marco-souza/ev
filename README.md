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
$ go run main.go
A vault for your secrets

Usage:
  ev [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        initialize a vault
  sec         secrets command
  var         variables command

Flags:
  -h, --help   help for ev

Use "ev [command] --help" for more information about a command.
```
