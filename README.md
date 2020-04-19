# Redis for Golang

learn

##  Perf
```text
➜  ~ wrk -t12 -c200 -d30s http://127.0.0.1:8080
Running 30s test @ http://127.0.0.1:8080
  12 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.57ms    2.17ms  79.47ms   98.92%
    Req/Sec     3.45k   303.20     5.91k    82.87%
  1237660 requests in 30.10s, 153.44MB read
  Socket errors: connect 0, read 37, write 0, timeout 0
Requests/sec:  41112.81
Transfer/sec:      5.10MB
```