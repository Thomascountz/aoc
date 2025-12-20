# 03

### Named vs Naked Returns

- Functions can have named return values (e.g. `func f() (result int) { ... return }`), which allows for calling `return` without a value
  - Go will return the value stored in the `result` variable at the time of the `return` statement
- This is generally discouraged except in short functions, furthermore, named returns are automatically initialized to their zero values

### Iterating on Strings (Runes)

- Go strings are readonly byte slices, i.e. `[]byte`. 
- They are arbitrary bytes, not necessarily unicode, UTF-8, or ASCII code points.
  - However, a string literal (e.g. `"hello"`) is UTF-8 encoded by default (unless there are escape sequences).
- Individual bytes which represent character code points are of type `rune`, alias for `int32`.
- Iterating over a string with `for i, c := range str` will yield the byte index `i` and the rune `c` at that position.
  - Note that `i` is the byte index, not the rune index. Runes can be multiple bytes in UTF-8, therefore `i` may not increment by 1 for each iteration.
  
### Converting a "Character" to an Integer

- Strings `"0"` to `"9"` can be converted to their integer values by subtracting the rune value of `'0'` from the rune value of the character.
- In Unicode, `0` is represented by `48`, and the digits are sequentially ordered until `9` which is `57`.
- Therefore, `rune('3') - rune('0')` yields `3`.
