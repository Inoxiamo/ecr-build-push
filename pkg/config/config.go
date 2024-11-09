// pkg/config/config.go
package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Environment struct {
	Name            string            `json:"name"`
	AWSRegion       string            `json:"aws_region"`
	AWSAccountID    string            `json:"aws_account_id"`
	AWSProfile      string            `json:"aws_profile"`
	PathContext     string            `json:"path_context"`
	DockerImageName string            `json:"docker_image_name"`
	DockerImageTag  string            `json:"docker_image_tag"`
	DockerFilePath  string            `json:"docker_file_path"`
	DockerEnvVars   map[string]string `json:"docker_env_vars"`
	DockerArgVars   map[string]string `json:"docker_arg_vars"`
}

type Config struct {
	Environments []Environment `json:"env"`
}

// LoadConfigForEnv carica la configurazione specifica per un ambiente dal file JSON.
func LoadConfigForEnv(filename, envName string) (*Environment, error) {
	var config Config
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("File di configurazione non trovato, procedura guidata di configurazione...")
		return promptForConfig(envName), nil
	}
	if err := json.Unmarshal(data, &config); err != nil {
		fmt.Println("Errore nel parsing del file di configurazione, procedura guidata di configurazione...")
		return promptForConfig(envName), nil
	}

	for _, env := range config.Environments {
		if env.Name == envName {
			return completeConfig(&env), nil
		}
	}

	fmt.Println("Ambiente non trovato nel file di configurazione, procedura guidata di configurazione...")
	return promptForConfig(envName), nil
}

func promptForConfig(envName string) *Environment {
	env := Environment{Name: envName}
	return completeConfig(&env)
}

func completeConfig(env *Environment) *Environment {
	reader := bufio.NewReader(os.Stdin)

	if env.AWSRegion == "" {
		fmt.Print("Inserire AWS Region: ")
		env.AWSRegion, _ = reader.ReadString('\n')
		env.AWSRegion = strings.TrimSpace(env.AWSRegion)
	}

	if env.AWSAccountID == "" {
		fmt.Print("Inserire AWS Account ID: ")
		env.AWSAccountID, _ = reader.ReadString('\n')
		env.AWSAccountID = strings.TrimSpace(env.AWSAccountID)
	}

	if env.AWSProfile == "" {
		fmt.Print("Inserire AWS Profile: ")
		env.AWSProfile, _ = reader.ReadString('\n')
		env.AWSProfile = strings.TrimSpace(env.AWSProfile)
	}

	if env.DockerImageName == "" {
		fmt.Print("Inserire Docker Image Name: ")
		env.DockerImageName, _ = reader.ReadString('\n')
		env.DockerImageName = strings.TrimSpace(env.DockerImageName)
	}

	if env.DockerImageTag == "" {
		fmt.Print("Inserire Docker Image Tag: ")
		env.DockerImageTag, _ = reader.ReadString('\n')
		env.DockerImageTag = strings.TrimSpace(env.DockerImageTag)
	}

	if env.DockerFilePath == "" {
		fmt.Print("Inserire Docker File Path: ")
		env.DockerFilePath, _ = reader.ReadString('\n')
		env.DockerFilePath = strings.TrimSpace(env.DockerFilePath)
	}

	if env.PathContext == "" {
		fmt.Print("Inserire Path Context: ")
		env.PathContext, _ = reader.ReadString('\n')
		env.PathContext = strings.TrimSpace(env.PathContext)
	}

	// Qui potresti aggiungere logica simile per DockerEnvVars e DockerArgVars se necessario

	return env

}
