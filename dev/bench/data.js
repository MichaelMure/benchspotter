window.BENCHMARK_DATA = {
  "lastUpdate": 1778773505305,
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
          "id": "4b8c7fcd510aef8d23a7e43804d2a95a5d85ed3f",
          "message": "implement optimize with multiple strategies",
          "timestamp": "2026-04-24T01:50:40+02:00",
          "tree_id": "753e76bd89eb8029c516b8a81ec87c0a47c7ff4c",
          "url": "https://github.com/MichaelMure/benchspotter/commit/4b8c7fcd510aef8d23a7e43804d2a95a5d85ed3f"
        },
        "date": 1776988287213,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 145,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "8272362 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 145,
            "unit": "ns/op",
            "extra": "8272362 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "8272362 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8272362 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 140.4,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "8788302 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 140.4,
            "unit": "ns/op",
            "extra": "8788302 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "8788302 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8788302 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 149.7,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7888240 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 149.7,
            "unit": "ns/op",
            "extra": "7888240 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7888240 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7888240 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 7.811,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "153639667 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 7.811,
            "unit": "ns/op",
            "extra": "153639667 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "153639667 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "153639667 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 7.789,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "153224145 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 7.789,
            "unit": "ns/op",
            "extra": "153224145 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "153224145 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "153224145 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal)",
            "value": 538755,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2223 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - ns/op",
            "value": 538755,
            "unit": "ns/op",
            "extra": "2223 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2223 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2223 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal)",
            "value": 170940,
            "unit": "ns/op\t  295552 B/op\t      33 allocs/op",
            "extra": "7311 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - ns/op",
            "value": 170940,
            "unit": "ns/op",
            "extra": "7311 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - B/op",
            "value": 295552,
            "unit": "B/op",
            "extra": "7311 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "7311 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.58,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "153781189 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.58,
            "unit": "ns/op",
            "extra": "153781189 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "153781189 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "153781189 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.703,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "158642552 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.703,
            "unit": "ns/op",
            "extra": "158642552 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "158642552 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "158642552 times\n2 procs"
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
          "id": "1b0842fc34f0511e1bbe2ae4d8f059ae4f759797",
          "message": "inline analysis + general improvement on escape/inline UI",
          "timestamp": "2026-04-25T13:13:40+02:00",
          "tree_id": "d6ba696fe621b4fec7344c5d1cde719ed0714ac6",
          "url": "https://github.com/MichaelMure/benchspotter/commit/1b0842fc34f0511e1bbe2ae4d8f059ae4f759797"
        },
        "date": 1777115693071,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 159.7,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "6774448 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 159.7,
            "unit": "ns/op",
            "extra": "6774448 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "6774448 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "6774448 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 159.6,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7761378 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 159.6,
            "unit": "ns/op",
            "extra": "7761378 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7761378 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7761378 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 173.6,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "6933956 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 173.6,
            "unit": "ns/op",
            "extra": "6933956 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "6933956 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "6933956 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 7.361,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "162284282 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 7.361,
            "unit": "ns/op",
            "extra": "162284282 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "162284282 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "162284282 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 7.473,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "163619804 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 7.473,
            "unit": "ns/op",
            "extra": "163619804 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "163619804 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "163619804 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal)",
            "value": 503554,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2474 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - ns/op",
            "value": 503554,
            "unit": "ns/op",
            "extra": "2474 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2474 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2474 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal)",
            "value": 188134,
            "unit": "ns/op\t  295552 B/op\t      33 allocs/op",
            "extra": "6304 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - ns/op",
            "value": 188134,
            "unit": "ns/op",
            "extra": "6304 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - B/op",
            "value": 295552,
            "unit": "B/op",
            "extra": "6304 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "6304 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.312,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "164359660 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.312,
            "unit": "ns/op",
            "extra": "164359660 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "164359660 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "164359660 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.466,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "160493715 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.466,
            "unit": "ns/op",
            "extra": "160493715 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "160493715 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "160493715 times\n2 procs"
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
          "id": "d7025f9a56aa3dad238cf1f314ff7bcb3ba081a0",
          "message": "move to charm/* v2, cleanup code",
          "timestamp": "2026-04-28T23:45:41+02:00",
          "tree_id": "9361487e45b9bd56e9e9055245d3a0dc12fd9848",
          "url": "https://github.com/MichaelMure/benchspotter/commit/d7025f9a56aa3dad238cf1f314ff7bcb3ba081a0"
        },
        "date": 1777413002926,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 163.8,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "6493503 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 163.8,
            "unit": "ns/op",
            "extra": "6493503 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "6493503 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "6493503 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 164.2,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "8077081 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 164.2,
            "unit": "ns/op",
            "extra": "8077081 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "8077081 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8077081 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 163.9,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "6464910 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 163.9,
            "unit": "ns/op",
            "extra": "6464910 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "6464910 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "6464910 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 7.347,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "161723646 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 7.347,
            "unit": "ns/op",
            "extra": "161723646 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "161723646 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "161723646 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 7.316,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "163946380 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 7.316,
            "unit": "ns/op",
            "extra": "163946380 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "163946380 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "163946380 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal)",
            "value": 486155,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2473 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - ns/op",
            "value": 486155,
            "unit": "ns/op",
            "extra": "2473 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2473 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2473 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal)",
            "value": 182899,
            "unit": "ns/op\t  295552 B/op\t      33 allocs/op",
            "extra": "6306 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - ns/op",
            "value": 182899,
            "unit": "ns/op",
            "extra": "6306 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - B/op",
            "value": 295552,
            "unit": "B/op",
            "extra": "6306 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "6306 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.712,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "163160860 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.712,
            "unit": "ns/op",
            "extra": "163160860 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "163160860 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "163160860 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.521,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "150022071 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.521,
            "unit": "ns/op",
            "extra": "150022071 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "150022071 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "150022071 times\n2 procs"
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
          "id": "2f0d8733c0f6617810247abfc4fe8b62e8098d88",
          "message": "fix tests",
          "timestamp": "2026-04-29T00:56:19+02:00",
          "tree_id": "574cc4e01c34409e6f2d6f16aad9e41138a5480c",
          "url": "https://github.com/MichaelMure/benchspotter/commit/2f0d8733c0f6617810247abfc4fe8b62e8098d88"
        },
        "date": 1777417035418,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 118.6,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "9944482 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 118.6,
            "unit": "ns/op",
            "extra": "9944482 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "9944482 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9944482 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 115.6,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "10773092 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 115.6,
            "unit": "ns/op",
            "extra": "10773092 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "10773092 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10773092 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 135.6,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "8681334 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 135.6,
            "unit": "ns/op",
            "extra": "8681334 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "8681334 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8681334 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 6.034,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "196217098 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 6.034,
            "unit": "ns/op",
            "extra": "196217098 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "196217098 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "196217098 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 6.049,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "190713180 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 6.049,
            "unit": "ns/op",
            "extra": "190713180 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "190713180 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "190713180 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal)",
            "value": 413181,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2865 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - ns/op",
            "value": 413181,
            "unit": "ns/op",
            "extra": "2865 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2865 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2865 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal)",
            "value": 121140,
            "unit": "ns/op\t  295552 B/op\t      33 allocs/op",
            "extra": "8932 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - ns/op",
            "value": 121140,
            "unit": "ns/op",
            "extra": "8932 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - B/op",
            "value": 295552,
            "unit": "B/op",
            "extra": "8932 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "8932 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 5.861,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "204221976 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 5.861,
            "unit": "ns/op",
            "extra": "204221976 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "204221976 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "204221976 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 6.015,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "198816897 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 6.015,
            "unit": "ns/op",
            "extra": "198816897 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "198816897 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "198816897 times\n2 procs"
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
          "id": "32f7e139640bb38c43302c047bc4beb2e5c3641c",
          "message": "move color handling into execenv.Style",
          "timestamp": "2026-04-29T09:21:00+02:00",
          "tree_id": "d1a3a6cc2e5d8f5d94e0a7af5f3be878475caead",
          "url": "https://github.com/MichaelMure/benchspotter/commit/32f7e139640bb38c43302c047bc4beb2e5c3641c"
        },
        "date": 1777447385595,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 152.2,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7030714 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 152.2,
            "unit": "ns/op",
            "extra": "7030714 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7030714 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7030714 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 149.7,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7847220 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 149.7,
            "unit": "ns/op",
            "extra": "7847220 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7847220 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7847220 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 168.6,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7422243 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 168.6,
            "unit": "ns/op",
            "extra": "7422243 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7422243 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7422243 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 7.48,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "162685690 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 7.48,
            "unit": "ns/op",
            "extra": "162685690 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "162685690 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "162685690 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 7.688,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "163887816 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 7.688,
            "unit": "ns/op",
            "extra": "163887816 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "163887816 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "163887816 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal)",
            "value": 482997,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2476 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - ns/op",
            "value": 482997,
            "unit": "ns/op",
            "extra": "2476 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2476 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2476 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal)",
            "value": 170343,
            "unit": "ns/op\t  295553 B/op\t      33 allocs/op",
            "extra": "6598 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - ns/op",
            "value": 170343,
            "unit": "ns/op",
            "extra": "6598 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - B/op",
            "value": 295553,
            "unit": "B/op",
            "extra": "6598 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "6598 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "163367427 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.5,
            "unit": "ns/op",
            "extra": "163367427 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "163367427 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "163367427 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.336,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "162770499 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.336,
            "unit": "ns/op",
            "extra": "162770499 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "162770499 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "162770499 times\n2 procs"
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
          "id": "676695210feb46de4a680d425b4f4d9309702d44",
          "message": "fix go's tabwriter not handling ansi colors",
          "timestamp": "2026-04-29T22:47:04+02:00",
          "tree_id": "4274630ea51a17d6b217ca6c366d9dec929a7c34",
          "url": "https://github.com/MichaelMure/benchspotter/commit/676695210feb46de4a680d425b4f4d9309702d44"
        },
        "date": 1777495788982,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter)",
            "value": 1909,
            "unit": "ns/op\t    2144 B/op\t      31 allocs/op",
            "extra": "554290 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 1909,
            "unit": "ns/op",
            "extra": "554290 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 2144,
            "unit": "B/op",
            "extra": "554290 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "554290 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter)",
            "value": 1050,
            "unit": "ns/op\t     640 B/op\t      10 allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 1050,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter)",
            "value": 139699,
            "unit": "ns/op\t  174944 B/op\t    2024 allocs/op",
            "extra": "8397 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 139699,
            "unit": "ns/op",
            "extra": "8397 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 174944,
            "unit": "B/op",
            "extra": "8397 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 2024,
            "unit": "allocs/op",
            "extra": "8397 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 100181,
            "unit": "ns/op\t   64011 B/op\t    1000 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 100181,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64011,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter)",
            "value": 29369047,
            "unit": "ns/op\t24685027 B/op\t  200056 allocs/op",
            "extra": "39 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 29369047,
            "unit": "ns/op",
            "extra": "39 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 24685027,
            "unit": "B/op",
            "extra": "39 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 200056,
            "unit": "allocs/op",
            "extra": "39 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 11810545,
            "unit": "ns/op\t 6582850 B/op\t  101000 allocs/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 11810545,
            "unit": "ns/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6582850,
            "unit": "B/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101000,
            "unit": "allocs/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter)",
            "value": 9487,
            "unit": "ns/op\t   11440 B/op\t     131 allocs/op",
            "extra": "129904 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 9487,
            "unit": "ns/op",
            "extra": "129904 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 11440,
            "unit": "B/op",
            "extra": "129904 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "129904 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter)",
            "value": 6968,
            "unit": "ns/op\t    6400 B/op\t     100 allocs/op",
            "extra": "171936 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 6968,
            "unit": "ns/op",
            "extra": "171936 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6400,
            "unit": "B/op",
            "extra": "171936 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "171936 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter)",
            "value": 830204,
            "unit": "ns/op\t 1035121 B/op\t   11038 allocs/op",
            "extra": "1472 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 830204,
            "unit": "ns/op",
            "extra": "1472 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 1035121,
            "unit": "B/op",
            "extra": "1472 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 11038,
            "unit": "allocs/op",
            "extra": "1472 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 705413,
            "unit": "ns/op\t  640244 B/op\t   10000 allocs/op",
            "extra": "1644 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 705413,
            "unit": "ns/op",
            "extra": "1644 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640244,
            "unit": "B/op",
            "extra": "1644 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10000,
            "unit": "allocs/op",
            "extra": "1644 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter)",
            "value": 108451021,
            "unit": "ns/op\t111012768 B/op\t 1100072 allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 108451021,
            "unit": "ns/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 111012768,
            "unit": "B/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1100072,
            "unit": "allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 82714535,
            "unit": "ns/op\t67358052 B/op\t 1007148 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 82714535,
            "unit": "ns/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 67358052,
            "unit": "B/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1007148,
            "unit": "allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter)",
            "value": 76604,
            "unit": "ns/op\t  103280 B/op\t    1041 allocs/op",
            "extra": "15514 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 76604,
            "unit": "ns/op",
            "extra": "15514 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 103280,
            "unit": "B/op",
            "extra": "15514 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1041,
            "unit": "allocs/op",
            "extra": "15514 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter)",
            "value": 64373,
            "unit": "ns/op\t   64002 B/op\t    1000 allocs/op",
            "extra": "18572 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 64373,
            "unit": "ns/op",
            "extra": "18572 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64002,
            "unit": "B/op",
            "extra": "18572 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "18572 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter)",
            "value": 8515441,
            "unit": "ns/op\t 9670516 B/op\t  101052 allocs/op",
            "extra": "138 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 8515441,
            "unit": "ns/op",
            "extra": "138 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 9670516,
            "unit": "B/op",
            "extra": "138 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101052,
            "unit": "allocs/op",
            "extra": "138 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 6778439,
            "unit": "ns/op\t 6418540 B/op\t  100006 allocs/op",
            "extra": "177 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 6778439,
            "unit": "ns/op",
            "extra": "177 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6418540,
            "unit": "B/op",
            "extra": "177 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100006,
            "unit": "allocs/op",
            "extra": "177 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter)",
            "value": 808111250,
            "unit": "ns/op\t974042208 B/op\t10100090 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 808111250,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 974042208,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10100090,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 773175928,
            "unit": "ns/op\t807021048 B/op\t10050044 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 773175928,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 807021048,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10050044,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter)",
            "value": 6226,
            "unit": "ns/op\t    7736 B/op\t      80 allocs/op",
            "extra": "180973 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 6226,
            "unit": "ns/op",
            "extra": "180973 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7736,
            "unit": "B/op",
            "extra": "180973 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 80,
            "unit": "allocs/op",
            "extra": "180973 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter)",
            "value": 455530,
            "unit": "ns/op\t  727320 B/op\t    5180 allocs/op",
            "extra": "2620 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 455530,
            "unit": "ns/op",
            "extra": "2620 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 727320,
            "unit": "B/op",
            "extra": "2620 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 5180,
            "unit": "allocs/op",
            "extra": "2620 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter)",
            "value": 43786311,
            "unit": "ns/op\t67574810 B/op\t  501554 allocs/op",
            "extra": "31 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 43786311,
            "unit": "ns/op",
            "extra": "31 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 67574810,
            "unit": "B/op",
            "extra": "31 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 501554,
            "unit": "allocs/op",
            "extra": "31 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter)",
            "value": 6648,
            "unit": "ns/op\t    7888 B/op\t      87 allocs/op",
            "extra": "171964 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 6648,
            "unit": "ns/op",
            "extra": "171964 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7888,
            "unit": "B/op",
            "extra": "171964 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "171964 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter)",
            "value": 62992,
            "unit": "ns/op\t   78368 B/op\t     750 allocs/op",
            "extra": "19459 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 62992,
            "unit": "ns/op",
            "extra": "19459 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 78368,
            "unit": "B/op",
            "extra": "19459 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 750,
            "unit": "allocs/op",
            "extra": "19459 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter)",
            "value": 657864,
            "unit": "ns/op\t  768978 B/op\t    7285 allocs/op",
            "extra": "1983 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 657864,
            "unit": "ns/op",
            "extra": "1983 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 768978,
            "unit": "B/op",
            "extra": "1983 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 7285,
            "unit": "allocs/op",
            "extra": "1983 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter)",
            "value": 2093,
            "unit": "ns/op\t    1344 B/op\t      33 allocs/op",
            "extra": "567308 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - ns/op",
            "value": 2093,
            "unit": "ns/op",
            "extra": "567308 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - B/op",
            "value": 1344,
            "unit": "B/op",
            "extra": "567308 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "567308 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 156.6,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7417497 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 156.6,
            "unit": "ns/op",
            "extra": "7417497 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7417497 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7417497 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 164.8,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7857399 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 164.8,
            "unit": "ns/op",
            "extra": "7857399 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7857399 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7857399 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 178.7,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7391215 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 178.7,
            "unit": "ns/op",
            "extra": "7391215 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7391215 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7391215 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 7.52,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "158495684 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 7.52,
            "unit": "ns/op",
            "extra": "158495684 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "158495684 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "158495684 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 7.449,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "164072196 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 7.449,
            "unit": "ns/op",
            "extra": "164072196 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "164072196 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "164072196 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal)",
            "value": 500308,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2476 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - ns/op",
            "value": 500308,
            "unit": "ns/op",
            "extra": "2476 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2476 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2476 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal)",
            "value": 183957,
            "unit": "ns/op\t  295552 B/op\t      33 allocs/op",
            "extra": "5977 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - ns/op",
            "value": 183957,
            "unit": "ns/op",
            "extra": "5977 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - B/op",
            "value": 295552,
            "unit": "B/op",
            "extra": "5977 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "5977 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.283,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "164431935 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.283,
            "unit": "ns/op",
            "extra": "164431935 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "164431935 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "164431935 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.536,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "159937210 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.536,
            "unit": "ns/op",
            "extra": "159937210 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "159937210 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "159937210 times\n2 procs"
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
          "id": "ee258cb9c12b7524119c2097a9bf0e2d8fb736a0",
          "message": "long description for all commands",
          "timestamp": "2026-05-01T00:41:07+02:00",
          "tree_id": "206aa2c3e998e710aab0e80135fdf8877c8c5fb0",
          "url": "https://github.com/MichaelMure/benchspotter/commit/ee258cb9c12b7524119c2097a9bf0e2d8fb736a0"
        },
        "date": 1777628359007,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter)",
            "value": 1883,
            "unit": "ns/op\t    2144 B/op\t      31 allocs/op",
            "extra": "614181 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 1883,
            "unit": "ns/op",
            "extra": "614181 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 2144,
            "unit": "B/op",
            "extra": "614181 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "614181 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter)",
            "value": 1045,
            "unit": "ns/op\t     640 B/op\t      10 allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 1045,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter)",
            "value": 152489,
            "unit": "ns/op\t  174944 B/op\t    2024 allocs/op",
            "extra": "7894 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 152489,
            "unit": "ns/op",
            "extra": "7894 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 174944,
            "unit": "B/op",
            "extra": "7894 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 2024,
            "unit": "allocs/op",
            "extra": "7894 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 100021,
            "unit": "ns/op\t   64011 B/op\t    1000 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 100021,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64011,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter)",
            "value": 29124520,
            "unit": "ns/op\t24685057 B/op\t  200056 allocs/op",
            "extra": "54 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 29124520,
            "unit": "ns/op",
            "extra": "54 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 24685057,
            "unit": "B/op",
            "extra": "54 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 200056,
            "unit": "allocs/op",
            "extra": "54 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 11801428,
            "unit": "ns/op\t 6586608 B/op\t  101021 allocs/op",
            "extra": "98 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 11801428,
            "unit": "ns/op",
            "extra": "98 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6586608,
            "unit": "B/op",
            "extra": "98 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101021,
            "unit": "allocs/op",
            "extra": "98 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter)",
            "value": 9732,
            "unit": "ns/op\t   11440 B/op\t     131 allocs/op",
            "extra": "118802 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 9732,
            "unit": "ns/op",
            "extra": "118802 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 11440,
            "unit": "B/op",
            "extra": "118802 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "118802 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter)",
            "value": 7286,
            "unit": "ns/op\t    6400 B/op\t     100 allocs/op",
            "extra": "161254 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 7286,
            "unit": "ns/op",
            "extra": "161254 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6400,
            "unit": "B/op",
            "extra": "161254 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "161254 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter)",
            "value": 893749,
            "unit": "ns/op\t 1035122 B/op\t   11038 allocs/op",
            "extra": "1413 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 893749,
            "unit": "ns/op",
            "extra": "1413 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 1035122,
            "unit": "B/op",
            "extra": "1413 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 11038,
            "unit": "allocs/op",
            "extra": "1413 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 752901,
            "unit": "ns/op\t  640247 B/op\t   10000 allocs/op",
            "extra": "1614 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 752901,
            "unit": "ns/op",
            "extra": "1614 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640247,
            "unit": "B/op",
            "extra": "1614 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10000,
            "unit": "allocs/op",
            "extra": "1614 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter)",
            "value": 100956732,
            "unit": "ns/op\t111012817 B/op\t 1100072 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 100956732,
            "unit": "ns/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 111012817,
            "unit": "B/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1100072,
            "unit": "allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 76611767,
            "unit": "ns/op\t67134204 B/op\t 1006671 allocs/op",
            "extra": "15 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 76611767,
            "unit": "ns/op",
            "extra": "15 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 67134204,
            "unit": "B/op",
            "extra": "15 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1006671,
            "unit": "allocs/op",
            "extra": "15 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter)",
            "value": 75988,
            "unit": "ns/op\t  103280 B/op\t    1041 allocs/op",
            "extra": "16093 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 75988,
            "unit": "ns/op",
            "extra": "16093 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 103280,
            "unit": "B/op",
            "extra": "16093 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1041,
            "unit": "allocs/op",
            "extra": "16093 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter)",
            "value": 64487,
            "unit": "ns/op\t   64002 B/op\t    1000 allocs/op",
            "extra": "18752 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 64487,
            "unit": "ns/op",
            "extra": "18752 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64002,
            "unit": "B/op",
            "extra": "18752 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "18752 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter)",
            "value": 8360111,
            "unit": "ns/op\t 9670581 B/op\t  101052 allocs/op",
            "extra": "145 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 8360111,
            "unit": "ns/op",
            "extra": "145 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 9670581,
            "unit": "B/op",
            "extra": "145 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101052,
            "unit": "allocs/op",
            "extra": "145 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 6673426,
            "unit": "ns/op\t 6419062 B/op\t  100006 allocs/op",
            "extra": "172 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 6673426,
            "unit": "ns/op",
            "extra": "172 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6419062,
            "unit": "B/op",
            "extra": "172 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100006,
            "unit": "allocs/op",
            "extra": "172 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter)",
            "value": 773440888,
            "unit": "ns/op\t974042152 B/op\t10100089 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 773440888,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 974042152,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10100089,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 772204720,
            "unit": "ns/op\t807021048 B/op\t10050044 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 772204720,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 807021048,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10050044,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter)",
            "value": 6485,
            "unit": "ns/op\t    7736 B/op\t      80 allocs/op",
            "extra": "179977 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 6485,
            "unit": "ns/op",
            "extra": "179977 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7736,
            "unit": "B/op",
            "extra": "179977 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 80,
            "unit": "allocs/op",
            "extra": "179977 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter)",
            "value": 490499,
            "unit": "ns/op\t  727323 B/op\t    5180 allocs/op",
            "extra": "2593 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 490499,
            "unit": "ns/op",
            "extra": "2593 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 727323,
            "unit": "B/op",
            "extra": "2593 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 5180,
            "unit": "allocs/op",
            "extra": "2593 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter)",
            "value": 44489158,
            "unit": "ns/op\t67574894 B/op\t  501554 allocs/op",
            "extra": "31 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 44489158,
            "unit": "ns/op",
            "extra": "31 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 67574894,
            "unit": "B/op",
            "extra": "31 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 501554,
            "unit": "allocs/op",
            "extra": "31 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter)",
            "value": 6984,
            "unit": "ns/op\t    7888 B/op\t      87 allocs/op",
            "extra": "170179 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 6984,
            "unit": "ns/op",
            "extra": "170179 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7888,
            "unit": "B/op",
            "extra": "170179 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "170179 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter)",
            "value": 64306,
            "unit": "ns/op\t   78368 B/op\t     750 allocs/op",
            "extra": "18668 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 64306,
            "unit": "ns/op",
            "extra": "18668 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 78368,
            "unit": "B/op",
            "extra": "18668 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 750,
            "unit": "allocs/op",
            "extra": "18668 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter)",
            "value": 649311,
            "unit": "ns/op\t  768980 B/op\t    7285 allocs/op",
            "extra": "1885 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 649311,
            "unit": "ns/op",
            "extra": "1885 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 768980,
            "unit": "B/op",
            "extra": "1885 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 7285,
            "unit": "allocs/op",
            "extra": "1885 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter)",
            "value": 2173,
            "unit": "ns/op\t    1344 B/op\t      33 allocs/op",
            "extra": "546832 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - ns/op",
            "value": 2173,
            "unit": "ns/op",
            "extra": "546832 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - B/op",
            "value": 1344,
            "unit": "B/op",
            "extra": "546832 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "546832 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 149.6,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7125957 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 149.6,
            "unit": "ns/op",
            "extra": "7125957 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7125957 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7125957 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 166.8,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "6576770 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 166.8,
            "unit": "ns/op",
            "extra": "6576770 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "6576770 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "6576770 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 172.3,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7376023 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 172.3,
            "unit": "ns/op",
            "extra": "7376023 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7376023 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7376023 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 7.342,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "163134578 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 7.342,
            "unit": "ns/op",
            "extra": "163134578 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "163134578 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "163134578 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 7.332,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "164048919 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 7.332,
            "unit": "ns/op",
            "extra": "164048919 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "164048919 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "164048919 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal)",
            "value": 482650,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2474 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - ns/op",
            "value": 482650,
            "unit": "ns/op",
            "extra": "2474 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2474 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2474 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal)",
            "value": 186480,
            "unit": "ns/op\t  295553 B/op\t      33 allocs/op",
            "extra": "5750 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - ns/op",
            "value": 186480,
            "unit": "ns/op",
            "extra": "5750 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - B/op",
            "value": 295553,
            "unit": "B/op",
            "extra": "5750 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "5750 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.505,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "162674967 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.505,
            "unit": "ns/op",
            "extra": "162674967 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "162674967 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "162674967 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.666,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "159274117 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.666,
            "unit": "ns/op",
            "extra": "159274117 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "159274117 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "159274117 times\n2 procs"
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
          "id": "32d5a0b784d949fb3eaf81b24079580d9cde2e9e",
          "message": "release process",
          "timestamp": "2026-05-02T00:33:21+02:00",
          "tree_id": "868b78d473e1526191292c0744a25d7ed1cb7b70",
          "url": "https://github.com/MichaelMure/benchspotter/commit/32d5a0b784d949fb3eaf81b24079580d9cde2e9e"
        },
        "date": 1777674893440,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter)",
            "value": 1913,
            "unit": "ns/op\t    2144 B/op\t      31 allocs/op",
            "extra": "601598 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 1913,
            "unit": "ns/op",
            "extra": "601598 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 2144,
            "unit": "B/op",
            "extra": "601598 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "601598 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter)",
            "value": 1053,
            "unit": "ns/op\t     640 B/op\t      10 allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 1053,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter)",
            "value": 141901,
            "unit": "ns/op\t  174944 B/op\t    2024 allocs/op",
            "extra": "8533 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 141901,
            "unit": "ns/op",
            "extra": "8533 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 174944,
            "unit": "B/op",
            "extra": "8533 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 2024,
            "unit": "allocs/op",
            "extra": "8533 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 101268,
            "unit": "ns/op\t   64011 B/op\t    1000 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 101268,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64011,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter)",
            "value": 31816185,
            "unit": "ns/op\t24685051 B/op\t  200056 allocs/op",
            "extra": "49 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 31816185,
            "unit": "ns/op",
            "extra": "49 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 24685051,
            "unit": "B/op",
            "extra": "49 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 200056,
            "unit": "allocs/op",
            "extra": "49 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 12095226,
            "unit": "ns/op\t 6588512 B/op\t  101031 allocs/op",
            "extra": "97 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 12095226,
            "unit": "ns/op",
            "extra": "97 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6588512,
            "unit": "B/op",
            "extra": "97 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101031,
            "unit": "allocs/op",
            "extra": "97 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter)",
            "value": 10198,
            "unit": "ns/op\t   11440 B/op\t     131 allocs/op",
            "extra": "120428 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 10198,
            "unit": "ns/op",
            "extra": "120428 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 11440,
            "unit": "B/op",
            "extra": "120428 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "120428 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter)",
            "value": 7088,
            "unit": "ns/op\t    6400 B/op\t     100 allocs/op",
            "extra": "167485 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 7088,
            "unit": "ns/op",
            "extra": "167485 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6400,
            "unit": "B/op",
            "extra": "167485 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "167485 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter)",
            "value": 819580,
            "unit": "ns/op\t 1035123 B/op\t   11038 allocs/op",
            "extra": "1378 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 819580,
            "unit": "ns/op",
            "extra": "1378 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 1035123,
            "unit": "B/op",
            "extra": "1378 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 11038,
            "unit": "allocs/op",
            "extra": "1378 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 715874,
            "unit": "ns/op\t  640245 B/op\t   10000 allocs/op",
            "extra": "1646 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 715874,
            "unit": "ns/op",
            "extra": "1646 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640245,
            "unit": "B/op",
            "extra": "1646 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10000,
            "unit": "allocs/op",
            "extra": "1646 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter)",
            "value": 104125354,
            "unit": "ns/op\t111012770 B/op\t 1100072 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 104125354,
            "unit": "ns/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 111012770,
            "unit": "B/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1100072,
            "unit": "allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 79285101,
            "unit": "ns/op\t67358052 B/op\t 1007148 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 79285101,
            "unit": "ns/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 67358052,
            "unit": "B/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1007148,
            "unit": "allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter)",
            "value": 85796,
            "unit": "ns/op\t  103280 B/op\t    1041 allocs/op",
            "extra": "15292 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 85796,
            "unit": "ns/op",
            "extra": "15292 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 103280,
            "unit": "B/op",
            "extra": "15292 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1041,
            "unit": "allocs/op",
            "extra": "15292 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter)",
            "value": 65591,
            "unit": "ns/op\t   64002 B/op\t    1000 allocs/op",
            "extra": "18244 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 65591,
            "unit": "ns/op",
            "extra": "18244 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64002,
            "unit": "B/op",
            "extra": "18244 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "18244 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter)",
            "value": 8681208,
            "unit": "ns/op\t 9670513 B/op\t  101052 allocs/op",
            "extra": "138 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 8681208,
            "unit": "ns/op",
            "extra": "138 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 9670513,
            "unit": "B/op",
            "extra": "138 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101052,
            "unit": "allocs/op",
            "extra": "138 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 7569476,
            "unit": "ns/op\t 6419519 B/op\t  100006 allocs/op",
            "extra": "168 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 7569476,
            "unit": "ns/op",
            "extra": "168 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6419519,
            "unit": "B/op",
            "extra": "168 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100006,
            "unit": "allocs/op",
            "extra": "168 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter)",
            "value": 806727576,
            "unit": "ns/op\t974041984 B/op\t10100088 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 806727576,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 974041984,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10100088,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 788699008,
            "unit": "ns/op\t807020992 B/op\t10050044 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 788699008,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 807020992,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10050044,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter)",
            "value": 6625,
            "unit": "ns/op\t    7736 B/op\t      80 allocs/op",
            "extra": "179431 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 6625,
            "unit": "ns/op",
            "extra": "179431 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7736,
            "unit": "B/op",
            "extra": "179431 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 80,
            "unit": "allocs/op",
            "extra": "179431 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter)",
            "value": 481307,
            "unit": "ns/op\t  727320 B/op\t    5180 allocs/op",
            "extra": "2476 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 481307,
            "unit": "ns/op",
            "extra": "2476 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 727320,
            "unit": "B/op",
            "extra": "2476 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 5180,
            "unit": "allocs/op",
            "extra": "2476 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter)",
            "value": 46258421,
            "unit": "ns/op\t67574809 B/op\t  501554 allocs/op",
            "extra": "28 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 46258421,
            "unit": "ns/op",
            "extra": "28 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 67574809,
            "unit": "B/op",
            "extra": "28 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 501554,
            "unit": "allocs/op",
            "extra": "28 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter)",
            "value": 7515,
            "unit": "ns/op\t    7888 B/op\t      87 allocs/op",
            "extra": "167792 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 7515,
            "unit": "ns/op",
            "extra": "167792 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7888,
            "unit": "B/op",
            "extra": "167792 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "167792 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter)",
            "value": 65770,
            "unit": "ns/op\t   78368 B/op\t     750 allocs/op",
            "extra": "18177 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 65770,
            "unit": "ns/op",
            "extra": "18177 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 78368,
            "unit": "B/op",
            "extra": "18177 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 750,
            "unit": "allocs/op",
            "extra": "18177 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter)",
            "value": 657660,
            "unit": "ns/op\t  768979 B/op\t    7285 allocs/op",
            "extra": "1872 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 657660,
            "unit": "ns/op",
            "extra": "1872 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 768979,
            "unit": "B/op",
            "extra": "1872 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 7285,
            "unit": "allocs/op",
            "extra": "1872 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter)",
            "value": 2215,
            "unit": "ns/op\t    1344 B/op\t      33 allocs/op",
            "extra": "552351 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - ns/op",
            "value": 2215,
            "unit": "ns/op",
            "extra": "552351 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - B/op",
            "value": 1344,
            "unit": "B/op",
            "extra": "552351 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "552351 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 159,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7206514 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 159,
            "unit": "ns/op",
            "extra": "7206514 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7206514 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7206514 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 160.8,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7527549 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 160.8,
            "unit": "ns/op",
            "extra": "7527549 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7527549 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7527549 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 155.1,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "6510079 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 155.1,
            "unit": "ns/op",
            "extra": "6510079 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "6510079 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "6510079 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 7.544,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "159949735 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 7.544,
            "unit": "ns/op",
            "extra": "159949735 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "159949735 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "159949735 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 7.837,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "161683886 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 7.837,
            "unit": "ns/op",
            "extra": "161683886 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "161683886 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "161683886 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal)",
            "value": 483172,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2479 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - ns/op",
            "value": 483172,
            "unit": "ns/op",
            "extra": "2479 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2479 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2479 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal)",
            "value": 195003,
            "unit": "ns/op\t  295552 B/op\t      33 allocs/op",
            "extra": "6187 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - ns/op",
            "value": 195003,
            "unit": "ns/op",
            "extra": "6187 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - B/op",
            "value": 295552,
            "unit": "B/op",
            "extra": "6187 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "6187 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.304,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "163248565 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.304,
            "unit": "ns/op",
            "extra": "163248565 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "163248565 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "163248565 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.344,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "164007184 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.344,
            "unit": "ns/op",
            "extra": "164007184 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "164007184 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "164007184 times\n2 procs"
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
          "id": "391cfe9204756c0f9fe8f51ecccd38cb952f646f",
          "message": "bench: --all option and regex matching",
          "timestamp": "2026-05-02T11:27:17+02:00",
          "tree_id": "2d5401fb133c869f25c7fd272873a831f90d91ad",
          "url": "https://github.com/MichaelMure/benchspotter/commit/391cfe9204756c0f9fe8f51ecccd38cb952f646f"
        },
        "date": 1777714138715,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter)",
            "value": 1822,
            "unit": "ns/op\t    2144 B/op\t      31 allocs/op",
            "extra": "628599 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 1822,
            "unit": "ns/op",
            "extra": "628599 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 2144,
            "unit": "B/op",
            "extra": "628599 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "628599 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter)",
            "value": 1054,
            "unit": "ns/op\t     640 B/op\t      10 allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 1054,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter)",
            "value": 138808,
            "unit": "ns/op\t  174945 B/op\t    2024 allocs/op",
            "extra": "8166 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 138808,
            "unit": "ns/op",
            "extra": "8166 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 174945,
            "unit": "B/op",
            "extra": "8166 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 2024,
            "unit": "allocs/op",
            "extra": "8166 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 100426,
            "unit": "ns/op\t   64011 B/op\t    1000 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 100426,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64011,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter)",
            "value": 31258890,
            "unit": "ns/op\t24685026 B/op\t  200056 allocs/op",
            "extra": "42 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 31258890,
            "unit": "ns/op",
            "extra": "42 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 24685026,
            "unit": "B/op",
            "extra": "42 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 200056,
            "unit": "allocs/op",
            "extra": "42 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 12028032,
            "unit": "ns/op\t 6607784 B/op\t  101137 allocs/op",
            "extra": "88 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 12028032,
            "unit": "ns/op",
            "extra": "88 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6607784,
            "unit": "B/op",
            "extra": "88 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101137,
            "unit": "allocs/op",
            "extra": "88 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter)",
            "value": 8802,
            "unit": "ns/op\t   11440 B/op\t     131 allocs/op",
            "extra": "131678 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 8802,
            "unit": "ns/op",
            "extra": "131678 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 11440,
            "unit": "B/op",
            "extra": "131678 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "131678 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter)",
            "value": 7372,
            "unit": "ns/op\t    6400 B/op\t     100 allocs/op",
            "extra": "176120 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 7372,
            "unit": "ns/op",
            "extra": "176120 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6400,
            "unit": "B/op",
            "extra": "176120 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "176120 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter)",
            "value": 790127,
            "unit": "ns/op\t 1035129 B/op\t   11038 allocs/op",
            "extra": "1530 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 790127,
            "unit": "ns/op",
            "extra": "1530 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 1035129,
            "unit": "B/op",
            "extra": "1530 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 11038,
            "unit": "allocs/op",
            "extra": "1530 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 688630,
            "unit": "ns/op\t  640237 B/op\t   10000 allocs/op",
            "extra": "1710 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 688630,
            "unit": "ns/op",
            "extra": "1710 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640237,
            "unit": "B/op",
            "extra": "1710 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10000,
            "unit": "allocs/op",
            "extra": "1710 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter)",
            "value": 106953816,
            "unit": "ns/op\t111012752 B/op\t 1100072 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 106953816,
            "unit": "ns/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 111012752,
            "unit": "B/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1100072,
            "unit": "allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 80234687,
            "unit": "ns/op\t67358076 B/op\t 1007148 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 80234687,
            "unit": "ns/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 67358076,
            "unit": "B/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1007148,
            "unit": "allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter)",
            "value": 75803,
            "unit": "ns/op\t  103281 B/op\t    1041 allocs/op",
            "extra": "15868 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 75803,
            "unit": "ns/op",
            "extra": "15868 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 103281,
            "unit": "B/op",
            "extra": "15868 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1041,
            "unit": "allocs/op",
            "extra": "15868 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter)",
            "value": 65102,
            "unit": "ns/op\t   64002 B/op\t    1000 allocs/op",
            "extra": "18246 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 65102,
            "unit": "ns/op",
            "extra": "18246 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64002,
            "unit": "B/op",
            "extra": "18246 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "18246 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter)",
            "value": 8557448,
            "unit": "ns/op\t 9670561 B/op\t  101052 allocs/op",
            "extra": "140 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 8557448,
            "unit": "ns/op",
            "extra": "140 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 9670561,
            "unit": "B/op",
            "extra": "140 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101052,
            "unit": "allocs/op",
            "extra": "140 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 6889848,
            "unit": "ns/op\t 6419401 B/op\t  100006 allocs/op",
            "extra": "169 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 6889848,
            "unit": "ns/op",
            "extra": "169 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6419401,
            "unit": "B/op",
            "extra": "169 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100006,
            "unit": "allocs/op",
            "extra": "169 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter)",
            "value": 839426872,
            "unit": "ns/op\t974042096 B/op\t10100089 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 839426872,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 974042096,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10100089,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 805282850,
            "unit": "ns/op\t807020992 B/op\t10050044 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 805282850,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 807020992,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10050044,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter)",
            "value": 6404,
            "unit": "ns/op\t    7736 B/op\t      80 allocs/op",
            "extra": "181273 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 6404,
            "unit": "ns/op",
            "extra": "181273 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7736,
            "unit": "B/op",
            "extra": "181273 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 80,
            "unit": "allocs/op",
            "extra": "181273 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter)",
            "value": 507831,
            "unit": "ns/op\t  727321 B/op\t    5180 allocs/op",
            "extra": "2436 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 507831,
            "unit": "ns/op",
            "extra": "2436 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 727321,
            "unit": "B/op",
            "extra": "2436 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 5180,
            "unit": "allocs/op",
            "extra": "2436 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter)",
            "value": 46494091,
            "unit": "ns/op\t67574840 B/op\t  501554 allocs/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 46494091,
            "unit": "ns/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 67574840,
            "unit": "B/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 501554,
            "unit": "allocs/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter)",
            "value": 7289,
            "unit": "ns/op\t    7888 B/op\t      87 allocs/op",
            "extra": "146158 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 7289,
            "unit": "ns/op",
            "extra": "146158 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7888,
            "unit": "B/op",
            "extra": "146158 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "146158 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter)",
            "value": 67565,
            "unit": "ns/op\t   78368 B/op\t     750 allocs/op",
            "extra": "18285 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 67565,
            "unit": "ns/op",
            "extra": "18285 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 78368,
            "unit": "B/op",
            "extra": "18285 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 750,
            "unit": "allocs/op",
            "extra": "18285 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter)",
            "value": 633844,
            "unit": "ns/op\t  768979 B/op\t    7285 allocs/op",
            "extra": "1964 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 633844,
            "unit": "ns/op",
            "extra": "1964 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 768979,
            "unit": "B/op",
            "extra": "1964 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 7285,
            "unit": "allocs/op",
            "extra": "1964 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter)",
            "value": 2013,
            "unit": "ns/op\t    1344 B/op\t      33 allocs/op",
            "extra": "549429 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - ns/op",
            "value": 2013,
            "unit": "ns/op",
            "extra": "549429 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - B/op",
            "value": 1344,
            "unit": "B/op",
            "extra": "549429 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "549429 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 155.7,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "8114893 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 155.7,
            "unit": "ns/op",
            "extra": "8114893 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "8114893 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8114893 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 151.2,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "8215852 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 151.2,
            "unit": "ns/op",
            "extra": "8215852 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "8215852 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8215852 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 152.7,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "8135031 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 152.7,
            "unit": "ns/op",
            "extra": "8135031 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "8135031 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8135031 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 7.737,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "158557503 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 7.737,
            "unit": "ns/op",
            "extra": "158557503 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "158557503 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "158557503 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 7.571,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "158463256 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 7.571,
            "unit": "ns/op",
            "extra": "158463256 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "158463256 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "158463256 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal)",
            "value": 539317,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2220 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - ns/op",
            "value": 539317,
            "unit": "ns/op",
            "extra": "2220 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2220 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2220 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal)",
            "value": 158872,
            "unit": "ns/op\t  295554 B/op\t      33 allocs/op",
            "extra": "7092 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - ns/op",
            "value": 158872,
            "unit": "ns/op",
            "extra": "7092 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - B/op",
            "value": 295554,
            "unit": "B/op",
            "extra": "7092 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "7092 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.787,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "157385431 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.787,
            "unit": "ns/op",
            "extra": "157385431 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "157385431 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "157385431 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.56,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "153954970 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.56,
            "unit": "ns/op",
            "extra": "153954970 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "153954970 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "153954970 times\n2 procs"
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
          "id": "f62bf251b2c312a37ee297685987350aedf8c5ab",
          "message": "trend: page up/down navigation",
          "timestamp": "2026-05-02T13:45:45+02:00",
          "tree_id": "45787c3eed423480b87d3d575f9bfa5bffbe68e9",
          "url": "https://github.com/MichaelMure/benchspotter/commit/f62bf251b2c312a37ee297685987350aedf8c5ab"
        },
        "date": 1777722438869,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter)",
            "value": 2059,
            "unit": "ns/op\t    2144 B/op\t      31 allocs/op",
            "extra": "539858 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 2059,
            "unit": "ns/op",
            "extra": "539858 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 2144,
            "unit": "B/op",
            "extra": "539858 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "539858 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter)",
            "value": 1005,
            "unit": "ns/op\t     640 B/op\t      10 allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 1005,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter)",
            "value": 133147,
            "unit": "ns/op\t  174944 B/op\t    2024 allocs/op",
            "extra": "8332 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 133147,
            "unit": "ns/op",
            "extra": "8332 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 174944,
            "unit": "B/op",
            "extra": "8332 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 2024,
            "unit": "allocs/op",
            "extra": "8332 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 93539,
            "unit": "ns/op\t   64009 B/op\t    1000 allocs/op",
            "extra": "12798 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 93539,
            "unit": "ns/op",
            "extra": "12798 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64009,
            "unit": "B/op",
            "extra": "12798 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "12798 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter)",
            "value": 27158488,
            "unit": "ns/op\t24685040 B/op\t  200056 allocs/op",
            "extra": "54 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 27158488,
            "unit": "ns/op",
            "extra": "54 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 24685040,
            "unit": "B/op",
            "extra": "54 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 200056,
            "unit": "allocs/op",
            "extra": "54 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 11980644,
            "unit": "ns/op\t 6598771 B/op\t  101087 allocs/op",
            "extra": "92 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 11980644,
            "unit": "ns/op",
            "extra": "92 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6598771,
            "unit": "B/op",
            "extra": "92 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101087,
            "unit": "allocs/op",
            "extra": "92 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter)",
            "value": 8856,
            "unit": "ns/op\t   11440 B/op\t     131 allocs/op",
            "extra": "134778 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 8856,
            "unit": "ns/op",
            "extra": "134778 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 11440,
            "unit": "B/op",
            "extra": "134778 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "134778 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter)",
            "value": 6594,
            "unit": "ns/op\t    6400 B/op\t     100 allocs/op",
            "extra": "181226 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 6594,
            "unit": "ns/op",
            "extra": "181226 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6400,
            "unit": "B/op",
            "extra": "181226 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "181226 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter)",
            "value": 782167,
            "unit": "ns/op\t 1035120 B/op\t   11038 allocs/op",
            "extra": "1542 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 782167,
            "unit": "ns/op",
            "extra": "1542 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 1035120,
            "unit": "B/op",
            "extra": "1542 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 11038,
            "unit": "allocs/op",
            "extra": "1542 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 643692,
            "unit": "ns/op\t  640229 B/op\t   10000 allocs/op",
            "extra": "1718 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 643692,
            "unit": "ns/op",
            "extra": "1718 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640229,
            "unit": "B/op",
            "extra": "1718 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10000,
            "unit": "allocs/op",
            "extra": "1718 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter)",
            "value": 101834456,
            "unit": "ns/op\t111012729 B/op\t 1100072 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 101834456,
            "unit": "ns/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 111012729,
            "unit": "B/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1100072,
            "unit": "allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 78486226,
            "unit": "ns/op\t67358052 B/op\t 1007148 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 78486226,
            "unit": "ns/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 67358052,
            "unit": "B/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1007148,
            "unit": "allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter)",
            "value": 68763,
            "unit": "ns/op\t  103280 B/op\t    1041 allocs/op",
            "extra": "17511 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 68763,
            "unit": "ns/op",
            "extra": "17511 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 103280,
            "unit": "B/op",
            "extra": "17511 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1041,
            "unit": "allocs/op",
            "extra": "17511 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter)",
            "value": 59134,
            "unit": "ns/op\t   64002 B/op\t    1000 allocs/op",
            "extra": "20178 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 59134,
            "unit": "ns/op",
            "extra": "20178 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64002,
            "unit": "B/op",
            "extra": "20178 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "20178 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter)",
            "value": 7713304,
            "unit": "ns/op\t 9670512 B/op\t  101052 allocs/op",
            "extra": "154 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 7713304,
            "unit": "ns/op",
            "extra": "154 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 9670512,
            "unit": "B/op",
            "extra": "154 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101052,
            "unit": "allocs/op",
            "extra": "154 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 6344010,
            "unit": "ns/op\t 6417396 B/op\t  100005 allocs/op",
            "extra": "188 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 6344010,
            "unit": "ns/op",
            "extra": "188 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6417396,
            "unit": "B/op",
            "extra": "188 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100005,
            "unit": "allocs/op",
            "extra": "188 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter)",
            "value": 821483930,
            "unit": "ns/op\t974041984 B/op\t10100088 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 821483930,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 974041984,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10100088,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 796857853,
            "unit": "ns/op\t807020992 B/op\t10050044 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 796857853,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 807020992,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10050044,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter)",
            "value": 6004,
            "unit": "ns/op\t    7736 B/op\t      80 allocs/op",
            "extra": "193045 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 6004,
            "unit": "ns/op",
            "extra": "193045 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7736,
            "unit": "B/op",
            "extra": "193045 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 80,
            "unit": "allocs/op",
            "extra": "193045 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter)",
            "value": 414201,
            "unit": "ns/op\t  727322 B/op\t    5180 allocs/op",
            "extra": "2852 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 414201,
            "unit": "ns/op",
            "extra": "2852 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 727322,
            "unit": "B/op",
            "extra": "2852 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 5180,
            "unit": "allocs/op",
            "extra": "2852 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter)",
            "value": 42713219,
            "unit": "ns/op\t67574811 B/op\t  501554 allocs/op",
            "extra": "25 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 42713219,
            "unit": "ns/op",
            "extra": "25 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 67574811,
            "unit": "B/op",
            "extra": "25 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 501554,
            "unit": "allocs/op",
            "extra": "25 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter)",
            "value": 6410,
            "unit": "ns/op\t    7888 B/op\t      87 allocs/op",
            "extra": "180889 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 6410,
            "unit": "ns/op",
            "extra": "180889 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7888,
            "unit": "B/op",
            "extra": "180889 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "180889 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter)",
            "value": 59132,
            "unit": "ns/op\t   78368 B/op\t     750 allocs/op",
            "extra": "20422 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 59132,
            "unit": "ns/op",
            "extra": "20422 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 78368,
            "unit": "B/op",
            "extra": "20422 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 750,
            "unit": "allocs/op",
            "extra": "20422 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter)",
            "value": 588027,
            "unit": "ns/op\t  768976 B/op\t    7285 allocs/op",
            "extra": "2001 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 588027,
            "unit": "ns/op",
            "extra": "2001 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 768976,
            "unit": "B/op",
            "extra": "2001 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 7285,
            "unit": "allocs/op",
            "extra": "2001 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter)",
            "value": 1889,
            "unit": "ns/op\t    1344 B/op\t      33 allocs/op",
            "extra": "593523 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - ns/op",
            "value": 1889,
            "unit": "ns/op",
            "extra": "593523 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - B/op",
            "value": 1344,
            "unit": "B/op",
            "extra": "593523 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "593523 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 159.5,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7269733 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 159.5,
            "unit": "ns/op",
            "extra": "7269733 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7269733 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7269733 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 156.9,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7602949 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 156.9,
            "unit": "ns/op",
            "extra": "7602949 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7602949 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7602949 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 161.6,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7318719 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 161.6,
            "unit": "ns/op",
            "extra": "7318719 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7318719 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7318719 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 6.627,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "181417930 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 6.627,
            "unit": "ns/op",
            "extra": "181417930 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "181417930 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "181417930 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 6.334,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "189482204 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 6.334,
            "unit": "ns/op",
            "extra": "189482204 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "189482204 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "189482204 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal)",
            "value": 518490,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2395 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - ns/op",
            "value": 518490,
            "unit": "ns/op",
            "extra": "2395 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2395 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2395 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal)",
            "value": 125308,
            "unit": "ns/op\t  295552 B/op\t      33 allocs/op",
            "extra": "8553 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - ns/op",
            "value": 125308,
            "unit": "ns/op",
            "extra": "8553 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - B/op",
            "value": 295552,
            "unit": "B/op",
            "extra": "8553 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "8553 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 6.895,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "171488437 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 6.895,
            "unit": "ns/op",
            "extra": "171488437 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171488437 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171488437 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 6.897,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "190126186 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 6.897,
            "unit": "ns/op",
            "extra": "190126186 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "190126186 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "190126186 times\n2 procs"
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
          "id": "382dd02ced1cb0952c85820b40a3ca24839a75c8",
          "message": "--no-prompt flag",
          "timestamp": "2026-05-03T01:04:22+02:00",
          "tree_id": "468627fc090784b1631a503458b790b97de7c057",
          "url": "https://github.com/MichaelMure/benchspotter/commit/382dd02ced1cb0952c85820b40a3ca24839a75c8"
        },
        "date": 1777763149887,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter)",
            "value": 1906,
            "unit": "ns/op\t    2144 B/op\t      31 allocs/op",
            "extra": "623276 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 1906,
            "unit": "ns/op",
            "extra": "623276 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 2144,
            "unit": "B/op",
            "extra": "623276 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "623276 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter)",
            "value": 1039,
            "unit": "ns/op\t     640 B/op\t      10 allocs/op",
            "extra": "994719 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 1039,
            "unit": "ns/op",
            "extra": "994719 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "994719 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "994719 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter)",
            "value": 155321,
            "unit": "ns/op\t  174944 B/op\t    2024 allocs/op",
            "extra": "7215 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 155321,
            "unit": "ns/op",
            "extra": "7215 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 174944,
            "unit": "B/op",
            "extra": "7215 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 2024,
            "unit": "allocs/op",
            "extra": "7215 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 106836,
            "unit": "ns/op\t   64011 B/op\t    1000 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 106836,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64011,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter)",
            "value": 29075998,
            "unit": "ns/op\t24685072 B/op\t  200056 allocs/op",
            "extra": "45 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 29075998,
            "unit": "ns/op",
            "extra": "45 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 24685072,
            "unit": "B/op",
            "extra": "45 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 200056,
            "unit": "allocs/op",
            "extra": "45 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 11769780,
            "unit": "ns/op\t 6584697 B/op\t  101010 allocs/op",
            "extra": "99 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 11769780,
            "unit": "ns/op",
            "extra": "99 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6584697,
            "unit": "B/op",
            "extra": "99 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101010,
            "unit": "allocs/op",
            "extra": "99 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter)",
            "value": 9205,
            "unit": "ns/op\t   11440 B/op\t     131 allocs/op",
            "extra": "127981 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 9205,
            "unit": "ns/op",
            "extra": "127981 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 11440,
            "unit": "B/op",
            "extra": "127981 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "127981 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter)",
            "value": 7096,
            "unit": "ns/op\t    6400 B/op\t     100 allocs/op",
            "extra": "170174 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 7096,
            "unit": "ns/op",
            "extra": "170174 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6400,
            "unit": "B/op",
            "extra": "170174 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "170174 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter)",
            "value": 817989,
            "unit": "ns/op\t 1035126 B/op\t   11038 allocs/op",
            "extra": "1471 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 817989,
            "unit": "ns/op",
            "extra": "1471 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 1035126,
            "unit": "B/op",
            "extra": "1471 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 11038,
            "unit": "allocs/op",
            "extra": "1471 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 718259,
            "unit": "ns/op\t  640239 B/op\t   10000 allocs/op",
            "extra": "1671 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 718259,
            "unit": "ns/op",
            "extra": "1671 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640239,
            "unit": "B/op",
            "extra": "1671 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10000,
            "unit": "allocs/op",
            "extra": "1671 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter)",
            "value": 106584035,
            "unit": "ns/op\t111012728 B/op\t 1100072 allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 106584035,
            "unit": "ns/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 111012728,
            "unit": "B/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1100072,
            "unit": "allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 81968573,
            "unit": "ns/op\t67358052 B/op\t 1007148 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 81968573,
            "unit": "ns/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 67358052,
            "unit": "B/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1007148,
            "unit": "allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter)",
            "value": 75962,
            "unit": "ns/op\t  103280 B/op\t    1041 allocs/op",
            "extra": "15506 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 75962,
            "unit": "ns/op",
            "extra": "15506 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 103280,
            "unit": "B/op",
            "extra": "15506 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1041,
            "unit": "allocs/op",
            "extra": "15506 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter)",
            "value": 65129,
            "unit": "ns/op\t   64002 B/op\t    1000 allocs/op",
            "extra": "18127 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 65129,
            "unit": "ns/op",
            "extra": "18127 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64002,
            "unit": "B/op",
            "extra": "18127 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "18127 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter)",
            "value": 8710524,
            "unit": "ns/op\t 9670530 B/op\t  101052 allocs/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 8710524,
            "unit": "ns/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 9670530,
            "unit": "B/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101052,
            "unit": "allocs/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 6916842,
            "unit": "ns/op\t 6419269 B/op\t  100006 allocs/op",
            "extra": "170 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 6916842,
            "unit": "ns/op",
            "extra": "170 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6419269,
            "unit": "B/op",
            "extra": "170 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100006,
            "unit": "allocs/op",
            "extra": "170 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter)",
            "value": 842429892,
            "unit": "ns/op\t974041976 B/op\t10100088 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 842429892,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 974041976,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10100088,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 776428708,
            "unit": "ns/op\t807020992 B/op\t10050044 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 776428708,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 807020992,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10050044,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter)",
            "value": 6386,
            "unit": "ns/op\t    7736 B/op\t      80 allocs/op",
            "extra": "181371 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 6386,
            "unit": "ns/op",
            "extra": "181371 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7736,
            "unit": "B/op",
            "extra": "181371 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 80,
            "unit": "allocs/op",
            "extra": "181371 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter)",
            "value": 460217,
            "unit": "ns/op\t  727321 B/op\t    5180 allocs/op",
            "extra": "2593 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 460217,
            "unit": "ns/op",
            "extra": "2593 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 727321,
            "unit": "B/op",
            "extra": "2593 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 5180,
            "unit": "allocs/op",
            "extra": "2593 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter)",
            "value": 45477690,
            "unit": "ns/op\t67574810 B/op\t  501554 allocs/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 45477690,
            "unit": "ns/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 67574810,
            "unit": "B/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 501554,
            "unit": "allocs/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter)",
            "value": 6912,
            "unit": "ns/op\t    7888 B/op\t      87 allocs/op",
            "extra": "169512 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 6912,
            "unit": "ns/op",
            "extra": "169512 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7888,
            "unit": "B/op",
            "extra": "169512 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "169512 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter)",
            "value": 64802,
            "unit": "ns/op\t   78368 B/op\t     750 allocs/op",
            "extra": "18652 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 64802,
            "unit": "ns/op",
            "extra": "18652 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 78368,
            "unit": "B/op",
            "extra": "18652 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 750,
            "unit": "allocs/op",
            "extra": "18652 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter)",
            "value": 627231,
            "unit": "ns/op\t  768977 B/op\t    7285 allocs/op",
            "extra": "1948 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 627231,
            "unit": "ns/op",
            "extra": "1948 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 768977,
            "unit": "B/op",
            "extra": "1948 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 7285,
            "unit": "allocs/op",
            "extra": "1948 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter)",
            "value": 2084,
            "unit": "ns/op\t    1344 B/op\t      33 allocs/op",
            "extra": "557108 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - ns/op",
            "value": 2084,
            "unit": "ns/op",
            "extra": "557108 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - B/op",
            "value": 1344,
            "unit": "B/op",
            "extra": "557108 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "557108 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 165.6,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "6722910 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 165.6,
            "unit": "ns/op",
            "extra": "6722910 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "6722910 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "6722910 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 166.8,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7229208 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 166.8,
            "unit": "ns/op",
            "extra": "7229208 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7229208 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7229208 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 166.8,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "6683328 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 166.8,
            "unit": "ns/op",
            "extra": "6683328 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "6683328 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "6683328 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 7.374,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "159877758 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 7.374,
            "unit": "ns/op",
            "extra": "159877758 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "159877758 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "159877758 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 7.328,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "163791760 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 7.328,
            "unit": "ns/op",
            "extra": "163791760 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "163791760 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "163791760 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal)",
            "value": 482914,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2472 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - ns/op",
            "value": 482914,
            "unit": "ns/op",
            "extra": "2472 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2472 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2472 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal)",
            "value": 192956,
            "unit": "ns/op\t  295552 B/op\t      33 allocs/op",
            "extra": "5800 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - ns/op",
            "value": 192956,
            "unit": "ns/op",
            "extra": "5800 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - B/op",
            "value": 295552,
            "unit": "B/op",
            "extra": "5800 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "5800 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.297,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "163592502 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.297,
            "unit": "ns/op",
            "extra": "163592502 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "163592502 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "163592502 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.461,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "160969582 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.461,
            "unit": "ns/op",
            "extra": "160969582 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "160969582 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "160969582 times\n2 procs"
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
          "id": "09401e461acd5815ce8ffabcf8bdaabfd4b8f73c",
          "message": "--no-prompt flag",
          "timestamp": "2026-05-03T01:05:53+02:00",
          "tree_id": "18883b6b6f32ecb875375e758bd6109d84319908",
          "url": "https://github.com/MichaelMure/benchspotter/commit/09401e461acd5815ce8ffabcf8bdaabfd4b8f73c"
        },
        "date": 1777763242955,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter)",
            "value": 1934,
            "unit": "ns/op\t    2144 B/op\t      31 allocs/op",
            "extra": "589140 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 1934,
            "unit": "ns/op",
            "extra": "589140 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 2144,
            "unit": "B/op",
            "extra": "589140 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "589140 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter)",
            "value": 1044,
            "unit": "ns/op\t     640 B/op\t      10 allocs/op",
            "extra": "961174 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 1044,
            "unit": "ns/op",
            "extra": "961174 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "961174 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "961174 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter)",
            "value": 139814,
            "unit": "ns/op\t  174944 B/op\t    2024 allocs/op",
            "extra": "7658 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 139814,
            "unit": "ns/op",
            "extra": "7658 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 174944,
            "unit": "B/op",
            "extra": "7658 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 2024,
            "unit": "allocs/op",
            "extra": "7658 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 100940,
            "unit": "ns/op\t   64011 B/op\t    1000 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 100940,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64011,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter)",
            "value": 29028228,
            "unit": "ns/op\t24685027 B/op\t  200056 allocs/op",
            "extra": "40 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 29028228,
            "unit": "ns/op",
            "extra": "40 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 24685027,
            "unit": "B/op",
            "extra": "40 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 200056,
            "unit": "allocs/op",
            "extra": "40 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 11733171,
            "unit": "ns/op\t 6584714 B/op\t  101010 allocs/op",
            "extra": "99 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 11733171,
            "unit": "ns/op",
            "extra": "99 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6584714,
            "unit": "B/op",
            "extra": "99 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101010,
            "unit": "allocs/op",
            "extra": "99 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter)",
            "value": 9686,
            "unit": "ns/op\t   11440 B/op\t     131 allocs/op",
            "extra": "132939 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 9686,
            "unit": "ns/op",
            "extra": "132939 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 11440,
            "unit": "B/op",
            "extra": "132939 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "132939 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter)",
            "value": 7049,
            "unit": "ns/op\t    6400 B/op\t     100 allocs/op",
            "extra": "163514 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 7049,
            "unit": "ns/op",
            "extra": "163514 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6400,
            "unit": "B/op",
            "extra": "163514 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "163514 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter)",
            "value": 815736,
            "unit": "ns/op\t 1035128 B/op\t   11038 allocs/op",
            "extra": "1395 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 815736,
            "unit": "ns/op",
            "extra": "1395 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 1035128,
            "unit": "B/op",
            "extra": "1395 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 11038,
            "unit": "allocs/op",
            "extra": "1395 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 730455,
            "unit": "ns/op\t  640250 B/op\t   10000 allocs/op",
            "extra": "1610 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 730455,
            "unit": "ns/op",
            "extra": "1610 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640250,
            "unit": "B/op",
            "extra": "1610 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10000,
            "unit": "allocs/op",
            "extra": "1610 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter)",
            "value": 103761960,
            "unit": "ns/op\t111012736 B/op\t 1100072 allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 103761960,
            "unit": "ns/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 111012736,
            "unit": "B/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1100072,
            "unit": "allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 78457904,
            "unit": "ns/op\t67358052 B/op\t 1007148 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 78457904,
            "unit": "ns/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 67358052,
            "unit": "B/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1007148,
            "unit": "allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter)",
            "value": 80904,
            "unit": "ns/op\t  103280 B/op\t    1041 allocs/op",
            "extra": "15673 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 80904,
            "unit": "ns/op",
            "extra": "15673 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 103280,
            "unit": "B/op",
            "extra": "15673 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1041,
            "unit": "allocs/op",
            "extra": "15673 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter)",
            "value": 66888,
            "unit": "ns/op\t   64002 B/op\t    1000 allocs/op",
            "extra": "18134 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 66888,
            "unit": "ns/op",
            "extra": "18134 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64002,
            "unit": "B/op",
            "extra": "18134 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "18134 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter)",
            "value": 8757030,
            "unit": "ns/op\t 9670548 B/op\t  101052 allocs/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 8757030,
            "unit": "ns/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 9670548,
            "unit": "B/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101052,
            "unit": "allocs/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 7124048,
            "unit": "ns/op\t 6419390 B/op\t  100006 allocs/op",
            "extra": "169 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 7124048,
            "unit": "ns/op",
            "extra": "169 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6419390,
            "unit": "B/op",
            "extra": "169 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100006,
            "unit": "allocs/op",
            "extra": "169 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter)",
            "value": 775737081,
            "unit": "ns/op\t974042040 B/op\t10100088 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 775737081,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 974042040,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10100088,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 758168512,
            "unit": "ns/op\t807020992 B/op\t10050044 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 758168512,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 807020992,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10050044,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter)",
            "value": 5698,
            "unit": "ns/op\t    7736 B/op\t      80 allocs/op",
            "extra": "182887 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 5698,
            "unit": "ns/op",
            "extra": "182887 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7736,
            "unit": "B/op",
            "extra": "182887 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 80,
            "unit": "allocs/op",
            "extra": "182887 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter)",
            "value": 434968,
            "unit": "ns/op\t  727324 B/op\t    5180 allocs/op",
            "extra": "2360 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 434968,
            "unit": "ns/op",
            "extra": "2360 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 727324,
            "unit": "B/op",
            "extra": "2360 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 5180,
            "unit": "allocs/op",
            "extra": "2360 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter)",
            "value": 44030435,
            "unit": "ns/op\t67574886 B/op\t  501554 allocs/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 44030435,
            "unit": "ns/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 67574886,
            "unit": "B/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 501554,
            "unit": "allocs/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter)",
            "value": 6286,
            "unit": "ns/op\t    7888 B/op\t      87 allocs/op",
            "extra": "179796 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 6286,
            "unit": "ns/op",
            "extra": "179796 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7888,
            "unit": "B/op",
            "extra": "179796 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "179796 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter)",
            "value": 60263,
            "unit": "ns/op\t   78368 B/op\t     750 allocs/op",
            "extra": "20252 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 60263,
            "unit": "ns/op",
            "extra": "20252 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 78368,
            "unit": "B/op",
            "extra": "20252 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 750,
            "unit": "allocs/op",
            "extra": "20252 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter)",
            "value": 596301,
            "unit": "ns/op\t  768984 B/op\t    7285 allocs/op",
            "extra": "1970 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 596301,
            "unit": "ns/op",
            "extra": "1970 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 768984,
            "unit": "B/op",
            "extra": "1970 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 7285,
            "unit": "allocs/op",
            "extra": "1970 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter)",
            "value": 2018,
            "unit": "ns/op\t    1344 B/op\t      33 allocs/op",
            "extra": "541562 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - ns/op",
            "value": 2018,
            "unit": "ns/op",
            "extra": "541562 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - B/op",
            "value": 1344,
            "unit": "B/op",
            "extra": "541562 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "541562 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 150.6,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7114810 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 150.6,
            "unit": "ns/op",
            "extra": "7114810 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7114810 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7114810 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 169.2,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7104219 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 169.2,
            "unit": "ns/op",
            "extra": "7104219 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7104219 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7104219 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 158,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7719729 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 158,
            "unit": "ns/op",
            "extra": "7719729 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7719729 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7719729 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 7.363,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "162698356 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 7.363,
            "unit": "ns/op",
            "extra": "162698356 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "162698356 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "162698356 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 7.332,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "163805574 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 7.332,
            "unit": "ns/op",
            "extra": "163805574 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "163805574 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "163805574 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal)",
            "value": 482715,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2473 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - ns/op",
            "value": 482715,
            "unit": "ns/op",
            "extra": "2473 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2473 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2473 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal)",
            "value": 184496,
            "unit": "ns/op\t  295552 B/op\t      33 allocs/op",
            "extra": "6300 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - ns/op",
            "value": 184496,
            "unit": "ns/op",
            "extra": "6300 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - B/op",
            "value": 295552,
            "unit": "B/op",
            "extra": "6300 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "6300 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.885,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "162375349 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.885,
            "unit": "ns/op",
            "extra": "162375349 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "162375349 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "162375349 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.333,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "160789108 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.333,
            "unit": "ns/op",
            "extra": "160789108 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "160789108 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "160789108 times\n2 procs"
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
          "id": "ed21e17af4bf543aa7ac6357c016c36385da0535",
          "message": "compare: add cpu/mem/mutex/block subcommands",
          "timestamp": "2026-05-14T13:44:57+02:00",
          "tree_id": "3f183a00e8ed45047b511fbbbfb8d6f4a453d598",
          "url": "https://github.com/MichaelMure/benchspotter/commit/ed21e17af4bf543aa7ac6357c016c36385da0535"
        },
        "date": 1778759245563,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter)",
            "value": 1931,
            "unit": "ns/op\t    2144 B/op\t      31 allocs/op",
            "extra": "607432 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 1931,
            "unit": "ns/op",
            "extra": "607432 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 2144,
            "unit": "B/op",
            "extra": "607432 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "607432 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter)",
            "value": 1054,
            "unit": "ns/op\t     640 B/op\t      10 allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 1054,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter)",
            "value": 140743,
            "unit": "ns/op\t  174944 B/op\t    2024 allocs/op",
            "extra": "7896 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 140743,
            "unit": "ns/op",
            "extra": "7896 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 174944,
            "unit": "B/op",
            "extra": "7896 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 2024,
            "unit": "allocs/op",
            "extra": "7896 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 102822,
            "unit": "ns/op\t   64011 B/op\t    1000 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 102822,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64011,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter)",
            "value": 29672651,
            "unit": "ns/op\t24685032 B/op\t  200056 allocs/op",
            "extra": "42 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 29672651,
            "unit": "ns/op",
            "extra": "42 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 24685032,
            "unit": "B/op",
            "extra": "42 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 200056,
            "unit": "allocs/op",
            "extra": "42 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 11643035,
            "unit": "ns/op\t 6584700 B/op\t  101010 allocs/op",
            "extra": "99 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 11643035,
            "unit": "ns/op",
            "extra": "99 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6584700,
            "unit": "B/op",
            "extra": "99 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101010,
            "unit": "allocs/op",
            "extra": "99 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter)",
            "value": 9741,
            "unit": "ns/op\t   11440 B/op\t     131 allocs/op",
            "extra": "132882 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 9741,
            "unit": "ns/op",
            "extra": "132882 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 11440,
            "unit": "B/op",
            "extra": "132882 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "132882 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter)",
            "value": 7037,
            "unit": "ns/op\t    6400 B/op\t     100 allocs/op",
            "extra": "168964 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 7037,
            "unit": "ns/op",
            "extra": "168964 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6400,
            "unit": "B/op",
            "extra": "168964 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "168964 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter)",
            "value": 812271,
            "unit": "ns/op\t 1035123 B/op\t   11038 allocs/op",
            "extra": "1472 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 812271,
            "unit": "ns/op",
            "extra": "1472 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 1035123,
            "unit": "B/op",
            "extra": "1472 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 11038,
            "unit": "allocs/op",
            "extra": "1472 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 702926,
            "unit": "ns/op\t  640234 B/op\t   10000 allocs/op",
            "extra": "1689 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 702926,
            "unit": "ns/op",
            "extra": "1689 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640234,
            "unit": "B/op",
            "extra": "1689 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10000,
            "unit": "allocs/op",
            "extra": "1689 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter)",
            "value": 106125190,
            "unit": "ns/op\t111012726 B/op\t 1100072 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 106125190,
            "unit": "ns/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 111012726,
            "unit": "B/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1100072,
            "unit": "allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 78514280,
            "unit": "ns/op\t67358052 B/op\t 1007148 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 78514280,
            "unit": "ns/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 67358052,
            "unit": "B/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1007148,
            "unit": "allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter)",
            "value": 75437,
            "unit": "ns/op\t  103280 B/op\t    1041 allocs/op",
            "extra": "15546 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 75437,
            "unit": "ns/op",
            "extra": "15546 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 103280,
            "unit": "B/op",
            "extra": "15546 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1041,
            "unit": "allocs/op",
            "extra": "15546 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter)",
            "value": 65439,
            "unit": "ns/op\t   64002 B/op\t    1000 allocs/op",
            "extra": "18586 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 65439,
            "unit": "ns/op",
            "extra": "18586 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64002,
            "unit": "B/op",
            "extra": "18586 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "18586 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter)",
            "value": 8626590,
            "unit": "ns/op\t 9670513 B/op\t  101052 allocs/op",
            "extra": "139 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 8626590,
            "unit": "ns/op",
            "extra": "139 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 9670513,
            "unit": "B/op",
            "extra": "139 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101052,
            "unit": "allocs/op",
            "extra": "139 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 6810650,
            "unit": "ns/op\t 6420189 B/op\t  100006 allocs/op",
            "extra": "162 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 6810650,
            "unit": "ns/op",
            "extra": "162 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6420189,
            "unit": "B/op",
            "extra": "162 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100006,
            "unit": "allocs/op",
            "extra": "162 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter)",
            "value": 819214404,
            "unit": "ns/op\t974041976 B/op\t10100088 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 819214404,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 974041976,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10100088,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 772605780,
            "unit": "ns/op\t807020992 B/op\t10050044 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 772605780,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 807020992,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10050044,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter)",
            "value": 6122,
            "unit": "ns/op\t    7736 B/op\t      80 allocs/op",
            "extra": "174955 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 6122,
            "unit": "ns/op",
            "extra": "174955 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7736,
            "unit": "B/op",
            "extra": "174955 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 80,
            "unit": "allocs/op",
            "extra": "174955 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter)",
            "value": 456444,
            "unit": "ns/op\t  727323 B/op\t    5180 allocs/op",
            "extra": "2682 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 456444,
            "unit": "ns/op",
            "extra": "2682 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 727323,
            "unit": "B/op",
            "extra": "2682 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 5180,
            "unit": "allocs/op",
            "extra": "2682 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter)",
            "value": 45279752,
            "unit": "ns/op\t67574841 B/op\t  501554 allocs/op",
            "extra": "31 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 45279752,
            "unit": "ns/op",
            "extra": "31 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 67574841,
            "unit": "B/op",
            "extra": "31 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 501554,
            "unit": "allocs/op",
            "extra": "31 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter)",
            "value": 6673,
            "unit": "ns/op\t    7888 B/op\t      87 allocs/op",
            "extra": "180889 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 6673,
            "unit": "ns/op",
            "extra": "180889 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7888,
            "unit": "B/op",
            "extra": "180889 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "180889 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter)",
            "value": 62454,
            "unit": "ns/op\t   78368 B/op\t     750 allocs/op",
            "extra": "19204 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 62454,
            "unit": "ns/op",
            "extra": "19204 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 78368,
            "unit": "B/op",
            "extra": "19204 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 750,
            "unit": "allocs/op",
            "extra": "19204 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter)",
            "value": 621354,
            "unit": "ns/op\t  768977 B/op\t    7285 allocs/op",
            "extra": "1904 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 621354,
            "unit": "ns/op",
            "extra": "1904 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 768977,
            "unit": "B/op",
            "extra": "1904 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 7285,
            "unit": "allocs/op",
            "extra": "1904 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter)",
            "value": 2107,
            "unit": "ns/op\t    1344 B/op\t      33 allocs/op",
            "extra": "557846 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - ns/op",
            "value": 2107,
            "unit": "ns/op",
            "extra": "557846 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - B/op",
            "value": 1344,
            "unit": "B/op",
            "extra": "557846 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "557846 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 155,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7345161 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 155,
            "unit": "ns/op",
            "extra": "7345161 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7345161 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7345161 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 153.7,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7882710 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 153.7,
            "unit": "ns/op",
            "extra": "7882710 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7882710 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7882710 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 158.9,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7138261 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 158.9,
            "unit": "ns/op",
            "extra": "7138261 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7138261 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7138261 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 7.458,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "157299842 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 7.458,
            "unit": "ns/op",
            "extra": "157299842 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "157299842 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "157299842 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 7.311,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "163568701 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 7.311,
            "unit": "ns/op",
            "extra": "163568701 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "163568701 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "163568701 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal)",
            "value": 482735,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2491 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - ns/op",
            "value": 482735,
            "unit": "ns/op",
            "extra": "2491 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2491 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2491 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal)",
            "value": 187137,
            "unit": "ns/op\t  295552 B/op\t      33 allocs/op",
            "extra": "5902 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - ns/op",
            "value": 187137,
            "unit": "ns/op",
            "extra": "5902 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - B/op",
            "value": 295552,
            "unit": "B/op",
            "extra": "5902 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "5902 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.573,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "160496540 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.573,
            "unit": "ns/op",
            "extra": "160496540 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "160496540 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "160496540 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.476,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "160699644 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.476,
            "unit": "ns/op",
            "extra": "160699644 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "160699644 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "160699644 times\n2 procs"
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
          "id": "a4f245d2155735641fdee3bac6dedc97cf2db99a",
          "message": "add \"compare inline\" and \"compare escape\" commands",
          "timestamp": "2026-05-14T17:43:22+02:00",
          "tree_id": "0bfc6a81681dfd333ea315678dc138530af7be13",
          "url": "https://github.com/MichaelMure/benchspotter/commit/a4f245d2155735641fdee3bac6dedc97cf2db99a"
        },
        "date": 1778773504308,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter)",
            "value": 1774,
            "unit": "ns/op\t    2144 B/op\t      31 allocs/op",
            "extra": "646970 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 1774,
            "unit": "ns/op",
            "extra": "646970 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 2144,
            "unit": "B/op",
            "extra": "646970 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "646970 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter)",
            "value": 1013,
            "unit": "ns/op\t     640 B/op\t      10 allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 1013,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter)",
            "value": 133787,
            "unit": "ns/op\t  174944 B/op\t    2024 allocs/op",
            "extra": "7771 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 133787,
            "unit": "ns/op",
            "extra": "7771 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 174944,
            "unit": "B/op",
            "extra": "7771 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 2024,
            "unit": "allocs/op",
            "extra": "7771 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 101989,
            "unit": "ns/op\t   64011 B/op\t    1000 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 101989,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64011,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter)",
            "value": 29526641,
            "unit": "ns/op\t24685029 B/op\t  200056 allocs/op",
            "extra": "43 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 29526641,
            "unit": "ns/op",
            "extra": "43 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 24685029,
            "unit": "B/op",
            "extra": "43 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 200056,
            "unit": "allocs/op",
            "extra": "43 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 11263337,
            "unit": "ns/op\t 6575818 B/op\t  100962 allocs/op",
            "extra": "104 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 11263337,
            "unit": "ns/op",
            "extra": "104 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6575818,
            "unit": "B/op",
            "extra": "104 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/1x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100962,
            "unit": "allocs/op",
            "extra": "104 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter)",
            "value": 9093,
            "unit": "ns/op\t   11440 B/op\t     131 allocs/op",
            "extra": "133640 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 9093,
            "unit": "ns/op",
            "extra": "133640 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 11440,
            "unit": "B/op",
            "extra": "133640 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "133640 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter)",
            "value": 6969,
            "unit": "ns/op\t    6400 B/op\t     100 allocs/op",
            "extra": "164403 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 6969,
            "unit": "ns/op",
            "extra": "164403 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6400,
            "unit": "B/op",
            "extra": "164403 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "164403 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter)",
            "value": 823576,
            "unit": "ns/op\t 1035120 B/op\t   11038 allocs/op",
            "extra": "1459 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 823576,
            "unit": "ns/op",
            "extra": "1459 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 1035120,
            "unit": "B/op",
            "extra": "1459 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 11038,
            "unit": "allocs/op",
            "extra": "1459 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 703329,
            "unit": "ns/op\t  640262 B/op\t   10000 allocs/op",
            "extra": "1533 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 703329,
            "unit": "ns/op",
            "extra": "1533 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 640262,
            "unit": "B/op",
            "extra": "1533 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10000,
            "unit": "allocs/op",
            "extra": "1533 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter)",
            "value": 104789787,
            "unit": "ns/op\t111012724 B/op\t 1100072 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 104789787,
            "unit": "ns/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 111012724,
            "unit": "B/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1100072,
            "unit": "allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 76931338,
            "unit": "ns/op\t67134182 B/op\t 1006671 allocs/op",
            "extra": "15 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 76931338,
            "unit": "ns/op",
            "extra": "15 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 67134182,
            "unit": "B/op",
            "extra": "15 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/10x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1006671,
            "unit": "allocs/op",
            "extra": "15 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter)",
            "value": 72027,
            "unit": "ns/op\t  103281 B/op\t    1041 allocs/op",
            "extra": "16561 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 72027,
            "unit": "ns/op",
            "extra": "16561 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - B/op",
            "value": 103281,
            "unit": "B/op",
            "extra": "16561 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1041,
            "unit": "allocs/op",
            "extra": "16561 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter)",
            "value": 62967,
            "unit": "ns/op\t   64002 B/op\t    1000 allocs/op",
            "extra": "18831 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 62967,
            "unit": "ns/op",
            "extra": "18831 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 64002,
            "unit": "B/op",
            "extra": "18831 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x10/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 1000,
            "unit": "allocs/op",
            "extra": "18831 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter)",
            "value": 8045534,
            "unit": "ns/op\t 9670516 B/op\t  101052 allocs/op",
            "extra": "148 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 8045534,
            "unit": "ns/op",
            "extra": "148 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 9670516,
            "unit": "B/op",
            "extra": "148 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 101052,
            "unit": "allocs/op",
            "extra": "148 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter)",
            "value": 6662717,
            "unit": "ns/op\t 6417978 B/op\t  100006 allocs/op",
            "extra": "183 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 6662717,
            "unit": "ns/op",
            "extra": "183 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 6417978,
            "unit": "B/op",
            "extra": "183 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x1000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 100006,
            "unit": "allocs/op",
            "extra": "183 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter)",
            "value": 821255358,
            "unit": "ns/op\t974042208 B/op\t10100090 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - ns/op",
            "value": 821255358,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - B/op",
            "value": 974042208,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/new (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10100090,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter)",
            "value": 796282252,
            "unit": "ns/op\t807021104 B/op\t10050045 allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - ns/op",
            "value": 796282252,
            "unit": "ns/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - B/op",
            "value": 807021104,
            "unit": "B/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkTable/100x100000/reuse (benchspotter/commands/tabwriter) - allocs/op",
            "value": 10050045,
            "unit": "allocs/op",
            "extra": "2 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter)",
            "value": 6073,
            "unit": "ns/op\t    7736 B/op\t      80 allocs/op",
            "extra": "179685 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 6073,
            "unit": "ns/op",
            "extra": "179685 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7736,
            "unit": "B/op",
            "extra": "179685 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 80,
            "unit": "allocs/op",
            "extra": "179685 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter)",
            "value": 449076,
            "unit": "ns/op\t  727325 B/op\t    5180 allocs/op",
            "extra": "2229 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 449076,
            "unit": "ns/op",
            "extra": "2229 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 727325,
            "unit": "B/op",
            "extra": "2229 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 5180,
            "unit": "allocs/op",
            "extra": "2229 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter)",
            "value": 43978717,
            "unit": "ns/op\t67574916 B/op\t  501554 allocs/op",
            "extra": "32 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 43978717,
            "unit": "ns/op",
            "extra": "32 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 67574916,
            "unit": "B/op",
            "extra": "32 times\n2 procs"
          },
          {
            "name": "BenchmarkPyramid/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 501554,
            "unit": "allocs/op",
            "extra": "32 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter)",
            "value": 6436,
            "unit": "ns/op\t    7888 B/op\t      87 allocs/op",
            "extra": "172780 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - ns/op",
            "value": 6436,
            "unit": "ns/op",
            "extra": "172780 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - B/op",
            "value": 7888,
            "unit": "B/op",
            "extra": "172780 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/10 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "172780 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter)",
            "value": 61999,
            "unit": "ns/op\t   78368 B/op\t     750 allocs/op",
            "extra": "19525 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - ns/op",
            "value": 61999,
            "unit": "ns/op",
            "extra": "19525 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - B/op",
            "value": 78368,
            "unit": "B/op",
            "extra": "19525 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/100 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 750,
            "unit": "allocs/op",
            "extra": "19525 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter)",
            "value": 617543,
            "unit": "ns/op\t  768980 B/op\t    7285 allocs/op",
            "extra": "1892 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - ns/op",
            "value": 617543,
            "unit": "ns/op",
            "extra": "1892 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - B/op",
            "value": 768980,
            "unit": "B/op",
            "extra": "1892 times\n2 procs"
          },
          {
            "name": "BenchmarkRagged/1000 (benchspotter/commands/tabwriter) - allocs/op",
            "value": 7285,
            "unit": "allocs/op",
            "extra": "1892 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter)",
            "value": 1987,
            "unit": "ns/op\t    1344 B/op\t      33 allocs/op",
            "extra": "568315 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - ns/op",
            "value": 1987,
            "unit": "ns/op",
            "extra": "568315 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - B/op",
            "value": 1344,
            "unit": "B/op",
            "extra": "568315 times\n2 procs"
          },
          {
            "name": "BenchmarkCode (benchspotter/commands/tabwriter) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "568315 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal)",
            "value": 147.2,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "8091331 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - ns/op",
            "value": 147.2,
            "unit": "ns/op",
            "extra": "8091331 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "8091331 times\n2 procs"
          },
          {
            "name": "BenchmarkFoo (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8091331 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal)",
            "value": 143.1,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "8456052 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - ns/op",
            "value": 143.1,
            "unit": "ns/op",
            "extra": "8456052 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "8456052 times\n2 procs"
          },
          {
            "name": "BenchmarkBar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8456052 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal)",
            "value": 149.1,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "7736070 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - ns/op",
            "value": 149.1,
            "unit": "ns/op",
            "extra": "7736070 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "7736070 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz/foo/bar (benchspotter/engine/internal) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7736070 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal)",
            "value": 7.959,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "158600815 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - ns/op",
            "value": 7.959,
            "unit": "ns/op",
            "extra": "158600815 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "158600815 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/first (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "158600815 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal)",
            "value": 7.559,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "158500428 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - ns/op",
            "value": 7.559,
            "unit": "ns/op",
            "extra": "158500428 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "158500428 times\n2 procs"
          },
          {
            "name": "BenchmarkSiblings/second (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "158500428 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal)",
            "value": 529079,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2235 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - ns/op",
            "value": 529079,
            "unit": "ns/op",
            "extra": "2235 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2235 times\n2 procs"
          },
          {
            "name": "BenchmarkHybridSort (benchspotter/engine/internal) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2235 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal)",
            "value": 153102,
            "unit": "ns/op\t  295552 B/op\t      33 allocs/op",
            "extra": "6968 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - ns/op",
            "value": 153102,
            "unit": "ns/op",
            "extra": "6968 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - B/op",
            "value": 295552,
            "unit": "B/op",
            "extra": "6968 times\n2 procs"
          },
          {
            "name": "BenchmarkMapOps (benchspotter/engine/internal) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "6968 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.795,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "158138319 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.795,
            "unit": "ns/op",
            "extra": "158138319 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "158138319 times\n2 procs"
          },
          {
            "name": "BenchmarkBaz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "158138319 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage)",
            "value": 7.574,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "158210851 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - ns/op",
            "value": 7.574,
            "unit": "ns/op",
            "extra": "158210851 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "158210851 times\n2 procs"
          },
          {
            "name": "BenchmarkBoz (benchspotter/engine/internal/anotherpackage) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "158210851 times\n2 procs"
          }
        ]
      }
    ]
  }
}