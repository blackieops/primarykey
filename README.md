# `go.b8s.dev/primarykey`

[![Go Report Card](https://goreportcard.com/badge/go.b8s.dev/primarykey)](https://goreportcard.com/report/go.b8s.dev/primarykey)

The goal with `primarykey` is to provide a bridge between
[`github.com/google/uuid`][0] and [`github.com/lithammer/shortuuid`][1] while also
supporting database interfaces so they can continue to be used as column types
in things like [`gorm.io/gorm`][2].

## Usage

The `primarykey.ID` type is byte-compatible with UUIDs, and provides all the
interfaces required to be used as a database column type. For example, if you
have a [gorm][2] model:

```go
import "go.b8s.dev/primarykey"

type MyModel struct {
	ID primarykey.ID
}
```

There is also a public interface to encode and decode directly:

```go
newOne := primarykey.New()
Encode(newOne) //=> "KwSysDpxcBU9FNhGkn2dCf"

id := Decode("KwSysDpxcBU9FNhGkn2dCf") //=> ID
id.UUID() //=> uuid.UUID
id.String() //=> "KwSysDpxcBU9FNhGkn2dCf"
```

## Development

This is a very standard Go project with very minimal dependencies and no
required setup. Dependencies are vendored.

To run the test suite,

```
$ go test .
```

That's about it.

[0]: https://github.com/google/uuid
[1]: https://github.com/lithammer/shortuuid
[2]: https://gorm.io/gorm
