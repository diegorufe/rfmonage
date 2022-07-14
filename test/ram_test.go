package test

import (
	"fmt"
	"rfmonage/system/ram"
	"testing"
)

func exampleTotalMemory() {
	fmt.Printf("Total system memory: %d\n", ram.TotalMemory())
}
func exampleFreeMemory() {
	fmt.Printf("Free system memory: %d\n", ram.FreeMemory())
}

func exampleMemoryInUse() {
	fmt.Printf("Memory in use: %d\n", ram.MemoryInUse())
}

func TestTotalRam(t *testing.T) {
	exampleFreeMemory()

	exampleTotalMemory()

	exampleMemoryInUse()

	var estimateMaxRam uint64 = 18000000000
	totalRam := ram.TotalMemory()

	if estimateMaxRam < totalRam {
		t.Errorf("Expected: %v, got: %v", estimateMaxRam, totalRam)
	}
}
