package primarykey

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v4"
)

// ID is a type that stores the bytes of a UUID. It delegates a lot of the
// handling and parsing to `github.com/google/uuid` and stringifies the UUIDs
// using `github.com/lithammer/shortuuid`. It is compatible with the various
// database interfaces required to use it as a UUID column type.
type ID [16]byte

// String returns the shortuuid string representation of the UUID.
func (i ID) String() string {
	return shortuuid.DefaultEncoder.Encode(i.UUID())
}

// UUID returns the underlying "long" UUID, which can be useful for
// compatibility with things that need the not-short UUID.
func (i ID) UUID() uuid.UUID {
	return uuid.UUID(i)
}

// Scan implements the sql.Scanner interface to support reading UUID values out
// of a database into an ID type.
func (i *ID) Scan(src interface{}) error {
	var u uuid.UUID
	err := u.Scan(src)
	if err != nil {
		return err
	}
	*i = ID(u)
	return nil
}

// Value implements the sql.Valuer interface to support converting the UUID to
// a value that can be stored by a database driver.
func (i ID) Value() (driver.Value, error) {
	return i.UUID().String(), nil
}

// MarshalJSON satisfies the [encoding/json.Marshaler] interface so IDs get
// serialized to JSON in their short string form automatically.
func (i ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON satisfies the [encoding/json.Unmarshaler] interface so the
// stringified IDs get decoded into [ID] structs automatically.
func (i *ID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	v, err := Decode(s)
	if err != nil {
		return err
	}
	*i = v
	return nil
}

// Decode parses the string as a shortuuid into an `ID`.
func Decode(src string) (ID, error) {
	u, err := shortuuid.DefaultEncoder.Decode(src)
	return ID(u), err
}

// MustDecode parses the string as an ID, similar to `Decode`, but panics if
// the string cannot be parsed.
func MustDecode(src string) ID {
	id, err := Decode(src)
	if err != nil {
		panic(err)
	}
	return id
}

// Encode takes an ID and encodes it to a shortuuid string. This is equivalent
// to calling `String()` on the `ID`, but is provided for symmetry.
func Encode(id ID) string {
	return id.String()
}

// New generates a new ID.
func New() ID {
	return ID(uuid.New())
}

// Empty returns a "valid", but zeroed-out UUID.
func Empty() ID {
	return [16]byte{}
}

// FromBytes initializes an [ID] from a UUID byte slice.
func FromBytes(b []byte) (ID, error) {
	uid, err := uuid.FromBytes(b)
	if err != nil {
		return Empty(), err
	}
	return ID(uid), nil
}
