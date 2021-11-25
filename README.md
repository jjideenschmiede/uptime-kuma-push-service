# Uptime Kuma Push Service

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/uptime-kuma-push-service.svg)](https://golang.org/) [![Go](https://github.com/jjideenschmiede/uptime-kuma-push-service/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/uptime-kuma-server-push/actions/workflows/go.yml) [![Docker Image CI](https://github.com/jjideenschmiede/uptime-kuma-push-service/actions/workflows/docker-image.yml/badge.svg)](https://github.com/jjideenschmiede/uptime-kuma-server-push/actions/workflows/docker-image.yml) [![Docker Hub](https://img.shields.io/docker/pulls/jjdevelopment/uptime-kuma-push-service.svg)](https://hub.docker.com/r/jjdevelopment/uptime-kuma-push-service)

This Docker image is for sending a heartbeat to an [Uptime Kuma](https://github.com/louislam/uptime-kuma) server. Here you will find a little introduction on how to use it.

The whole thing is brought to run in a Docker container. For this, a few variables from the Dockerfile are needed, if they are not stored, then they have a default value stored. You can find the variables here.

| Variable | Default value |
|----------|:-------------:|
| URL      | default       |
| MSG      | OK            |
| URL      | * * * * *     |

The URL must be stored. The other values can be used directly as default values. The milliseconds for the ping are calculated directly during the execution of the software.

## Launch Docker Container

To start the container properly, here is a small template. No volumes need to be mapped and no port is needed, because the software does not need to be accessed directly.

```console
docker run -d --restart always --name uptime-kuma-push-service -e URL='https://uptime-kuma.test.de/api/push/M4KzP0tSTB' jjdevelopment/uptime-kuma-push-service
```

Now the container can be started, so that the service can access your Uptime Kuma service and you always know if your servers in the office or at home are still running.

Click [here](https://hub.docker.com/r/jjdevelopment/uptime-kuma-push-service) to go directly to the Docker HUB.

## Contribute

If you want to help with development, or have found a bug, open a [new issue](https://github.com/jjideenschmiede/uptime-kuma-push-service/issues).
