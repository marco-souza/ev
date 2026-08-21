# Tasks

## Minimum Viable Product

### Manage Vault

- [x] `ev init`
- [x] `ev drop`

### Manage Variables & Secrets

- [x] `ev set [NAME] [VALUE]`
- [x] `ev get [NAME]`
- [x] `ev rm [NAME]`

### Encrypt secrets

- [ ] ask encryption password when initializing
- [ ] encrypt/decrypt secreets to persist based on encryption password (no file saved)

### Run commands

- [x] `ev run -- [your application start command]` - implement secret expansion

## Minimum Lovable Product

### Activation

- [ ] `. <(ev activate)` - activate environment injecting secrets and var envs

### Support Multi environments

- [ ] `ev get|set|rm|ls --env=<dev|staging|production>`
- [ ] `ev run --env=<dev|staging|production> -- [commands]`

### Security hardening

- [ ] ensure we don't log the execution to bash history
- [x] remove secrets and variables from output (replace with `*****`)
