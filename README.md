"# golang-microservices-gRPC" 
## Required

* [Go](http://www.dropwizard.io/1.0.2/docs/) - above version 1.11.4
* [dep](https://maven.apache.org/) - Dependency Management tool for Golang 
* [protoc](https://developers.google.com/protocol-buffers/) - Protocol Buffers (a.k.a., protobuf) are Google's language-neutral,
 platform-neutral, extensible mechanism for serializing structured data
* [docker](https://www.docker.com/) - version 18.09.2
* [GO-SWAGGER](https://github.com/go-swagger/go-swagger) - simple yet powerful representation of your RESTful API.
## USAGE

 1. Run simple client and server communicate between gRPC(simple):
    
        1.1 . execute go run cmd\gRPC\server\main.go
        
        1.2 . execute go run cmd\gRPC\client\main.go
 
 2 . Microservices patterns with golang communicate between gRPC: 
 
    //TODO
## Deployment
    //TODO 

## Using protocol buffers with Go
    protoc.exe -I $env:GOPATH\src --go_out=$env:GOPATH\src $env:GOPATH\src\.*.proto
       
    protoc.exe -I $env:GOPATH\src --go_out=plugins=gRPC:$env:GOPATH\src $env:GOPATH\src\.*.proto
## Description
This project is for study purpose (microservices patterns communicate between gRPC in golang)

## License
Pratice follwing this  [article](https://bitbucket.org/blog/writing-a-microservice-in-golang-which-communicates-over-grpc)
