package main

import (
	"fmt"

	"benchspotter/locate"
)

func main() {
	benchmarks, err := locate.Benchmarks(".")
	if err != nil {
		panic(err)
	}
	for _, b := range benchmarks {
		fmt.Println(b)
	}
}
