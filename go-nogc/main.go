package main

import (
	"fmt"
	"runtime"
)

const N = 1_000_000

func a() {
	s := make([]int, N)
	for i := range N {
		s = append(s, i)
	}
}

func main() {
	a()
	mem()
	for {
	}
}

func mem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("HeapAlloc = %v B\n", m.HeapAlloc)
	fmt.Printf("TotalAlloc = %v B\n", m.TotalAlloc)
	fmt.Printf("Sys = %v B\n", m.Sys)
	fmt.Printf("NumGC = %v\n", m.NumGC)
}
