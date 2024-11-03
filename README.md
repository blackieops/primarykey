# `go.b8s.dev/primarykey`

[![Test Suite](https://github.com/blackieops/primarykey/actions/workflows/tests.yml/badge.svg)](https://github.com/blackieops/primarykey/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/go.b8s.dev/primarykey)](https://goreportcard.com/report/go.b8s.dev/primarykey)

`primarykey` is a replacement for UUID primary keys in your database-driven Go
programs. `primarykey` utilises the [`shortuuid`][1] library to generate
binary-compatible UUIDs in a more human-friendly format, allowing you to still
benefit from native UUID storage in systems like PostgreSQL but while providing
less obnoxious IDs to your users.

[1]: https://github.com/lithammer/shortuuid

## Usage

Use `primarykey.ID` for your ID fields in your model structs:

```go
import "go.b8s.dev/primarykey"

type MyModel struct {
	ID primarykey.ID
}
```

The ID will be passed to the datastore as a regular binary UUID (eg., `uuid` in
PostgreSQL).

You can pass primarykeys directly into queries:

```go
db.Query("SELECT * FROM my_model WHERE id = $1", myModel.ID)
```

And marshal them to serializable formats like JSON automatically:

```go
b, _ := json.Marshal(&MyModel{ID: primarykey.New()})
//=> {"ID":"KwSysDpxcBU9FNhGkn2dCf"}
```

There is also a public interface to encode and decode directly:

```go
newOne := primarykey.New()
primarykey.Encode(newOne) //=> "KwSysDpxcBU9FNhGkn2dCf"

id := primarykey.Decode("KwSysDpxcBU9FNhGkn2dCf") //=> ID
id.UUID() //=> uuid.UUID
id.String() //=> "KwSysDpxcBU9FNhGkn2dCf"
```


## Development

This is a very standard Go project with very minimal dependencies and no
required setup. Dependencies are vendored.

To run the test suite:

```
$ go test .
```
