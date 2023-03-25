# Cimpex - Container Import Export Utility
Export and import docker images from a registry without docker

## Description
The solution enables you to easily export and import Docker containers from a registry without the need to install docker.
It also can be run as a web-server allowing you to perform the Import Export via REST API

## When to use Cimpex
Cimpex can be used when you want to automate the import and export of docker images from a registry if you need to transport images between sites that don't have internet access. Also can be used as part of your CI/CD pipeline.

---
## Requirements
* A Container registry to import or export from.
* you will need to install go 1.8 [https://go.dev/doc/install](https://go.dev/doc/install) to run and install Cimpex

---

## Project folders
Below is a description cimpex project folders and what they contain
|   Folder        | Description  | 
|-----------|---|
| charts    | Contains the helm chart for cimpex  |
| docs      | Contains the swagger documents |
| documents | Contains cli and api markdown files  |
| modules   | Contains cimpex modules and code  |
| config    | Contains Example payload config files  |
| cmd       | Contains code for cimpex CLI   |
|           |   |

---

## Installation

```yaml
go install github.com/Mrpye/cimpex
```

## Run as a container

<details>
<summary>1. Clone the repository</summary>

```
git clone https://github.com/Mrpye/cimpex.git
```
</details>

<details>
<summary>2. Build the container as API endpoint</summary>

```
sudo docker build . -t  cimpex:v1.0.0 -f Dockerfile
```

</details>

<details>
<summary>3. Run the container</summary>

```
sudo docker run -d -p 9020:8080 --name=cimpex --restart always  -v /host_path/images:/go/bin/images --env=BASE_FOLDER=/go/bin/images --env=WEB_IP=0.0.0.0  --env=WEB_PORT=8080 -t cimpex:1.0.0
```

</details>

---

### Environment  variables
- BASE_FOLDER (set where the images will be exported)
- WEB_IP (set the listening ip address 0.0.0.0 allow from everywhere)
- WEB_PORT (set the port to listen on)


---
## How to use Cimpex CLI
Check out the CLI documentation [here](./documents/cimpex.md)

---

# Using the API

## Run web server
```bash
    cimpex web -f /go/bin/images -p 8080 -i 0.0.0.0
```
---

## Examples

<details>
<summary>Export image from a registry</summary>

Export and image fromm a registry with no authentication

``` bash
curl --location --request POST 'http://localhost:8080/export' \
--header 'Content-Type: application/json' \
--data-raw '{
    "target":"cimpex:v1.0.0",
    "tar":"cimpex-v1-0-0.tar",
    "ignore_ssl":true
}
```

Export from a registry with authentication

```bash
curl --location --request POST 'http://localhost:8080/export' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "target":"127.0.0.1/library/cimpex:v1.0.0",
    "tar":"cimpex-v1-0-0.tar",
    "ignore_ssl":true
}
```
</details>

<details>
<summary>Export multiple images from a registry</summary>

## Export images from a registry
``` bash
curl --location --request POST 'http://localhost:8080/exports' \
--header 'Content-Type: application/json' \
--data-raw '[{
    "target":"cimpex:v1.0.0",
    "tar":"cimpex-v1-0-0.tar",
    "ignore_ssl":true
},
{
    "target":"helm-api:v1.0.0",
    "tar":"helm-api-v1-0-0.tar",
    "ignore_ssl":true
}
]
```
</details>


<details>
<summary>Import image to a registry</summary>

Import an image and specify the name at tag to use

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

Import an image and use the name at tag from the tar file

```bash
curl --location --request POST 'localhost:8080/import' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "target":"library/",
    "tar":"bind-latest.tar",
    "ignore_ssl":true
}'
```

## Json Payload

- target (location of the docker image import/export)
- tar (name of the tar file will be saved in the export folder)
- ignore_ssl (Ignore ssl cert)
</details>

<details>
<summary>Import images to a registry from the BASE_FOLDER
 folder</summary>

Import an images and imports to the target registry

```bash
curl --location 'localhost:8080/imports' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--header 'Content-Type: application/json' \
--data '{
    "target":"127.0.0.1/library/",
    "ignore_ssl":true
}'

```

## Json Payload

- target (the docker registry)
- ignore_ssl (Ignore ssl cert)
</details>


<details>
<summary>List images in the folder</summary>

```bash
curl --location --request POST ''http://localhost:8080/list' \
--header 'Content-Type: application/json' 
```
</details>


<details>
<summary>Test is alive</summary>

```bash
curl --location --request GET ''http://localhost:8080/'
```
Return OK
</details>

---
## cimpex Helm chart
This guide will show you how to build the helm chart package for cimpex, you will need to have helm installed to build the package.

<details>
<summary>1. Build the helm chart package for cimpex</summary>

```bash
# change into the chart directory
cd charts
# Package the cimpex chart
helm package cimpex

```

the helm chart package will be saved under the charts folder cimpex-0.1.0.tgz

</details>

---

## Update the swagger document
The code below shows you how to update the swagger API documents.

If you need more helm on using these tools please refer to the links below
- gin-swagge [https://github.com/swaggo/gin-swagger](https://github.com/swaggo/gin-swagger)
- swag [https://github.com/swaggo/swag](https://github.com/swaggo/swag)

<details>
<summary>1. Install swag</summary>

```bash
#Install swag
go install github.com/swaggo/swag/cmd/swag
```
</details>

<details>
<summary>2. Update APi document</summary>

```bash
#update the API document
swag init
```
</details>
<details>
<summary>3. Update the api.md</summary>

```bash
swagger generate markdown -f .\docs\swagger.json --output .\documents\api.md 
```
</details>

---

# To Do
- improve validation and error handling

---

# 3rd party Libraries
https://github.com/google/go-containerregistry

---
# license
cimplex is Apache 2.0 licensed.