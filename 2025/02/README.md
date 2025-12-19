# 02

### `if err != nil`

I wrapped this infamous Go idiom in a `func check(e error)` function:

```go
func check(e error) {
	if e != nil {
		panic(e)
	}
}
```

### Slice Operator
- Similar to Ruby's `[low...high]` range slicing on Arrays: `slice[low(inclusive):high(exclusive)]`

### `break`, `continue`, and Labels
- `break` is used to exit a immediate loop early, similar to Ruby's `break`
- `continue` is used to skip to the next iteration of a loop, similar to Ruby's `next`
- Labels can be used with `break` and `continue` to specify which loop to affect
  - This is useful for nested loops
  
```go
package main

import "fmt"

func main() {
outer:
	for i := range 5 {
		for j := range 5 {
			fmt.Printf("i: %d, j: %d\n", i, j)
			if j == i {
				fmt.Println("Continue j == i")
				continue outer
			}

		}
	}
}
```

```
i: 0, j: 0
Continue j == i
i: 1, j: 0
i: 1, j: 1
Continue j == i
i: 2, j: 0
i: 2, j: 1
i: 2, j: 2
Continue j == i
i: 3, j: 0
i: 3, j: 1
i: 3, j: 2
i: 3, j: 3
Continue j == i
i: 4, j: 0
i: 4, j: 1
i: 4, j: 2
i: 4, j: 3
i: 4, j: 4
Continue j == i
```

### Appending to Slices
- To append to a slice, we use the built-in `func append(slice []Type, elems ...Type) []Type` function, which returns a new slice with the new element added.

### Aliased Imports
- You can alias imports to avoid name conflicts or for convenience, e.g. 

```go
import(
	s "strings"
	m "math"
)
```

### `os.ReadFile`
- `os.ReadFile` reads the entire content of a file and returns it as a byte slice (`[]byte`), along with an error if any. 
- I found this more convenient than using `bufio.Scanner` for reading the small, non-newline-delimited file for this challenge.
