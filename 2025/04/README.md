# 04

I accidentally found out you don't need a double buffer for this puzzle. Unlike Game of Life, you can update the grid in place because the grid will converge to a maximal number of removed rolls.

### `string` vs `[]byte`
- `string`s are essentially byte slices (`[]byte`), but they are immutable/need to be copied.
- `[]byte` can be used without copying/allocating new memory, though like all slices, if you append past capacity, new memory will be allocated.

### Slices are References
- Slices are reference types, meaning that when you pass a slice to a function or assign it to another variable, you are passing a reference to the "slice header"
- The "slice header" contains a pointer to the underlying array, the length of the slice, and its capacity.

### runtime.MemStats
- The `runtime` package provides access to runtime information about the Go program, including allocs and garbage collection stats.

```go
func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	fmt.Printf("Total Mallocs: %d\n", m.Mallocs)
}
```
