


# cimpex
cimpex is a CLI application written in Golang that gives the ability import and export docker images from a repository. GitHub repository at https://github.com/Mrpye/compex
  

## Informations

### Version

1.0

### License

[Apache 2.0 licensed](https://github.com/Mrpye/cimpex/blob/main/LICENSE)

### Contact

  https://github.com/Mrpye/cimpex

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  operations

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | / | [check api endpoint](#check-api-endpoint) | Check API Endpoint |
| POST | /export | [post export docker image](#post-export-docker-image) | Export Docker Image from Registry to tar file |
| POST | /exports | [post exports docker images](#post-exports-docker-images) | Exports Docker Images from Registry to tar file |
| POST | /import | [post import docker image](#post-import-docker-image) | Import Docker Image to Registry from tar file |
| POST | /imports | [post import list docker images](#post-import-list-docker-images) | Import the tar files in the directory |
| POST | /list | [post list docker images](#post-list-docker-images) | List the docker images tar files in the directory |
  


## Paths

### <span id="check-api-endpoint"></span> Check API Endpoint (*check-api-endpoint*)

```
GET /
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#check-api-endpoint-200) | OK | ok |  | [schema](#check-api-endpoint-200-schema) |

#### Responses


##### <span id="check-api-endpoint-200"></span> 200 - ok
Status: OK

###### <span id="check-api-endpoint-200-schema"></span> Schema
   
  



### <span id="post-export-docker-image"></span> Export Docker Image from Registry to tar file (*post-export-docker-image*)

```
POST /export
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| request | `body` | [BodyTypesImportExportRequest](#body-types-import-export-request) | `models.BodyTypesImportExportRequest` | | ✓ | | query params |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-export-docker-image-200) | OK | tar file Exported |  | [schema](#post-export-docker-image-200-schema) |
| [404](#post-export-docker-image-404) | Not Found | error |  | [schema](#post-export-docker-image-404-schema) |

#### Responses


##### <span id="post-export-docker-image-200"></span> 200 - tar file Exported
Status: OK

###### <span id="post-export-docker-image-200-schema"></span> Schema
   
  



##### <span id="post-export-docker-image-404"></span> 404 - error
Status: Not Found

###### <span id="post-export-docker-image-404-schema"></span> Schema
   
  



### <span id="post-exports-docker-images"></span> Exports Docker Images from Registry to tar file (*post-exports-docker-images*)

```
POST /exports
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| request | `body` | [][BodyTypesImportExportRequest](#body-types-import-export-request) | `[]*models.BodyTypesImportExportRequest` | | ✓ | | query params |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-exports-docker-images-200) | OK | tar file Exported |  | [schema](#post-exports-docker-images-200-schema) |
| [404](#post-exports-docker-images-404) | Not Found | error |  | [schema](#post-exports-docker-images-404-schema) |

#### Responses


##### <span id="post-exports-docker-images-200"></span> 200 - tar file Exported
Status: OK

###### <span id="post-exports-docker-images-200-schema"></span> Schema
   
  



##### <span id="post-exports-docker-images-404"></span> 404 - error
Status: Not Found

###### <span id="post-exports-docker-images-404-schema"></span> Schema
   
  



### <span id="post-import-docker-image"></span> Import Docker Image to Registry from tar file (*post-import-docker-image*)

```
POST /import
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| request | `body` | [BodyTypesImportExportRequest](#body-types-import-export-request) | `models.BodyTypesImportExportRequest` | | ✓ | | query params |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-import-docker-image-200) | OK | tar file imported |  | [schema](#post-import-docker-image-200-schema) |
| [404](#post-import-docker-image-404) | Not Found | error |  | [schema](#post-import-docker-image-404-schema) |

#### Responses


##### <span id="post-import-docker-image-200"></span> 200 - tar file imported
Status: OK

###### <span id="post-import-docker-image-200-schema"></span> Schema
   
  



##### <span id="post-import-docker-image-404"></span> 404 - error
Status: Not Found

###### <span id="post-import-docker-image-404-schema"></span> Schema
   
  



### <span id="post-import-list-docker-images"></span> Import the tar files in the directory (*post-import-list-docker-images*)

```
POST /imports
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| request | `body` | [BodyTypesImportExportRequest](#body-types-import-export-request) | `models.BodyTypesImportExportRequest` | | ✓ | | query params |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-import-list-docker-images-200) | OK | OK |  | [schema](#post-import-list-docker-images-200-schema) |
| [404](#post-import-list-docker-images-404) | Not Found | error |  | [schema](#post-import-list-docker-images-404-schema) |

#### Responses


##### <span id="post-import-list-docker-images-200"></span> 200 - OK
Status: OK

###### <span id="post-import-list-docker-images-200-schema"></span> Schema
   
  

[][BodyTypesPackageInfo](#body-types-package-info)

##### <span id="post-import-list-docker-images-404"></span> 404 - error
Status: Not Found

###### <span id="post-import-list-docker-images-404-schema"></span> Schema
   
  



### <span id="post-list-docker-images"></span> List the docker images tar files in the directory (*post-list-docker-images*)

```
POST /list
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-list-docker-images-200) | OK | OK |  | [schema](#post-list-docker-images-200-schema) |
| [404](#post-list-docker-images-404) | Not Found | error |  | [schema](#post-list-docker-images-404-schema) |

#### Responses


##### <span id="post-list-docker-images-200"></span> 200 - OK
Status: OK

###### <span id="post-list-docker-images-200-schema"></span> Schema
   
  

[][BodyTypesPackageInfo](#body-types-package-info)

##### <span id="post-list-docker-images-404"></span> 404 - error
Status: Not Found

###### <span id="post-list-docker-images-404-schema"></span> Schema
   
  



## Models

### <span id="body-types-import-export-request"></span> body_types.ImportExportRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ignore_ssl | boolean| `bool` |  | |  |  |
| tar | string| `string` |  | |  |  |
| target | string| `string` |  | |  |  |



### <span id="body-types-package-info"></span> body_types.PackageInfo


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| image_name_tag | string| `string` |  | |  |  |
| status | string| `string` |  | |  |  |
| tar_path | string| `string` |  | |  |  |


