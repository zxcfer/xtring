# xtring

String parsing package for Go. Converts strings to best guess value type.

Supported types (in this order):

  * `nil` or `null`
  * `bool`
  * `int`
  * `int64`
  * `uint`
  * `uint64`
  * `float64`

Quoted values are always strings with the quotes stripped.

## Get started

```
go get gopkg.in/matryer/xtring.v1
```

## Usage

The simplest usage:

```
val := xtring.Parse("123")
```

To add additional `ParseFunc` functions:

```
parser := append(xtring.DefaultParser, toMyRangeType)
val := parser.Parse("[1,3]")
```

To use your own parsers entirely, use `New`:

```
parser := xtring.New(toMyRangeType)
val := parser.Parse("[1,3]")
```

### Time

To add parsers for time, import the `parsers` package and use the `Time` method:

```
parser := append(xtring.DefaultParser, parsers.Time(time.RFC822))
val := parser.Parse("02 Jan 06 15:04 MST")
```