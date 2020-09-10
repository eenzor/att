# att (ANZ Technical Test)
![CICD](https://github.com/eenzor/att/workflows/CICD/badge.svg)

att is a simple webserver which responds to the /version endpoint and prints some simple metadata.  
It is is written in golang using only the standard libraries.  

## Building the WebServer Binary

To build the binary run the script `./bin/build.sh`  
This will extract the latest version number from the change-log file and the latest commit SHA from git and pass these to the build so the /version endpoint shows us the correct metadata.

## Running the WebServer Binary

First build the binary,

```
Usage of ./att:
  -address string
        The TCP address to listen on (default "127.0.0.1")
  -log string
        the log format to use [none|combined|json|kv] (default "kv")
  -port int
        The TCP port to listen on (default 8000)
```

Additionally, the description can be changed by setting the `DESCRIPTION` environment variable.

## Building the WebServer Container

A dockerfile has been included which will build the binary into a docker image.
The script `./bin/docker-build.sh` can be used to build the image with the latest  
version tag from the change-log file.

## Running the WebServer Container

First build the container
```
docker run -p 8000:8000 att:latest
```

## Pipeline

A pipeline has been created using Github Actions.  
This is configured in `.github/workflows/main.yml`  
This pipeline will:  
- run the gosec security scanner
- run go vet
- run the golangci-lint linter
- unit test the code
- build the binary
- tag and publish the release to github
- build and publish the docker image to docker hub

## Versioning

The application is versioned using the CHANGELOG.md file following semantic versioning.  
The version.sh script extracts the latest version from the CHANGELOG file,  
and is used in the build.sh, docker-build.sh and pipelines to inject the version number  
into the binary and to tag the releases and docker images.  

## Testing

The `./bin/test.sh` script runs the same tests as the pipeline but locally.  
This makes it easy to check the code and the tests will pass before pushing.  

## Issues and Limitations

 - No TLS
 - No healthz endpoint
 - No (prometheus style) metrics endpoint
 - No load testing has been performedon the server
 - Only successful requests are logged
 - The server only listens on IPv4 addresses
 - The scripts assume the dependencies are installed (yamllint, go, docker)
 - The pipeline could be split into different jobs to improve performance
 - The pipeline fails when publishing a release if it already exists
 - The pipeline does not perform any integration or end-to-end tests
 - Unit test coverage is low (only happy paths are tested)
