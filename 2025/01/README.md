# 01

### Modulus and Remainders with Negative Numbers

My first solution was to use the modulus and figure out how many times we "carry over" across zero. In that, I learned that the `%` (remainder) operator didn't work the way I expected for negative numbers. That's where this function came in:

```go
func mod(a, b int64) int64 {
	return ((a % b) + b) % b
}
```

This makes it so that `mod(-1, 5)` returns `4`, (like `-1%5` in Ruby) instead of `-1`. Furthermore, `math.Mod(-1.0, 5.0)` returns `-1.0`, so it behaves like Go's `%` operator.

### Math Module

The `func Mod(x, y float64) float64` function operates on `float64`, like nearly all `math` module functions in Go, including `func Abs(x float64) float64`. This is why I wrote an absolute value integer function, `absInt`:

```go
func absInt(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
```

### Padding Using `fmt.Printf`

- Integers
  - To specify the width of an integer, use a number after the `%` in the verb, e.g. `%5d` for a width of 5.
    - This is right-aligned and space-padded by default.
  - To left-align, use a `-` before the width, e.g. `%-5d`
  - To pad with leading zeros, use a `0` before the width, e.g. `%05d`
- Floats
  - With floats, you can also specify the decimal precision using a `.`, e.g. `%.2f` for 2 decimal places.
    - You can combine width and precision, e.g. `%8.2f` for a width of 8 and 2 decimal places.
- Strings
  - Strings can also be padded using width, e.g. `%10s` for a width of 10, or left-align with `%-10s`
  
### Relative File Paths

- I haven't found a way to get a file descriptor relative to the go file being executed, v.s. the current working directory of the process.

### Unreachable Code

- In `zig`, `unreachable` is a built-in that acts as a runtime assertion that the code should never be reached.
- I like that, so in Go, I used `panic("unreachable")` when specifying defining all valid invariants, rather than using an implicit default branch.

### `func (*Scanner) Scan` vs `func (*Reader) ReadLine`
- See: https://pkg.go.dev/bufio
- "Scanner provides a convenient interface for reading data such as a file of newline-delimited lines of text."
  - "Scan advances the Scanner to the next token, which will then be available through the Scanner.Bytes or Scanner.Text method. It returns false when there are no more tokens, either by reaching the end of the input or an error."
- "Reader implements buffering for an io.Reader object."
  - "ReadLine is a low-level line-reading primitive. Most callers should use Reader.ReadBytes('\n') or Reader.ReadString('\n') instead or use a Scanner."

### `strconv`
- `strconv.Atoi` - string to int
- `strconv.Itoa` - int to string
