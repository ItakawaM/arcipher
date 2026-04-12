package benchmark

import (
	"fmt"
	"os"
	"runtime"
	"text/tabwriter"
	"time"
)

func getMemoryUsage() string {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	return fmt.Sprintf("Alloc\t%v MB\nTotalAlloc\t%v MB\nSys\t%v MB\nNumGC\t%v\n", bToMb(stats.Alloc), bToMb(stats.TotalAlloc), bToMb(stats.Sys), bToMb(uint64(stats.NumGC)))
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func MeasurePerformance(name string) func() {
	start := time.Now()
	tab := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	return func() {
		fmt.Fprintf(tab, "\nPerformance %s\nTime\t%s\n%s", name, time.Since(start), getMemoryUsage())
		tab.Flush()
	}
}
