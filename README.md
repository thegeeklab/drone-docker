# drone-docker

Drone plugin to build multiarch Docker images

[![Build Status](https://img.shields.io/drone/build/thegeeklab/drone-docker?logo=drone&server=https%3A%2F%2Fdrone.thegeeklab.de)](https://drone.thegeeklab.de/thegeeklab/drone-docker)
[![Docker Hub](https://img.shields.io/badge/dockerhub-latest-blue.svg?logo=docker&logoColor=white)](https://hub.docker.com/r/thegeeklab/drone-docker)
[![Quay.io](https://img.shields.io/badge/quay-latest-blue.svg?logo=docker&logoColor=white)](https://quay.io/repository/thegeeklab/drone-docker)
[![Go Report Card](https://goreportcard.com/badge/github.com/thegeeklab/drone-docker)](https://goreportcard.com/report/github.com/thegeeklab/drone-docker)
[![GitHub contributors](https://img.shields.io/github/contributors/thegeeklab/drone-docker)](https://github.com/thegeeklab/drone-docker/graphs/contributors)
[![Source: GitHub](https://img.shields.io/badge/source-github-blue.svg?logo=github&logoColor=white)](https://github.com/thegeeklab/drone-docker)
[![License: MIT](https://img.shields.io/github/license/thegeeklab/drone-docker)](https://github.com/thegeeklab/drone-docker/blob/main/LICENSE)

Drone plugin to build multiarch Docker images. This plugin is a fork of [drone-plugins/drone-docker](https://github.com/drone-plugins/drone-docker).

## Docker Tags

Tags are following the main Docker version e.g. `19.03`, the second part is reflecting the plugin "version". A full example would be `19.03.5`.

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

## Parameters

dry_run
: disables docker push

drone_remote_url
: sets the git remote url

mirror
: sets a registry mirror to pull images

storage_driver
: sets the docker daemon storage driver

storage_path (default `/var/lib/docker`)
: sets the docker daemon storage path

bip
: allows the docker daemon to bride ip address

mtu
: sets docker daemon custom mtu setting

custom_dns
: sets custom docker daemon dns server

custom_dns_search
: sets custom docker daemon dns search domain

insecure
: allows the docker daemon to use insecure registries

ipv6
: enables docker daemon ipv6 support

experimental
: enables docker daemon experimental mode

debug "docker_launch_debug
: enables verbose debug mode for the docker daemon

daemon_off
: disables the startup of the docker daemon

dockerfile (default `./Dockerfile`)
: sets dockerfile to use for the image build

context (default `./`)
: sets the path of the build context to use

tags (default `latest`)
: sets repository tags to use for the image; tags can also be loaded from a `.tags` file

auto_tag
: generates tag names automatically based on git branch and git tag

auto_tag_suffix
: generates tag names with the given suffix

build_args
: sets custom build arguments for the build

build_args_from_env
: forwards environment variables as custom arguments to the build

quiet
: enables suppression of the build output

target
: sets the build target to use

cache_from
: sets images to consider as cache sources

pull_image (default `true`)
: enforces to pull base image at build time

compress
: enables compression og the build context using gzip

repo
: sets repository name for the image

registry (default `https://index.docker.io/v1/`)
: sets docker registry to authenticate with

username
: sets username to authenticates with

password
: sets password to authenticates with

email
: sets email addres to authenticates with

config
: sets content of the docker daemon json config

purge (default `true`)
: enables cleanup of the docker environment at the end of a build

no_cache
: disables the usage of cached intermediate containers

add_host
: sets additional host:ip mapping

## Contributors

Special thanks goes to all [contributors](https://github.com/thegeeklab/drone-docker/graphs/contributors). If you would like to contribute,
please see the [instructions](https://github.com/thegeeklab/drone-docker/blob/main/CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/thegeeklab/drone-docker/blob/main/LICENSE) file for details.
