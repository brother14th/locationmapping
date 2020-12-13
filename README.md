# Location Mapping Sample API

This is a location mapping sample API. It's written in Go. 

Users can authenticate himself/herself and get a JWT. He/She can access other API services using this token. 

The token is valid for 20 minutes.  

However, to accessing Prometheus metrics, basic authentication is needed.  

GoMock package is used to facilitate the testing 

Database is built on SQLite. There are three tables, namely users, locations and location_reports.

All the tables have been prepopulated with the following data

**<u>Sample data in SQLite database</u>**

**<u>users table</u>**


| username | password  | preferred_location |
| :------- | --------- | ------------------ |
| user1    | password1 | location1          |
| user2    | password2 | location 2         |

**locations table**

| name      |
| :-------- |
| location1 |
| location2 |
| location3 |
| location4 |

**location_reports  table**

| location  | total_square_feet | price_per_month |
| :-------- | ----------------- | --------------- |
| location1 | 1000              | 3800            |
| location2 | 1100              | 3900            |
| location3 | 1200              | 4000            |
| location4 | 1300              | 4100            |



## Steps to run the API

1. Download and install GCC (http://mingw-w64.org/doku.php) 

   Note: GCC is required for access SQLite database. For windows, you can choose x86_64-win32-sjlj installer

2. Add GCC bin directory to PATH (e.g. C:\Program Files\mingw-w64\x86_64-8.1.0-win32-sjlj-rt_v6-rev0\mingw64\bin)

3. Run go get -v github.com/brother14th/locationmapping

4. Go to local go package directory (e.g. %gopath%\src\github.com\brother14th\locationmapping\db)

5. Open  userrepository.go, modify the path  C:/Users/hngkh/go/src/github.com/brother14th/locationmapping/db/locationmapping.db 

6. Open locationrepository.go, modify the path  C:/Users/hngkh/go/src/github.com/brother14th/locationmapping/db/locationmapping.db

7. Enter  locationmapping in command prompt

8. Use curl or Postman to access the API services (refer to REST API section below)

   

# REST API

The REST API  is described below.

## Authenticate User

- To authenticate user's credential and return JWT token

### Request

`POST /v1/authenticate`

    curl -i -X POST -H "Content-Type: application/json"  -d "{\"username\":\"user1\", \"password\":\"password1\"}" http://localhost:8080/v1/authenticate

### Response

    HTTP/1.1 200 OK
    Date: Sun, 13 Dec 2020 12:45:12 GMT
    Content-Length: 137
    Content-Type: text/plain; charset=utf-8
    
    {"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDc4NjQ4NTYsInN1YiI6InVzZXIxIn0.17AwlPIqY023Iu-0rnM6vKL0l9TPv0lm4XdScdEkHXg"}


## Set user's preferred location

- To set user's preferred location for himself/herself. User's name will be extracted from token.  

### Request

`PATCH /v1/user`

    curl -i -X PATCH -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDc4NjQ4NTYsInN1YiI6InVzZXIxIn0.17AwlPIqY023Iu-0rnM6vKL0l9TPv0lm4XdScdEkHXg" -d "{\"location\":\"location4\"}" http://localhost:8080/v1/user

### Response

    HTTP/1.1 200 OK
    Date: Sun, 13 Dec 2020 12:48:08 GMT
    Content-Length: 16
    Content-Type: text/plain; charset=utf-8
    
    {"status":true}


## Get user's preferred location

- To get user's preferred location. User's name will be extracted from token.  

### Request

`GET /v1/preferredlocation`

    curl -i -X GET -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDc4NjQ4NTYsInN1YiI6InVzZXIxIn0.17AwlPIqY023Iu-0rnM6vKL0l9TPv0lm4XdScdEkHXg" http://localhost:8080/v1/preferredlocation

### Response

    HTTP/1.1 200 OK
    Date: Sun, 13 Dec 2020 12:54:06 GMT
    Content-Length: 25
    Content-Type: text/plain; charset=utf-8
    
    {"Location":"location4"}


## Get location's report/summary for a selected location

- To get location's report/summary based on selected location

### Request

`GET /v1/locationreport?location=selectedlocation`

    curl -i -X GET -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDc4NjQ4NTYsInN1YiI6InVzZXIxIn0.17AwlPIqY023Iu-0rnM6vKL0l9TPv0lm4XdScdEkHXg" http://localhost:8080/v1/locationreport?location=location4

### Response

    HTTP/1.1 200 OK
    Date: Sun, 13 Dec 2020 12:57:46 GMT
    Content-Length: 175
    Content-Type: text/plain; charset=utf-8
    
    {"summary":{"ID":4,"CreatedAt":"2020-12-11T00:00:00Z",
    "UpdatedAt":"2020-12-11T00:00:00Z","DeletedAt":null,"Location":"location4",
    "TotalSquareFeet":1300,"PricePerMonth":4100}}


## Get Prometheus metrics

- To get Prometheus metrics

### Request

`GET /metrics`

    curl -u prometheus:password http://localhost:8080/metrics

### Response

    # HELP api_auth_service_request_count Number of requests received.
    # TYPE api_auth_service_request_count counter
    api_auth_service_request_count{error="false",method="authenticate"} 15
    # HELP api_auth_service_request_latency_microseconds Total duration of requests in microseconds.
    # TYPE api_auth_service_request_latency_microseconds summary
    api_auth_service_request_latency_microseconds_sum{error="false",method="authenticate"} 0.0081086
    api_auth_service_request_latency_microseconds_count{error="false",method="authenticate"} 15
    # HELP api_location_report_service_request_count Number of requests received.
    # TYPE api_location_report_service_request_count counter
    api_location_report_service_request_count{error="false",method="GetLocationReport"} 6
    # HELP api_location_report_service_request_latency_microseconds Total duration of requests in microseconds.
    # TYPE api_location_report_service_request_latency_microseconds summary
    api_location_report_service_request_latency_microseconds_sum{error="false",method="GetLocationReport"} 0.0016370999999999998
    api_location_report_service_request_latency_microseconds_count{error="false",method="GetLocationReport"} 6
    # HELP api_user_service_request_count Number of requests received.
    # TYPE api_user_service_request_count counter
    api_user_service_request_count{error="false",method="GetPreferredLocation"} 8
    api_user_service_request_count{error="false",method="SetPreferredLocation"} 5
    # HELP api_user_service_request_latency_microseconds Total duration of requests in microseconds.
    # TYPE api_user_service_request_latency_microseconds summary
    api_user_service_request_latency_microseconds_sum{error="false",method="GetPreferredLocation"} 0.0012837999999999999
    api_user_service_request_latency_microseconds_count{error="false",method="GetPreferredLocation"} 8
    api_user_service_request_latency_microseconds_sum{error="false",method="SetPreferredLocation"} 0.0513669
    api_user_service_request_latency_microseconds_count{error="false",method="SetPreferredLocation"} 5
    # HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
    # TYPE go_gc_duration_seconds summary
    go_gc_duration_seconds{quantile="0"} 0
    go_gc_duration_seconds{quantile="0.25"} 0
    go_gc_duration_seconds{quantile="0.5"} 0
    go_gc_duration_seconds{quantile="0.75"} 0
    go_gc_duration_seconds{quantile="1"} 0
    go_gc_duration_seconds_sum 0
    go_gc_duration_seconds_count 0
    # HELP go_goroutines Number of goroutines that currently exist.
    # TYPE go_goroutines gauge
    go_goroutines 11
    # HELP go_info Information about the Go environment.
    # TYPE go_info gauge
    go_info{version="go1.15.6"} 1
    # HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
    # TYPE go_memstats_alloc_bytes gauge
    go_memstats_alloc_bytes 3.819336e+06
    # HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
    # TYPE go_memstats_alloc_bytes_total counter
    go_memstats_alloc_bytes_total 3.819336e+06
    # HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
    # TYPE go_memstats_buck_hash_sys_bytes gauge
    go_memstats_buck_hash_sys_bytes 1.445154e+06
    # HELP go_memstats_frees_total Total number of frees.
    # TYPE go_memstats_frees_total counter
    go_memstats_frees_total 2813
    # HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
    # TYPE go_memstats_gc_cpu_fraction gauge
    go_memstats_gc_cpu_fraction 0
    # HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
    # TYPE go_memstats_gc_sys_bytes gauge
    go_memstats_gc_sys_bytes 2.338496e+06
    # HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
    # TYPE go_memstats_heap_alloc_bytes gauge
    go_memstats_heap_alloc_bytes 3.819336e+06
    # HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
    # TYPE go_memstats_heap_idle_bytes gauge
    go_memstats_heap_idle_bytes 3.047424e+06
    # HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
    # TYPE go_memstats_heap_inuse_bytes gauge
    go_memstats_heap_inuse_bytes 4.9152e+06
    # HELP go_memstats_heap_objects Number of allocated objects.
    # TYPE go_memstats_heap_objects gauge
    go_memstats_heap_objects 27771
    # HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
    # TYPE go_memstats_heap_released_bytes gauge
    go_memstats_heap_released_bytes 3.022848e+06
    # HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
    # TYPE go_memstats_heap_sys_bytes gauge
    go_memstats_heap_sys_bytes 7.962624e+06
    # HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
    # TYPE go_memstats_last_gc_time_seconds gauge
    go_memstats_last_gc_time_seconds 0
    # HELP go_memstats_lookups_total Total number of pointer lookups.
    # TYPE go_memstats_lookups_total counter
    go_memstats_lookups_total 0
    # HELP go_memstats_mallocs_total Total number of mallocs.
    # TYPE go_memstats_mallocs_total counter
    go_memstats_mallocs_total 30584
    # HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
    # TYPE go_memstats_mcache_inuse_bytes gauge
    go_memstats_mcache_inuse_bytes 13632
    # HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
    # TYPE go_memstats_mcache_sys_bytes gauge
    go_memstats_mcache_sys_bytes 16384
    # HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
    # TYPE go_memstats_mspan_inuse_bytes gauge
    go_memstats_mspan_inuse_bytes 81056
    # HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
    # TYPE go_memstats_mspan_sys_bytes gauge
    go_memstats_mspan_sys_bytes 81920
    # HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
    # TYPE go_memstats_next_gc_bytes gauge
    go_memstats_next_gc_bytes 4.473924e+06
    # HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
    # TYPE go_memstats_other_sys_bytes gauge
    go_memstats_other_sys_bytes 1.25303e+06
    # HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
    # TYPE go_memstats_stack_inuse_bytes gauge
    go_memstats_stack_inuse_bytes 425984
    # HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
    # TYPE go_memstats_stack_sys_bytes gauge
    go_memstats_stack_sys_bytes 425984
    # HELP go_memstats_sys_bytes Number of bytes obtained from system.
    # TYPE go_memstats_sys_bytes gauge
    go_memstats_sys_bytes 1.3523592e+07
    # HELP go_threads Number of OS threads created.
    # TYPE go_threads gauge
    go_threads 9
    # HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
    # TYPE process_cpu_seconds_total counter
    process_cpu_seconds_total 0.109375
    # HELP process_max_fds Maximum number of open file descriptors.
    # TYPE process_max_fds gauge
    process_max_fds 1.6777216e+07
    # HELP process_open_fds Number of open file descriptors.
    # TYPE process_open_fds gauge
    process_open_fds 127
    # HELP process_resident_memory_bytes Resident memory size in bytes.
    # TYPE process_resident_memory_bytes gauge
    process_resident_memory_bytes 1.0596352e+07
    # HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
    # TYPE process_start_time_seconds gauge
    process_start_time_seconds 1.607852979e+09
    # HELP process_virtual_memory_bytes Virtual memory size in bytes.
    # TYPE process_virtual_memory_bytes gauge
    process_virtual_memory_bytes 2.1127168e+07
    # HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
    # TYPE promhttp_metric_handler_requests_in_flight gauge
    promhttp_metric_handler_requests_in_flight 1
    # HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
    # TYPE promhttp_metric_handler_requests_total counter
    promhttp_metric_handler_requests_total{code="200"} 5
    promhttp_metric_handler_requests_total{code="500"} 0
    promhttp_metric_handler_requests_total{code="503"} 0


​    

