# 05

### Enum and `iota`

- Go doesn't have a native "enum" type
- Instead, you can create an "alias type" for `int`, and define a set of named constants of that type
- When defining a set of related constants, you can use the `iota` identifier to create a sequence of incrementing integer values
- Multiple variable or constant declarations in parenthesis will each have the same type and value from the first line

```go
type ParseMode int
const (
	RangeParsing ParseMode = iota // 0
	IdParsing                     // 1
)
```

### Variadic Functions and Slice Expansion
- `append` is a variadic function: `func append([]T, ...T) []T`
- The `...` expands a slice into individual elements when calling a variadic function

```go
// Remove element at index i from slice a
a = append(a[:i], a[i+1:]...)
```

- This pattern is commonly used to **remove an element from a slice**
  - `a[i+1:]` is a slice, but `append` expects elements (`T`), so `...` is required
  - Without `...`, the code does not compile
- Doing this while iterating over a slice isn't a great idea, as it forces Go to shift all the elements in memory

### `slices.BinarySearchFunc`
- Performs a binary search on a **sorted** slice using a custom comparison function
- The comparison function should return:
  - A negative number if the first argument is less than the second
  - Zero if they are equal
  - A positive number if the first argument is greater than the second
  
```go
	_, found := slices.BinarySearchFunc(r, id, func(ir IngredientRange, target int) int {
		if ir.Start > target {
			return 1
		}
		if ir.End < target {
			return -1
		}
		return 0
	})
	return found
	```
	
	- In this example, `id` is being passed into the comparison function as `target`
	- The function checks if `id` is within the range defined by `ir.Start` and `ir.End`, and conforms to the "`cmp`" signature.
