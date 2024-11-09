package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Inoxiamo/ecr-build-push/pkg/aws"
	"github.com/Inoxiamo/ecr-build-push/pkg/config"
	"github.com/Inoxiamo/ecr-build-push/pkg/docker"
)

func main() {
	// Controlla se l'ambiente è passato come argomento
	if len(os.Args) < 2 {
		log.Fatal("Specifica l'ambiente come argomento (es. DEV, PROD)")
	}
	envName := os.Args[1]

	// Carica la configurazione specifica per l'ambiente
	envConfig, err := config.LoadConfigForEnv("config-ebp.json", envName)
	if err != nil {
		log.Fatalf("Errore nel caricamento della configurazione: %v", err)
	}

	// Esegue il comando di login a Docker ECR
	loginCmd, err := aws.GetDockerLoginCommand(envConfig.AWSRegion, envConfig.AWSProfile, envConfig.AWSAccountID)
	if err != nil {
		log.Fatalf("Errore nel recuperare il comando di login Docker: %v", err)
	}

	fmt.Println("Esecuzione del comando di login ECR...")
	output, err := exec.Command("sh", "-c", loginCmd).CombinedOutput()
	if err != nil {
		log.Fatalf("Errore nel login a Docker: %v, Output: %s", err, output)
	}
	fmt.Println("Login ECR eseguito con successo.")

	// Costruisci l'immagine Docker
	fmt.Println("Costruzione dell'immagine Docker in corso...")
	fullName := getImageFullName(envConfig.AWSAccountID, envConfig.AWSRegion, envConfig.DockerImageName, envConfig.DockerImageTag)
	buildErr := docker.Build(fullName, envConfig.DockerFilePath, envConfig.PathContext, envConfig.DockerEnvVars, envConfig.DockerArgVars)
	if buildErr != nil {
		log.Fatalf("Errore nella costruzione dell'immagine Docker: %v", buildErr)
	}
	fmt.Println("Costruzione dell'immagine completata.")

	// Push dell'immagine a ECR
	fmt.Println("Push dell'immagine a ECR in corso...")
	pushErr := docker.Push(fullName)
	if pushErr != nil {
		log.Fatalf("Errore nel push dell'immagine Docker: %v", pushErr)
	}
	fmt.Println("Push dell'immagine completato con successo.")

	fmt.Println("Operazione completata con successo!")
}

func getImageFullName(awsAccountID, region, imageName, tag string) string {
	return fmt.Sprintf("%s.dkr.ecr.%s.amazonaws.com/%s:%s", awsAccountID, region, imageName, tag)
}
