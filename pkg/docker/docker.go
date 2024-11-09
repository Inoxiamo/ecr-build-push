package docker

import (
	"os"
	"os/exec"
)

// Build costruisce un'immagine Docker usando i parametri specificati, variabili d'ambiente, argomenti e un contesto di build.
func Build(imageName, dockerFilePath, buildContext string, envVars, argVars map[string]string) error {
	args := []string{"build", "-t", imageName}

	// Aggiunge gli argomenti specifici di Docker come argomenti del comando
	for key, value := range argVars {
		args = append(args, "--build-arg", key+"="+value)
	}

	// Aggiunge le variabili d'ambiente come argomenti del comando
	for key, value := range envVars {
		args = append(args, "--build-arg", key+"="+value)
	}

	// Completa il comando specificando il percorso del Dockerfile e il contesto di build
	args = append(args, "-f", dockerFilePath, buildContext)

	// Esegue il comando Docker
	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout // Imposta correttamente l'output standard
	cmd.Stderr = os.Stderr // Imposta correttamente l'output degli errori
	return cmd.Run()
}

// Push carica un'immagine Docker al registro specificato.
func Push(imageName string) error {
	cmd := exec.Command("docker", "push", imageName)
	cmd.Stdout = os.Stdout // Imposta correttamente l'output standard
	cmd.Stderr = os.Stderr // Imposta correttamente l'output degli errori
	return cmd.Run()
}
