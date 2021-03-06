# Go API client for kalaclient

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./kalaclient"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8000*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CreateJob**](docs/DefaultApi.md#createjob) | **Post** /api/v1/job/ | Creates a job
*DefaultApi* | [**DeleteJob**](docs/DefaultApi.md#deletejob) | **Delete** /api/v1/job/{jobId}/ | Deletes a job
*DefaultApi* | [**GetJob**](docs/DefaultApi.md#getjob) | **Get** /api/v1/job/{jobId}/ | Get job details
*DefaultApi* | [**Healthcheck**](docs/DefaultApi.md#healthcheck) | **Get** /api/v1/healthcheck | Check the avaialbility of the service


## Documentation For Models

 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponseDefault](docs/InlineResponseDefault.md)
 - [Job](docs/Job.md)
 - [JobWithId](docs/JobWithId.md)
 - [RemoteProperties](docs/RemoteProperties.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author



