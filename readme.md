# Validate Sudoku with Golang

## Requirement
You need to install this on your local :
- Golang, if you want to run manualy
- Docker, if you want to run without golang

## Installation
- Open your command prompt and direct to your `GOPATH` directory.
- Run this command: `go get -u -d github.com/yuhando/validatesudoku` to download it to your local gopath directory.

## Run with Golang
- Go to `validatesudoku` directory.
- Run this command: `go build`
- Run this command: `./foldername-of-this-apps` (the default is validatesudoku) 
- Open your browser and Run this : `http://localhost:8081/`

## Run with Docker
- Go to `validatesudoku` directory.
- Run this command: `docker build -t validatesudoku .`
- Run this command: `docker image ls` to check the docker image you just made.
- Run this command: `docker run -p 8081:8081 validatesudoku` to run the app.
- Open your browser and Run this : `http://localhost:8081/`

## Thank You.