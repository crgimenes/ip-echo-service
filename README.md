# IP echo service

This is a simple web service that returns the IP address of the client accessing it. The service supports multiple output formats (HTML, JSON, and plain text)

## Installation

### Build from Source

#### Prerequisites

- [Go](https://golang.org/dl/) 1.23 or later

```bash
go install github.com/crgimenes/ip-echo-service@latest
```

Or you can clone the repository and build it:

```bash
git clone git@github.com:crgimenes/ip-echo-service.git
cd ip-echo-service
go build -o ip-echo-service
```

Alternatively, you can use go run to run the service without building:

```bash
go run main.go
```


