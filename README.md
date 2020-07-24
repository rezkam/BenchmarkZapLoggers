# Benchmark for Logging With Zap

## Logger Types
### Discarded Logger
Sends all the writes to ioutil.Discard
### Sampled Logger
Creates a zap core that samples incoming N (here is 10) entries, with a given level and message
each tick. If more Entries with the same level and message are seen during the same interval, every Mth message is logged and the rest are dropped.
### Writer Logger
Sneds all the writes to a writer (stderr) without sampling.
### WriterSampler
Writer Logger with a sampler.

## Benchmark Results
```
BenchmarkDiscardedSugarLogger                817209           1292 ns/op         512 B/op          1 allocs/op
BenchmarkDiscardedStructuredLogger           1000000          1010 ns/op         256 B/op          1 allocs/op
BenchmarkSampledSugarLogger                  3972433          277 ns/op          51 B/op           0 allocs/op
BenchmarkSampledStructuredLogger             3486465          356 ns/op          256 B/op          1 allocs/op
BenchmarkWriterSugarLogger                   85296            14986 ns/op        760 B/op          4 allocs/op
BenchmarkWriterStructuredLogger              57351            17587 ns/op        504 B/op          4 allocs/op
BenchmarkWriterSamplerSugarLogger            3176926          563 ns/op          7 B/op            0 allocs/op
BenchmarkWriterSamplerStructuredLogger       2051127          610 ns/op          258 B/op          1 allocs/op
```
