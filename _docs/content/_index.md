---
title: drone-docker
---

[![Build Status](https://img.shields.io/drone/build/thegeeklab/drone-docker?logo=drone&server=https%3A%2F%2Fdrone.thegeeklab.de)](https://drone.thegeeklab.de/thegeeklab/drone-docker)
[![Docker Hub](https://img.shields.io/badge/dockerhub-latest-blue.svg?logo=docker&logoColor=white)](https://hub.docker.com/r/thegeeklab/drone-docker)
[![Quay.io](https://img.shields.io/badge/quay-latest-blue.svg?logo=docker&logoColor=white)](https://quay.io/repository/thegeeklab/drone-docker)
[![GitHub contributors](https://img.shields.io/github/contributors/thegeeklab/drone-docker)](https://github.com/thegeeklab/drone-docker/graphs/contributors)
[![Source: GitHub](https://img.shields.io/badge/source-github-blue.svg?logo=github&logoColor=white)](https://github.com/thegeeklab/drone-docker)
[![License: MIT](https://img.shields.io/github/license/thegeeklab/drone-docker)](https://github.com/thegeeklab/drone-docker/blob/main/LICENSE)

Drone plugin to build and publish multiarch Docker images.

<!-- prettier-ignore-start -->
<!-- spellchecker-disable -->
{{< toc >}}
<!-- spellchecker-enable -->
<!-- prettier-ignore-end -->

## Versioning

The tags follow the major version of Docker, e.g. `19`, the minor and patch part reflects the "version" of the plugin. A full example would be `19.6.5`. Minor versions may introduce breaking changes, while patch versions may be considered non-breaking.

## Usage

```YAML
kind: pipeline
name: default

steps:
  - name: docker
    image: thegeeklab/drone-docker
    settings:
      username: octocat
      password: secure
      repo: octocat/example
      tags: latest
```

### Parameters

<!-- prettier-ignore-start -->
<!-- spellchecker-disable -->
{{< propertylist name=drone-docker.data >}}
<!-- spellchecker-enable -->
<!-- prettier-ignore-end -->

## Build

Build the binary with the following command:

```Shell
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

make build
```

Build the Docker image with the following command:

```Shell
docker build --file docker/Dockerfile.amd64 --tag thegeeklab/drone-docker .
```

## Test

{{< hint type=important >}}
Be aware that the this plugin requires privileged capabilities, otherwise the integrated Docker daemon is not able to start.
{{< /hint >}}

```Shell
docker run --rm \
  -e PLUGIN_TAG=latest \
  -e PLUGIN_REPO=octocat/hello-world \
  -e DRONE_COMMIT_SHA=00000000 \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  --privileged \
  thegeeklab/drone-docker --dry-run
```
