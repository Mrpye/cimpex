# Cimpex - Container Import Export Utility
Export and import docker images from a registry without docker

# Description
The solution enables you to easily export and import Docker containers from a registry without the need to install docker.
It also can be run as a web-server allowing you to perform the Import Export via REST API

# When to use Cimpex
Cimpex can be used when you want to automate the import and export of docker images from a registry if you need to transport images between sites that don't have internet access. Also can be used as part of your CI/CD pipeline.

## Requirements
* A Container registry to import or export from.
* you will need to install go 1.8 [https://go.dev/doc/install](https://go.dev/doc/install) to run and install Cimpex

## Installation

```yaml
go install github.com/Mrpye/cimpex
```

## Run as a container
1. Clone the repository

```
git clone https://github.com/Mrpye/cimpex.git
```

2. Build the container as API endpoint
```
sudo docker build . -t  cimpex:v1.0.0 -f Dockerfile
```
3. Run the container
```
sudo docker run -d -p 9020:8080 --name=cimpex --restart always  -v /host_path/images:/go/bin/images --env=BASE_FOLDER=/go/bin/images --env=WEB_IP=0.0.0.0  --env=WEB_PORT=8080 -t cimpex:1.0.0
```

### Environment  variables
- BASE_FOLDER (set where the images will be exported)
- WEB_IP (set the listening ip address 0.0.0.0 allow from everywhere)
- WEB_PORT (set the port to listen on)

## How to use Cimpex CLI
Check out the CLI documentation [here](./documents/cimpex.md)


# Using the API
## Run web server
```bash
    cimpex web -f /go/bin/images -p 8080 -i 0.0.0.0
```
## Export image from a registry
``` bash
curl --location --request POST 'http://localhost:9020/export' \
--header 'Content-Type: application/json' \
--data-raw '{
    "target":"cimpex:v1.0.0",
    "tar":"cimpex-v1-0-0.tar",
    "ignore_ssl":true
}
```

Export from a registry with authentication
```bash
curl --location --request POST 'http://172.16.10.237:9020/export' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "target":"172.19.2.152/library/cimpex:v1.0.0",
    "tar":"cimpex-v1-0-0.tar",
    "ignore_ssl":true
}
```

## Import image from a registry
```bash
curl --location --request POST 'localhost:8080/import' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "target":"library/bind:test",
    "tar":"bind-latest.tar",
    "ignore_ssl":true
}'
```

## Json Payload

- target (location of the docker image import/export)
- tar (name of the tar file will be saved in the export folder)
- ignore_ssl (Ignore ssl cert)

# To Do
- want to be able to get the exported container details from the tar image name and tag so you can just point to the repo and the name and tag will be populated.
- improve validation and error handling

# 3rd party Libraries
https://github.com/google/go-containerregistry

# license
cimplex is Apache 2.0 licensed.