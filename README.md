# README

## Overview
This document provides detailed information on the configuration and operation of the Docker image building and pushing system. The system is designed to automate Docker operations based on environment-specific settings defined in a `config-ebp.json` file located in the same directory as the script.

## Configuration File (`config-ebp.json`)
The configuration file is a JSON document that specifies various settings for different environments such as development (dev), user acceptance testing (uat), etc. The file should be named `config-ebp.json` and located in the same directory from which the script is executed.

### Structure of `config-ebp.json`
The `config-ebp.json` contains an array of environment objects, each defining settings such as AWS region, account ID, Docker image details, and Docker build arguments. Below is the structure of the file with details for each field:

```json
{
    "env": [
        {
            "name": "dev",
            "aws_region": "us-central-1",
            "aws_account_id": "12321",
            "aws_profile": "profile-dev",
            "docker_arg_vars": {
                "ARG1": "VALUE",
                "ARG2": "VALUE2"
            }
        },
        {
            "name": "uat",
            "aws_region": "us-east-1",
            "aws_account_id": "123456789012",
            "aws_profile": "default",
            "docker_image_name": "my-docker-image",
            "docker_image_tag": "latest",
            "docker_file_path": "/home/me/Dockerfile"
        }
    ]
}
```

### Configuration Fields

| Field              | Description                                                       |
|--------------------|-------------------------------------------------------------------|
| `name`             | The environment name (e.g., dev, uat).                            |
| `aws_region`       | AWS region where resources will be managed.                       |
| `aws_account_id`   | Your AWS account ID.                                              |
| `aws_profile`      | The AWS CLI profile to use.                                       |
| `docker_image_name`| Name of the Docker image to be built or pushed.                   |
| `docker_image_tag` | Tag for the Docker image (e.g., latest, v1.0).                    |
| `docker_file_path` | Path to the Dockerfile if not in the current directory.           |
| `docker_arg_vars`  | Optional. Map of build-time variables for the Docker build command.|


## Script Operation

### Behavior
The script performs the following operations based on the environment configuration specified:

1. **Load Configuration**: Tries to read the `config-ebp.json` file. If it does not exist or an environment is not specified, it will prompt the user to enter configuration details interactively.

2. **Docker Build**: Builds a Docker image using the details specified in the environment configuration. Includes the use of AWS credentials and Docker build arguments if provided.

3. **Docker Push**: Pushes the built image to a Docker registry, typically AWS ECR, using the account and region details specified in the configuration.

### Usage
To use this script, execute it from the directory containing `config-ebp.json`. Specify the environment as an argument:
```bash
./your_script.sh dev
```

### FAQs
**Q: What happens if the config-ebp.json file is not found?**

A: The script will prompt you to input all required details for the Docker build and push operations.

**Q: How do I add a new environment configuration?** 

A: Add a new object to the env array in the config-ebp.json file with all required fields completed.