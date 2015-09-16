# Minimal Service

Example of simple Service Interface, Request/Response, Endpoints and Transport using [Go kit](https://github.com/go-kit).

### Build this Example
```bash
$ go build -o minimal-service main.go
```

### Run this Example
```bash
$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
{"v":"HELLO WORLD"}
$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
{"v":12}
```
