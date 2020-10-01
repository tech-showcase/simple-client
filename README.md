## SIMPLE CLIENT

### Description
This repo contains project that act as a **client**.
This client is part of a big system. 
The whole system will be used to present **micro-services without an orchestrator**.

### Features
- Serve movie data thru stdout by executing CLI

### How to run
#### Docker
- Install docker
- Create following environment variable and fill it with the right value
```shell script
  API_GATEWAY_ADDRESS=http://api-gateway-address
```
- Build and run docker image as below
```shell script
$ docker build -t simple-client .
$ docker run -p 8080:8080 simple-client
```

### Tech / Dependency
- [Go kit - service](https://github.com/go-kit/kit)
