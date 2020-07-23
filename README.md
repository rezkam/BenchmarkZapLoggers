# Benchmark for Logging With Zap

## Benchmark Results
on my MacBook Pro
```
➜  BenchmarkZapLoggers git:(master) ✗ go test -bench=. 
goos: darwin
goarch: amd64
pkg: github.com/rezakamalifard/BenchmarkZapLoggers
BenchmarkSugarLogger-4            793906              1260 ns/op             512 B/op          1 allocs/op
BenchmarkStructuredLogger-4      1262098               950 ns/op             256 B/op          1 allocs/op
PASS
ok      github.com/rezakamalifard/BenchmarkZapLoggers    3.575s
```
