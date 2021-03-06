# Go File Explorer

![status](https://github.com/cyprille/go-file-explorer/workflows/status/badge.svg)
[![Release](https://img.shields.io/github/release/cyprille/go-file-explorer.svg)](https://github.com/cyprille/go-file-explorer/releases/latest)
[![License](https://img.shields.io/github/license/cyprille/go-file-explorer)](/LICENSE)
[![Github Releases Stats of golangci-lint](https://img.shields.io/github/downloads/cyprille/go-file-explorer/total.svg?logo=github)](https://somsubhra.com/github-release-stats/?username=cyprille&repository=go-file-explorer)

## Introduction

This tool provides you a very lightweight interface to navigate through a file system.
It was developed to provide a UI to a NAS project running on a Raspberry Pi 3B+.

# Prerequisite

- [Golang](https://golang.org/dl/) with a working workspace (GOPATH, GOROOT...)

# Installation

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

# Run

## Run the application in development mode

Start the app
```bash
go run main.go
```

Then, open your browser and navigate to [http://localhost:SERVER_PORT](http://localhost:SERVER_PORT).
The ``SERVER_PORT`` variable is the one you defined in the ``.env`` file.

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
The ``SERVER_PORT`` variable is the one you defined in the ``.env`` file.

## Set the application as a background service

Copy the service file to your systemd
```bash
sudo cp ressources/go-file-explorer.service /etc/systemd/system/
```

Run the following command to start the service
```bash
sudo systemctl start go-file-explorer
```

Run the following command to auto start the service on boot
```bash
sudo systemctl enable go-file-explorer
```

# Deploy

## Automated update

If you want to automatically update the app from the main repository, I suggest the use of the extremely simple [Git-Auto-Deploy](https://github.com/olipo186/Git-Auto-Deploy).

# Code of conduct

[See CODE_OF_CONDUCT.md](https://github.com/cyprille/go-file-explorer/blob/master/CODE_OF_CONDUCT.md)

# Contributing

[See CONTRIBUTING.md](https://github.com/cyprille/go-file-explorer/blob/master/CONTRIBUTING.md)

# Licence

This project is licended under the MIT licence.

[See LICENCE.md](https://github.com/cyprille/go-file-explorer/blob/master/LICENSE.md)
