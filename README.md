# go-convert

A lightweight, zero-dependency (except `github.com/pkg/errors`) Go library for safely converting `any` (`interface{}`) values to standard Go types.  
All conversions handle `nil`, type mismatches, and out-of-range values gracefully — returning a sensible zero value or an error instead of panicking.

## Installation

```bash
go get github.com/iostrovok/go-convert
```

## Import

```go
import convert "github.com/iostrovok/go-convert"
```

## Overview

Every converter comes in two flavours:

| Flavour | Signature | On error |
|---------|-----------|----------|
| **Safe** | `Xxx(in any, ...) T` | returns zero value |
| **Error** | `XxxErr(in any, ...) (T, error)` | returns zero value + error |

Most functions also accept optional `debugKeys ...string` that are embedded in error messages to help with tracing the origin of a bad value.

---

## Bool

```go
convert.Bool(in any) bool
convert.BoolErr(in any) (bool, error)
```

| Input type | `false` when | `true` when |
|-----------|-------------|------------|
| `bool` | `false` | `true` |
| Numeric | value == `0` | value != `0` |
| `string` / `[]byte` | `""`, `"false"`, `"f"`, `"0"` | anything else |
| Slice / Map | length == 0 | length > 0 |
| `nil` | always | — |

```go
convert.Bool(0)         // false
convert.Bool("true")    // true
convert.Bool([]int{})   // false
convert.Bool([]int{1})  // true
```

---

## Int

```go
// Safe (returns 0 on error)
convert.Int(in any, debugKeys ...string) int
convert.Int32(in any, debugKeys ...string) int32
convert.Int64(in any, debugKeys ...string) int64

// With error
convert.Int32Err(in any, debugKeys ...string) (int32, error)
convert.Int64Err(in any, debugKeys ...string) (int64, error)

// Slice conversions
convert.ListOfInt32Err(in any, checkLen bool, debugKeys ...string) ([]int32, error)
convert.ListOfInt64Err(in any, checkLen bool, debugKeys ...string) ([]int64, error)
```

Accepted input types: all numeric types, `bool` (`true`→1, `false`→0), `string`, `[]byte`.  
`Int32` / `Int32Err` additionally validate that the value fits within `math.MinInt32`…`math.MaxInt32`.

```go
convert.Int("42")            // 42
convert.Int64(true)          // 1
convert.Int32Err(99999999999) // 0, error (overflow)

ids, err := convert.ListOfInt64Err([]any{"1", "2", "3"}, true)
// ids == []int64{1, 2, 3}
```

---

## Uint

```go
// Safe (returns 0 on error)
convert.Uint32(in any, debugKeys ...string) uint32
convert.Uint64(in any, debugKeys ...string) uint64

// With error
convert.Uint32Err(in any, debugKeys ...string) (uint32, error)
convert.Uint64Err(in any, debugKeys ...string) (uint64, error)
```

Accepted input types: all numeric types, `bool`, `string`, `[]byte`.  
`nil` is treated as `0` (no error).

```go
convert.Uint64("255")   // 255
convert.Uint32Err(-1)   // 0, error
```

---

## Float

```go
// Safe (returns 0 on error)
convert.Float32(in any, debugKeys ...string) float32
convert.Float64(in any, debugKeys ...string) float64

// With error
convert.Float32Err(in any, debugKeys ...string) (float32, error)
convert.Float64Err(in any, debugKeys ...string) (float64, error)

// Low-level (does not filter ±Inf)
convert.BaseFloat64Err(in any, debugKeys ...string) (float64, error)
```

Accepted input types: all numeric types, `bool`, `string`, `[]byte`.  
`Float64Err` additionally filters `±Inf` results.  
`Float32Err` validates that a `float64` input fits within `±math.MaxFloat32`.

```go
convert.Float64("3.14")   // 3.14
convert.Float32(true)     // 1.0
```

---

## String

```go
// Single value
convert.String(in any, _ ...string) string
convert.StringTS(in any, debugKeys ...string) string   // TrimSpace variant

// Slice of strings (include empty elements)
convert.ListOfStrings(in any, debugKeys ...string) []string
convert.ListOfStringsErr(in any, checkLen bool, debugKeys ...string) ([]string, error)

// Slice of strings — skip empty elements
convert.ListOfStringsP(in any, debugKeys ...string) []string
convert.ListOfStringsPErr(in any, checkLen bool, debugKeys ...string) ([]string, error)

// Slice of strings — fail if any element is empty
convert.ListOfStringsStrictPErr(in any, checkLen bool, debugKeys ...string) ([]string, error)
```

`String` converts scalar types using the fastest path (`strconv`). Non-scalar types and `nil` return `""`.  
`StringTS` is identical to `String` but trims surrounding whitespace.

When `checkLen` is `true`, an error is returned if the resulting slice is empty.

```go
convert.String(42)         // "42"
convert.String(3.14)       // "3.14"
convert.StringTS("  hi ") // "hi"

strs, _ := convert.ListOfStringsP([]any{"a", nil, "b", ""})
// strs == []string{"a", "b"}  — nil and empty elements skipped
```

### String list runners

| Constant | Behaviour |
|----------|-----------|
| `convert.PassAll` | include every element (even empty) |
| `convert.SkipEmpty` | silently skip empty strings |
| `convert.FallOnEmpty` | return an error if any element is empty |

---

## Iterator (slice)

A reflection-based iterator over any slice (`[]any`, `[]int`, `[]string`, …).

```go
it, err := convert.Iterator(slice, checkLen ...bool) (*It, error)

it.Len() int
it.NextNotNil() any                          // skips nil elements
it.NextNotEmptyString() string               // skips nil and empty strings
it.NextNotNilMapString() (map[string]any, bool) // skips non-map elements
```

Helper:

```go
convert.CheckMapStringType(t any) (map[string]any, bool)
```

```go
it, _ := convert.Iterator([]any{1, nil, "three"})
for i := 0; i < it.Len(); i++ {
    v := it.NextNotNil()
    fmt.Println(v) // prints: 1, "three"  (nil skipped)
}
```

---

## MapIterator (map)

A reflection-based iterator over any map.

```go
it, err := convert.MapIterator(m, checkLen ...bool) (*MapIt, error)

it.Len() int
it.HasNext() bool
it.Next() (key any, value any)
```

```go
it, _ := convert.MapIterator(map[string]any{"a": 1, "b": 2})
for it.HasNext() {
    k, v := it.Next()
    fmt.Printf("%v = %v\n", k, v)
}
```

---

## Error handling pattern

All `*Err` functions return a descriptive error that includes the bad value and any supplied `debugKeys`:

```go
v, err := convert.Int64Err("bad", "user.age")
// err: strconv.ParseInt: parsing "bad": invalid syntax
```

Use `debugKeys` to annotate where in your data structure the value came from — invaluable when processing deeply nested JSON/map structures.

---

## Requirements

- Go **1.22+**
- [`github.com/pkg/errors`](https://github.com/pkg/errors) v0.9.1

## License

See [LICENSE](LICENSE).
