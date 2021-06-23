package model

const DockerComposeLocalFile = `version: '3.7'

services:
  app:
    container_name: ai-unit-app
    build:
      context: services/app
      dockerfile: local.dockerfile
    volumes:
      - './services/app:/usr/src/app'
    entrypoint: go run main.go
    ports:
      - 9404:9404
`

const DockerComposeReleaseFile = `version: '3.7'

services:
  app:
    container_name: ginhelper
    image: ginhelper:v0.1.0
    restart: always
    volumes:
      - './logs:/logs'
      - './config.yaml:/config.yaml'
    entrypoint: /app
    ports:
      - 9404:9404
`
