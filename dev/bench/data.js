window.BENCHMARK_DATA = {
  "lastUpdate": 1776871168178,
  "repoUrl": "https://github.com/MichaelMure/benchspotter",
  "entries": {
    "Go Benchmark": [
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "20d42e74500e5df57982cd6db9bf8c288aa87298",
          "message": "some simplification",
          "timestamp": "2026-04-22T14:00:43+02:00",
          "tree_id": "207c73a6f592e20ea5521d130a46f09dc9e64dee",
          "url": "https://github.com/MichaelMure/benchspotter/commit/20d42e74500e5df57982cd6db9bf8c288aa87298"
        },
        "date": 1776871167261,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 178.6,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7513910 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 178.6,
            "unit": "ns/op",
            "extra": "7513910 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7513910 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7513910 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 198.9,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "6394071 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 198.9,
            "unit": "ns/op",
            "extra": "6394071 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "6394071 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "6394071 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 190.8,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7042672 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 190.8,
            "unit": "ns/op",
            "extra": "7042672 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7042672 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7042672 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 7.484,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "160178012 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 7.484,
            "unit": "ns/op",
            "extra": "160178012 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "160178012 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "160178012 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 7.455,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "164093800 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 7.455,
            "unit": "ns/op",
            "extra": "164093800 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "164093800 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "164093800 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.745,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "162330300 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.745,
            "unit": "ns/op",
            "extra": "162330300 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "162330300 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "162330300 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.474,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "160259538 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.474,
            "unit": "ns/op",
            "extra": "160259538 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "160259538 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "160259538 times\n2 procs"
          }
        ]
      }
    ]
  }
}