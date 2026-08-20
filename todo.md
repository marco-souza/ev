# Tasks

## Minimum Viable Product

### Manage Vault

- [x] `ev init`
- [x] `ev drop`

### Manage Variables & Secrets

- [x] `ev set [NAME] [VALUE]`
- [x] `ev get [NAME]`
- [x] `ev rm [NAME]`

### Run commands

- [ ] `ev run -- [your application start command]` - implement secret expantion

## Minimum Lovable Product

### Support Multi environments

- [ ] `ev get|set|rm|ls --env=<dev|staging|production>`
- [ ] `ev run --env=<dev|staging|production> -- [commands]`

### Security hardening

- [ ] ensure we don't log the execution to bash history
- [ ] remove secrets and variables from output (replace with `*****`)
