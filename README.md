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
