[![GuardRails badge](https://api.guardrails.io/v2/badges/ileossa/go-bank-backend.svg?token=47c732401be09e2f2af23f20cf9bf691a7555b8eb37fa16adf349bfd1e10818e&provider=github)](https://dashboard.guardrails.io/gh/ileossa/95105)

# Technical test - bank backend

## Requirements:

* docker >= 17.12.0+
* docker-compose
* golang >= 1.17
* make

# Quick Start

use `make dev` to run app locally

# Docker

## Build image

`make package`

## Publish image

`make publish-images`

# Documentation

Use `make doc` to generate the documentation.

ð You can publish documentation with `make doc-publish` on the branch /publish/documentation ð

## Postgresql & PgAdmin powered by compose

### Quick Start

* Clone or download this repository
* Go inside of directory,  `cd compose-postgres`
* Run this command `docker-compose up -d`


### Environments
This Compose file contains the following environment variables:

* `POSTGRES_USER` the default value is **postgres**
* `POSTGRES_PASSWORD` the default value is **changeme**
* `PGADMIN_PORT` the default value is **5050**
* `PGADMIN_DEFAULT_EMAIL` the default value is **pgadmin4@pgadmin.org**
* `PGADMIN_DEFAULT_PASSWORD` the default value is **admin**

### Access to postgres:
* `localhost:5432`
* **Username:** postgres (as a default)
* **Password:** changeme (as a default)

### Access to PgAdmin:
* **URL:** `http://localhost:5050`
* **Username:** pgadmin4@pgadmin.org (as a default)
* **Password:** admin (as a default)

### Add a new server in PgAdmin:
* **Host name/address** `postgres`
* **Port** `5432`
* **Username** as `POSTGRES_USER`, by default: `postgres`
* **Password** as `POSTGRES_PASSWORD`, by default `changeme`



## Troubleshooting

# if port 8080 already used

`lsof -i :8080 | grep LISTEN`
`kill -9 <pid>`
