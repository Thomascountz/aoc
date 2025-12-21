# 03

### Named vs Naked Returns

- Functions can have named return values (e.g. `func f() (result int) { ... return }`), which allows for calling `return` without a value
  - Go will return the value stored in the `result` variable at the time of the `return` statement
- This is generally discouraged except in short functions, furthermore, named returns are automatically initialized to their zero values

### Iterating on Strings (Runes)

- Go strings are essentially immutable byte slices, i.e. `[]byte` without a "capacity" concept.
- Indexing a string with `str[i]` yields the byte at index `i`, regardless of encoding.
  - I.e., a multi-byte UTF-8 character will return only the first byte of that character.
- Ranging over a string with `for i, c := range str` yields the byte index `i` and the rune (alias for `int32`) `c` at that position.
  - Note that runes can be multi-byte UTF-8, therefore `i` may not increment by 1 for each iteration.
- A string literal (e.g. `"hello"`) is UTF-8 encoded by default (unless there are escape sequences).
- UTF-8 encodes ASCII characters as single-byte equivalents, i.e., the byte value of `"A"` is `65` in both ASCII and UTF-8.
  
### Converting a "Character" to an Integer

- Strings `"0"` to `"9"` are represented as single-byte ASCII
- They can be converted to their integer values by subtracting the rune value of `'0'` from the rune value of the character.
- In Unicode, `0` is `48`, and the digits are sequentially ordered until `9` which is `57`.
- Therefore, `rune('3') - rune('0')` yields `3`.

### "Bit-Shifting" Integers in Base 10
- Given a sequence of digits `[]int{1, 2, 3}`, you can accumulate the value `123` by iterating over the digits and multiplying the current total by `10` before adding the next digit.
- This is bit-shifting but in base 10.

```go
digits := []int{1, 2, 3}
total := 0
for _, d := range digits {
	total = total*10 + d
}
fmt.Println(total) // prints 123
```

### Scanner Errors

> Scan advances the Scanner to the next token, which will then be available through the Scanner.Bytes or Scanner.Text method. It returns false when there are no more tokens, either by reaching the end of the input or an error. After Scan returns false, the Scanner.Err method will return any error that occurred during scanning, except that if it was io.EOF, Scanner.Err will return nil. Scan panics if the split function returns too many empty tokens without advancing the input. This is a common error mode for scanners. 
>
> source: https://pkg.go.dev/bufio#Scanner.Scan

- After calling `scanner.Scan()`, you can check for errors with `scanner.Err()`.

```go
scanner := bufio.NewScanner(file)
if err := scanner.Err(); err != nil {
  log.Fatal(err)
}
```
