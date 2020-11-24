# Go File Explorer

## Introduction

This tool provides you a very lightweight interface to navigate through a file system.
It was developed to provide a UI to a NAS project running on a Raspberry Pi.

## Prerequisite

- [Golang](https://golang.org/dl/) with a working workspace (GOPATH, GOROOT...)

## Installation

Clone the repository in your Gopath (usually $GOPATH/src/)
```bash
git clone git@github.com:cyprille/go-file-explorer.git $GOPATH/src/go-file-explorer
```

Enter the project
```bash
cd $GOPATH/src/go-file-explorer
```

Import and update dependencies
```bash
go get -u ./...
```

## Run the application

Start the app
```bash
go run main.go
```

Then, open your browser and navigate to [http://localhost:8080](http://localhost:8080)
