# Benchmark for Logging With Zap

## Benchmark Results
on my MacBook Pro
```
➜  TestLoggingWithZap git:(master) ✗ go test -bench=. 
goos: darwin
goarch: amd64
pkg: github.com/rezakamalifard/TestLoggingWithZap
BenchmarkSugarLogger-4            793906              1260 ns/op             512 B/op          1 allocs/op
BenchmarkStructuredLogger-4      1262098               950 ns/op             256 B/op          1 allocs/op
PASS
ok      github.com/rezakamalifard/TestLoggingWithZap    3.575s
```
