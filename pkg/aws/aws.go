package aws

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetDockerLoginCommand restituisce il comando di login a Docker utilizzando le credenziali AWS ECR.
func GetDockerLoginCommand(region, profile, awsAccountID string) (string, error) {
	cmd := exec.Command("aws", "ecr", "get-login-password", "--region", region, "--profile", profile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("errore nel recupero della password di login: %s, %v", string(output), err)
	}

	// Pulisci il token di autenticazione
	password := strings.TrimSpace(string(output))

	// Formatta il comando completo di login Docker
	loginCommand := fmt.Sprintf("echo %s | docker login --username AWS --password-stdin %s.dkr.ecr.%s.amazonaws.com", password, awsAccountID, region)
	return loginCommand, nil
}

func ExecuteCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
