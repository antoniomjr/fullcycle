
# Go Challenge for FullCycle Pós Go Expert

## Overview

This repository contains a Go application developed as a part of the "`Desafios/Client-Server-API`" (Go Challenge) from the Pós Go Expert.

## Challenge Requirements

- The Go application must to have a server and a client
- Server.go must extract the API containing the Dollar and Real exchange rate at the address: https://economia.awesomeapi.com.br/json/last/USD-BRL and then return the result to the client in JSON format .
- Using the "context" package, server.go must register each quote received in the SQLite database, with the maximum timeout to call the dollar quote API being 200ms and the maximum timeout to be able to persist the data in the database should be 10ms.
- Client.go will only need to receive the current exchange rate value from server.go ("bid" field in the JSON). Using the "context" package, client.go will have a maximum timeout of 300ms to receive the result from server.go.
- The 3 contexts should return an error in the logs if the execution time is insufficient.
- Client.go will have to save the current quote in a "quotacao.txt" file in the format: Dollar: {value}
- The necessary endpoint generated by server.go for this challenge will be: /cotacao and the port to be used by the HTTP server will be 8080.
  
## Project Structure

- `Server/server.go`: The Go source file that contains the application server side.
- `Client/client.go`: The Go source file that contains the application client side.
- `html-template/index.html`: The html source file that contains the template html and css style.
- `go.mod`: Specifies the Go module name and version (_This project does not have any external dependencies_).
- `Dockerfile.dev`: The Docker configuration file used to build the image for local development purposes.

## Instructions

To build and run this application, follow these steps:


#### Building the Docker Image

```bash
docker-compose up -d --remove-orphans 
``` 

###  Create the file to save the last dollar value
```bash
sqlite3 ./data/dolar_brl.db ".databases" 
```

### Script to create database
```bash
sqlite3 ./data/dolar_brl.db "CREATE TABLE dolar_brl (id VARCHAR(255) PRIMARY KEY,price VARCHAR(25), create_at VARCHAR(150);" 
```


#### Running the Go application
This will run the server side
```bash
go run Server/server.go
```
To redict to server side -> http://localhost:8080/cotacao

![Screenshot 2024-02-20 at 23 01 45](https://github.com/antoniomjr/fullcycle/assets/53837075/5cbf6a88-f229-4cf6-910e-f83c496175a5)


This will run the client side
```bash
go run Client/client.go
```
To redict to client side and call the server -> http://localhost:8090/cotacao?code=BRL

![Screenshot 2024-02-20 at 23 01 05](https://github.com/antoniomjr/fullcycle/assets/53837075/754af4b9-f072-4bf9-be8e-0f21edd7584b)

> [!WARNING]  
> Please note that will receive an error in case you call diffent endpoint.
> in case to check locali database -> docker exec -it sqlitebrowser /bin/bash                  



