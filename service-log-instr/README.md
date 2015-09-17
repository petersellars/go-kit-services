# Logging and Instrumentation Service

Example of Service with Logging and Instrumentation using [Go kit](https://github.com/go-kit).

### Build this Example
```bash
$ go build . 
$ ./service-log-instr
```

### Run this Example
```bash
$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
{"v":"HELLO WORLD"}
$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
{"v":12}
```

```bash
msg=HTTP addr=:8080
method=uppercase msg="calling endpoint"
method=uppercase input="hello, world" output="HELLO, WORLD" err=null took=3.544µs
method=uppercase msg="called endpoint"
method=count msg="calling endpoint"
method=count input="hello, world" n=12 took=1.189µs
method=count msg="called endpoint"
```

### Metrics - Prometheus Example

```bash
$ curl localhost:8080/metrics
```

```bash
# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0.000183487
go_gc_duration_seconds{quantile="0.25"} 0.000186979
go_gc_duration_seconds{quantile="0.5"} 0.000310535
go_gc_duration_seconds{quantile="0.75"} 0.00036327200000000004
go_gc_duration_seconds{quantile="1"} 0.000496432
go_gc_duration_seconds_sum 0.0017815760000000002
go_gc_duration_seconds_count 6
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 16
# HELP http_request_duration_microseconds The HTTP request latencies in microseconds.
# TYPE http_request_duration_microseconds summary
http_request_duration_microseconds{handler="prometheus",quantile="0.5"} NaN
http_request_duration_microseconds{handler="prometheus",quantile="0.9"} NaN
http_request_duration_microseconds{handler="prometheus",quantile="0.99"} NaN
http_request_duration_microseconds_sum{handler="prometheus"} 0
http_request_duration_microseconds_count{handler="prometheus"} 0
# HELP http_request_size_bytes The HTTP request sizes in bytes.
# TYPE http_request_size_bytes summary
http_request_size_bytes{handler="prometheus",quantile="0.5"} NaN
http_request_size_bytes{handler="prometheus",quantile="0.9"} NaN
http_request_size_bytes{handler="prometheus",quantile="0.99"} NaN
http_request_size_bytes_sum{handler="prometheus"} 0
http_request_size_bytes_count{handler="prometheus"} 0
# HELP http_response_size_bytes The HTTP response sizes in bytes.
# TYPE http_response_size_bytes summary
http_response_size_bytes{handler="prometheus",quantile="0.5"} NaN
http_response_size_bytes{handler="prometheus",quantile="0.9"} NaN
http_response_size_bytes{handler="prometheus",quantile="0.99"} NaN
http_response_size_bytes_sum{handler="prometheus"} 0
http_response_size_bytes_count{handler="prometheus"} 0
# HELP my_group_string_service_count_result The result of each count method.
# TYPE my_group_string_service_count_result summary
my_group_string_service_count_result{error="unknown",method="unknown",quantile="0.5"} 12
my_group_string_service_count_result{error="unknown",method="unknown",quantile="0.9"} 12
my_group_string_service_count_result{error="unknown",method="unknown",quantile="0.99"} 12
my_group_string_service_count_result_sum{error="unknown",method="unknown"} 12
my_group_string_service_count_result_count{error="unknown",method="unknown"} 1
# HELP my_group_string_service_request_count Number of requests received.
# TYPE my_group_string_service_request_count counter
my_group_string_service_request_count{error="<nil>",method="count"} 1
my_group_string_service_request_count{error="<nil>",method="uppercase"} 1
# HELP my_group_string_service_request_latency_microseconds Total duration of requests in microseconds.
# TYPE my_group_string_service_request_latency_microseconds summary
my_group_string_service_request_latency_microseconds{error="<nil>",method="count",quantile="0.5"} 88
my_group_string_service_request_latency_microseconds{error="<nil>",method="count",quantile="0.9"} 88
my_group_string_service_request_latency_microseconds{error="<nil>",method="count",quantile="0.99"} 88
my_group_string_service_request_latency_microseconds_sum{error="<nil>",method="count"} 88
my_group_string_service_request_latency_microseconds_count{error="<nil>",method="count"} 1
my_group_string_service_request_latency_microseconds{error="<nil>",method="uppercase",quantile="0.5"} 97
my_group_string_service_request_latency_microseconds{error="<nil>",method="uppercase",quantile="0.9"} 97
my_group_string_service_request_latency_microseconds{error="<nil>",method="uppercase",quantile="0.99"} 97
my_group_string_service_request_latency_microseconds_sum{error="<nil>",method="uppercase"} 97
my_group_string_service_request_latency_microseconds_count{error="<nil>",method="uppercase"} 1
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 1024
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 7
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 4.046848e+06
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.44248862524e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 3.3214464e+07
```
