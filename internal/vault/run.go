package vault

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/marco-souza/ev/internal/repository"
)

func (v *Vault) Run(db *sql.DB, shCmd string) (string, error) {
	cmd := exec.Command("bash", "-c", shCmd)

	cmd.Env = os.Environ()

	ctx := context.Background()
	varsRepo := repository.NewVariableRepository(db)
	secretsRepo := repository.NewSecretRepository(db)

	var injectedValues []string

	vars, err := varsRepo.ListAll(ctx)
	if err != nil {
		return "", err
	}

	for _, variable := range vars {
		env := fmt.Sprintf("%s=%s", variable.Name, variable.Value)
		cmd.Env = append(cmd.Env, env)
		injectedValues = append(injectedValues, variable.Value)
	}

	secrets, err := secretsRepo.ListAll(ctx)
	if err != nil {
		return "", err
	}

	for _, secret := range secrets {
		env := fmt.Sprintf("%s=%s", secret.Name, secret.Value)
		cmd.Env = append(cmd.Env, env)
		injectedValues = append(injectedValues, secret.Value)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	parsedOutput := string(output)
	secretPlaceholder := "*****"

	for _, value := range injectedValues {
		parsedOutput = strings.ReplaceAll(parsedOutput, value, secretPlaceholder)
	}

	return parsedOutput, nil
}
