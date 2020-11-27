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

Launch the install command
```bash
make install
```

Fill in the ``.env`` file at the project root with your parameters

## Run the application in development mode

Start the app
```bash
go run main.go
```

Then, open your browser and navigate to [http://localhost:SERVER_PORT](http://localhost:SERVER_PORT).
The ``SERVER_PORT`` variable is the one you defined in the .env file.

## Run the application in production mode

Build the app
```bash
go build
```

Start the app
```bash
./go-file-explorer
```

Then, open your browser and navigate to [http://localhost:SERVER_PORT](http://localhost:SERVER_PORT).
The ``SERVER_PORT`` variable is the one you defined in the .env file.

## Set the application as a background service

Copy the service file to your systemd
```bash
sudo cp go-file-explorer.service /etc/systemd/system/
```

Run the following command to start the service
```bash
sudo systemctl start myscript.service
```

Run the following command to auto start the service on boot
```bash
sudo systemctl enable myscript.service
```
