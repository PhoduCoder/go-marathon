package main

import (
	"fmt"
	"runtime"
)

func main() {
	//fmt.Println("Hello, 世界")
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		fmt.Println(s.Sys, s.HeapSys, s.HeapAlloc, s.StackSys, s.StackInuse)
		return s.HeapSys
	}
	fmt.Println(memConsumed())
}
