package builtins

import (
	"fmt"
	"runtime"
)
func Alloc() {
//may have to allocate memory yourself in the function to test it works 
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Printf("Total Alloc: %v MB\n", bToMb(memStats.Alloc))
	fmt.Printf("Total Free:  %v MB\n", bToMb(memStats.TotalAlloc))
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}