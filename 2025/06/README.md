# 06

For part 1, I felt it was clever to trade time for space by accumulating both additions and multiplications in a single pass, and then selecting only the results needed at the end. This means we're doing work upfront that we'll never use, but it also means we only have to loop through the input once. 

By also using `strings.FieldsSeq` to parse each row, we avoid allocating a slice to hold all the values each row, we save some memory there as well.

### Throwaway Err

- For quick scripts, we can noop on errors

```go
num, _ := strconv.Atoi(field)
```

### Better Stats

```go
// PrintMemUsage outputs the current memory statistics.
// Usage: defer PrintMemUsage()
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// Helper: Convert bytes to Mebibytes (MiB) rounded to 2 decimal places.
	bToMb := func(b uint64) float64 {
		return float64(b) / 1024.0 / 1024.0
	}

	fmt.Println()
	fmt.Println("====== Memory Stats ======")

	// Alloc (HeapAlloc):
	// The number of bytes currently allocated and in use.
	// This represents the actual RAM footprint of your solution right now.
	fmt.Printf("Current Heap:   %.5f MiB\n", bToMb(m.Alloc))

	// TotalAlloc:
	// The cumulative count of bytes allocated (even if they were freed).
	// A high number here with a low 'Current Heap' indicates high "churn"
	// (creating and discarding many temporary objects).
	fmt.Printf("Total Alloc:    %.5f MiB\n", bToMb(m.TotalAlloc))

	// Mallocs:
	// The total count of heap objects allocated since the program started.
	// In Go, high object counts can slow down the Garbage Collector.
	fmt.Printf("Mallocs:        %d objects\n", m.Mallocs)

	// Live Objects (Mallocs - Frees):
	// The number of objects currently alive on the heap.
	// If this number is very high, the GC has more work to do during scans.
	fmt.Printf("Live Objects:   %d objects\n", m.Mallocs-m.Frees)

	// NumGC:
	// The number of times the Garbage Collector has run.
	fmt.Printf("GC Cycles:      %d\n", m.NumGC)

	fmt.Println("==========================")
}
```
