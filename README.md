# MoneyTable Assignment

- [MoneyTable Assignment](#moneytable-assignment)
  - [Author : Chanchai Lee](#author--chanchai-lee)
  - [Description](#description)
  - [Prerequisites](#prerequisites)
  - [To Run Servers](#to-run-servers)
  - [Call APIs](#call-apis)
    - [sum](#sum)
    - [mul](#mul)
    - [div](#div)
    - [sub](#sub)

## Author : Chanchai Lee

## Description

Create 2 Go servers including the Proxy Server and the Calculate Server.

## Prerequisites

- Docker installed
- Docker Compose installed

## To Run Servers

```sh
$ cd assignment && docker-compose up

or

$ cd assignment && docker-compose up -d
```

## Call APIs

### sum

```sh
curl --request POST \
  --url http://localhost:8080/calculator.sum \
  --header 'content-type: application/json' \
  --data '{"a": 1110,"b": 200.99}'
```

### mul

```sh
curl --request POST \
  --url http://localhost:8080/calculator.mul \
  --header 'content-type: application/json' \
  --data '{"a": 1110,"b": 200.99}'
```

### div

```sh
curl --request POST \
  --url http://localhost:8080/calculator.div \
  --header 'content-type: application/json' \
  --data '{"a": 1110,"b": 200.99}'
```

### sub

```sh
curl --request POST \
  --url http://localhost:8080/calculator.sub \
  --header 'content-type: application/json' \
  --data '{"a": 1110,"b": 200.99}'
```
