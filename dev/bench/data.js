window.BENCHMARK_DATA = {
  "lastUpdate": 1777413003920,
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
      }
    ]
  }
}