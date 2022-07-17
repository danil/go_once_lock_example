# Once lock example for Go

[![Build Status](https://cloud.drone.io/api/badges/danil/oncelockexample/status.svg)](https://cloud.drone.io/danil/oncelockexample)
[![Go Reference](https://pkg.go.dev/badge/github.com/danil/oncelockexample.svg)](https://pkg.go.dev/github.com/danil/oncelockexample)

Source files are distributed under the BSD-style license.

## About

The software is considered to be at a alpha level of readiness,
its extremely slow and allocates a lots of memory.

## Benchmark

```sh
$ go test -count=1 -race -bench ./... 
goos: linux
goarch: amd64
pkg: github.com/danil/oncelockexample
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkAtomic/atomic-8            13257363            84.48 ns/op
BenchmarkUberAtomic/uber_atomic-8           10779932           111.4 ns/op
BenchmarkMutex/mutex-8                       4926771           244.5 ns/op
BenchmarkRWMutex/read-write_mutex-8          5244938           230.2 ns/op
BenchmarkChannel/channel-8                   9048637           125.2 ns/op
PASS
ok      github.com/danil/oncelockexample    6.723s
```
