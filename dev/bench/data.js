window.BENCHMARK_DATA = {
  "lastUpdate": 1776988174834,
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
      },
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
          "id": "2999cfc4b649048286d2f76ea329514f720e83d2",
          "message": "implement optimize with multiple strategies",
          "timestamp": "2026-04-24T01:48:29+02:00",
          "tree_id": "13d5833355cb41989849585e774f22c76f92b6ed",
          "url": "https://github.com/MichaelMure/benchspotter/commit/2999cfc4b649048286d2f76ea329514f720e83d2"
        },
        "date": 1776988174314,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 142.1,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7307590 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 142.1,
            "unit": "ns/op",
            "extra": "7307590 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7307590 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7307590 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 142.3,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7908082 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 142.3,
            "unit": "ns/op",
            "extra": "7908082 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7908082 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7908082 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 156.6,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7656878 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 156.6,
            "unit": "ns/op",
            "extra": "7656878 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7656878 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7656878 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 7.831,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "154378291 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 7.831,
            "unit": "ns/op",
            "extra": "154378291 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "154378291 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "154378291 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 8.067,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "153593222 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 8.067,
            "unit": "ns/op",
            "extra": "153593222 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "153593222 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "153593222 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.818,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "157906092 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.818,
            "unit": "ns/op",
            "extra": "157906092 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "157906092 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "157906092 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.815,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "157740163 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.815,
            "unit": "ns/op",
            "extra": "157740163 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "157740163 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "157740163 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/examples/tuning)",
            "value": 533803,
            "unit": "ns/op\t       2 B/op\t       0 allocs/op",
            "extra": "2196 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/examples/tuning) - ns/op",
            "value": 533803,
            "unit": "ns/op",
            "extra": "2196 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/examples/tuning) - B/op",
            "value": 2,
            "unit": "B/op",
            "extra": "2196 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/examples/tuning) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2196 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/examples/tuning)",
            "value": 155158,
            "unit": "ns/op\t  295553 B/op\t      33 allocs/op",
            "extra": "7504 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/examples/tuning) - ns/op",
            "value": 155158,
            "unit": "ns/op",
            "extra": "7504 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/examples/tuning) - B/op",
            "value": 295553,
            "unit": "B/op",
            "extra": "7504 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/examples/tuning) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "7504 times\n2 procs"
          }
        ]
      }
    ]
  }
}