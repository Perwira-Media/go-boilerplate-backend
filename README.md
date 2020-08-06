# Perwira Media Go Boiler

A boiler codebase for creating API based on golang.

## Main Features

- Database Connection

## Status
Under Development

## How to Start Development

### Clone Repository
```bash
$ cd $GOPATH/src/
$ mkdir Perwira-Media
$ cd Perwira-Media
$ git clone https://github.com/Perwira-Media/go-boilerplate-backend.git
```

## Running Services

### Netmonk PM-Boiler-Code Service

Create boiler-config.yaml in _bin/conf
Tamplate config file :
```bash
service_data:
  address: localhost:8080
  
source_data:
  postgres_server: localhost
  postgres_dbname: test
  postgres_username: root
  postgres_password: pass
  postgres_port: 5432
  postgres_timeout: 5

```

Build Service
```bash
./_scripts/build.sh
```

Run Service
```bash
./_scripts/run.sh
```