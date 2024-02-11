# sample-go-api

 Sample go lang api app 

## Pre-requisites

### GO

Install go version 1.20. [Download link.](https://go.dev/doc/install)

Verify go installation by running

```(bash)
go version
```

### OpenAPI generator

- Install the [openapi-generator](https://openapi-generator.tech/docs/installation/#homebrew)

Verify go installation by running

```(bash)
openapi-generator version
```

## Build the API using the OpenAPI spec

The [api.yaml](/api.yaml) file holds the OpenAPI spec for this repo.

### API server side code

- Update the API spec and generate the go server side code using the [openapi-generator](https://openapi-generator.tech/docs/installation/#homebrew)

- This repo is using openapi-generator to generate go-gin-server code

```(bash)
openapi-generator generate -i sample_api.yaml -g go-gin-server -o ./api --git-user-id shivinkapur --git-repo-id sample-go-api --package-name api --api-package api
```

## Build the app

Build the app locally using the following command on your root folder:

```(bash)
./build.sh
```

## Test the app

Run the following command on your root folder:

```(bash)
./test.sh
```

### Run the API module

To run the API module, you can use the following command from your root folder:

```(bash)
go run api/main.go
```

or if you have already run the build script, you can simply

```(bash)
./target/api
```

### Run the app

To run the go modules locally;

Create a file local_env.sh under project_root>cloud>config folder with the following content:

```(bash)
#!/bin/bash

if [[ -z "${SAMPLE_USER_DATA}" ]]; then
    export SAMPLE_USER_DATA="/path-to-repo//config/user_data.json"
fi

if [[ -z "${GIT_COMMIT}" ]]; then
    export GIT_COMMIT=$(git rev-parse HEAD)
fi

if [[ -z "${TAG_DATE}" ]]; then
    export TAG_DATE=$(date)
fi

if [[ -z "${BRANCH_NAME}" ]]; then
    export BRANCH_NAME=$(git branch --show-current)
fi

if [[ -z "${GIT_BRANCH}" ]]; then
    export GIT_BRANCH="$BRANCH_NAME@$GIT_COMMIT" 
fi

```


### Using VSCode

If you are using VSCode for this project, you can create the following launch file and local.env file under the `.vscode` folder

launch.json

```(json)
{
  // Use IntelliSense to learn about possible attributes.
  // Hover to view descriptions of existing attributes.
  // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Run API",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/api",
      "envFile": "${workspaceFolder}/.vscode/local.env"
    },
  ]
}
```

local.env

```(bash)
SAMPLE_USER_DATA="/path-to-repo/sample-go-api/config/user_data.json"
```

Note: Make sure to change the references to your relative path for the project
