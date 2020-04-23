## SIMPLE CLIENT

### Description
This repo contains project that act as a client to microservices system.
This service is part of a big system. 
The whole system will be used to present technology show case.

### Features
- Serve movie data

This service serve data that is mentioned above through HTTP.

### How to run
#### Docker
- Install docker
- Create `config-dev.json` under `config` dir which contains following content
```json
{
  "api_gateway_address": "http://localhost:8081"
}
```
- Build and run docker image as below
```shell script
$ docker build -t simple-client .
$ docker run -p 8080:8080 simple-client
```
