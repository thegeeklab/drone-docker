# drone-docker

Drone plugin to build multiarch Docker images

[![Build Status](https://img.shields.io/drone/build/thegeeklab/drone-docker?logo=drone)](https://cloud.drone.io/thegeeklab/drone-docker)
[![Docker Hub](https://img.shields.io/badge/dockerhub-latest-blue.svg?logo=docker&logoColor=white)](https://hub.docker.com/r/thegeeklab/drone-docker)
[![Quay.io](https://img.shields.io/badge/quay-latest-blue.svg?logo=docker&logoColor=white)](https://quay.io/repository/thegeeklab/drone-docker)
[![Go Report Card](https://goreportcard.com/badge/github.com/thegeeklab/drone-docker)](https://goreportcard.com/report/github.com/thegeeklab/drone-docker)
[![GitHub contributors](https://img.shields.io/github/contributors/thegeeklab/drone-docker)](https://github.com/thegeeklab/drone-docker/graphs/contributors)
[![Source: GitHub](https://img.shields.io/badge/source-github-blue.svg?logo=github&logoColor=white)](https://github.com/thegeeklab/drone-docker)
[![License: MIT](https://img.shields.io/github/license/thegeeklab/drone-docker)](https://github.com/thegeeklab/drone-docker/blob/main/LICENSE)

Drone plugin to build multiarch Docker images.

## Docker Tags

Tags are following the main Docker version e.g. `19`, the second part is reflecting the plugin "version". A full example would be `19.0.5`.

## Build

Build the binary with the following command:

```Shell
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -tags netgo -o release/drone-docker
```

Build the Docker image with the following command:

```Shell
docker build --file docker/Dockerfile.amd64 --tag thegeeklab/drone-docker .
```

## Usage

> Notice: Be aware that the tis plugin requires privileged capabilities, otherwise the integrated Docker daemon is not able to start.

```console
docker run --rm \
  -e PLUGIN_TAG=latest \
  -e PLUGIN_REPO=octocat/hello-world \
  -e DRONE_COMMIT_SHA=00000000 \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  --privileged \
  thegeeklab/drone-docker --dry-run
```

## Contributors

Special thanks goes to all [contributors](https://github.com/thegeeklab/drone-docker/graphs/contributors). If you would like to contribute,
please see the [instructions](https://github.com/thegeeklab/drone-docker/blob/main/CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/thegeeklab/drone-docker/blob/main/LICENSE) file for details.
