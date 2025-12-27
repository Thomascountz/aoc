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
