# OwnAPIGateway
<!-- Intro  -->
<h3 align="center">
        <samp>Hey, I am
                <b><a target="_blank" href="http://www.diegorojas.tech">Diego R.</a></b>
        </samp>
</h3>

## API Gateway's stack
[![KrakenD](https://img.shields.io/badge/KrakenD-0f265c?style=plastic)](https://www.krakend.io/docs/overview/)
[![Golang](https://img.shields.io/badge/golang-00ADD8?logo=Go&logoColor=white&style=plastic)](https://go.dev)
[![Docker](https://img.shields.io/badge/docker-2496ED?logo=docker&logoColor=white&style=plastic)](https://hub.docker.com/layers/devopsfaith/krakend/2.6.2/images/sha256-4dd3cff749206e3f27eb19bc7379e7b146adde2ade56f1f2e70cd4a6a27efbf8?context=explore)
![Github](https://img.shields.io/badge/github-181717?logo=github&logoColor=white&style=plastic)
![VSCode](https://img.shields.io/badge/Visual_Studio-0078d7?logo=visual%20studio&logoColor=white&style=plastic)

## Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Services](#services)
- [Aditional Information](#aditional-information)
- [License](#license)



## Overview

OwnAPIGateway is a project developed to create an API Gateway solution using the Krakend framework in Golang. It serves as a template for future projects and includes three types of services:
- Single API to Single API
- Single API to Multi-API
- Single REST API to Single SOAP API using a customized plugin.

## Prerequisites

- Golang version: 1.22.2
- Krakend version: 2.6.2
- Docker

## Getting Started

1. Clone the repository.
2. Run the following command:
docker compose up --build

3. All services can be accessed through port 8080.

## Managed API Endpoints
| Service                                  | Method | Type of Service | Allowed Query Params        | Backend(s)                                                                                | Backend Method(s) | Backend Type of Service(s) | Backend Body(s)                                                                                                                                                                                                                                                      |
|------------------------------------------|--------|-----------------|-----------------------------|-------------------------------------------------------------------------------------------|-------------------|----------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| /Gateway/Drojas9811/test                 | GET    | REST            | None                        | https://httpbin.org/get                                                                   | GET               | REST                       | None                                                                                                                                                                                                                                                                 |
| /Gateway/Drojas9811/GameofThrones/houses | GET    | REST            | ["name", "region", "words"] | https://www.anapioficeandfire.com/api/houses                                              | GET               | REST                       | None                                                                                                                                                                                                                                                                 |
| /Gateway/Drojas9811/GameofThrones/books  | GET    | REST            | ["country", "name"]         | https://www.anapioficeandfire.com/api/books                                               | GET               | REST                       | None                                                                                                                                                                                                                                                                 |
| /Gateway/Drojas9811/GameofThrones/       | GET    | REST            | ["region", "words"]         | https://www.anapioficeandfire.com/api/books, https://www.anapioficeandfire.com/api/houses | GET, GET          | REST, REST                 | None, None                                                                                                                                                                                                                                                           |
| /Gateway/Drojas9811/number2words         | GET    | REST            | ["number2words"]            | https://www.dataaccess.com/webservicesserver/NumberConversion.wso                         | POST              | SOAP                       | `<?xml version="1.0" encoding="utf-8"?> <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"> <soap:Body> <NumberToWords xmlns="http://www.dataaccess.com/webservicesserver/"> <ubiNum>500</ubiNum> </NumberToWords> </soap:Body> </soap:Envelope>` |


## Aditional Information
- Developer: [DRojas9811](https://github.com/drojas9811)
- Website: [www.diegorojas.tech](https://www.diegorojas.tech)
- Krakend Documentation: [https://www.krakend.io/docs/](https://www.krakend.io/docs/)

## License

[![CC0](https://licensebuttons.net/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)

This project is released to the public domain. You are free to use, modify, and distribute the software for any purpose without any restrictions.
