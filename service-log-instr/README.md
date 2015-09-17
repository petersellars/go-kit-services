# Logging and Instrumentation Service

Example of Service with Logging and Instrumentation using [Go kit](https://github.com/go-kit).

### Build this Example
```bash
$ go build -o service-li main.go
$ ./service-li
```

### Run this Example
```bash
$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
{"v":"HELLO WORLD"}
$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
{"v":12}
```
