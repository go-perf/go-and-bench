# go-and-bench

[![build-img]][build-url]

Collection of Go benchmarks. Literally "go and bench".

## Rationale

To have an easy way to run Go benchmarks to see how different code behaves on a specific hardware and/or OS.

## Usage

Install `Taskfile` first, see [https://taskfile.dev/installation/](https://taskfile.dev/installation/).

Compile benchmarks:

```bash
$ task compile 
```

Run benchmarks:

```bash
$ task bench
```

[build-img]: https://github.com/go-perf/go-and-bench/workflows/build/badge.svg
[build-url]: https://github.com/go-perf/go-and-bench/actions
