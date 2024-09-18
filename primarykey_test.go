package primarykey

import (
	"encoding/json"
	"testing"
)

var testUUIDBytes = []byte{
	0x7d, 0x44, 0x48, 0x40,
	0x9d, 0xc0,
	0x11, 0xd1,
	0xb2, 0x45,
	0x5f, 0xfd, 0xce, 0x74, 0xfa, 0xd2,
}

func TestIDString(t *testing.T) {
	id, err := FromBytes(testUUIDBytes)
	if err != nil {
		t.Errorf("Error while initializing ID from bytes: %v", err)
		return
	}
	if id.String() != "QJQQv92onGb5t9gCsNLDgT" {
		t.Errorf("String() generated invalid uuid: %v", id.String())
	}
}

func TestIDToUUID(t *testing.T) {
	id, err := FromBytes(testUUIDBytes)
	if err != nil {
		t.Errorf("Error while initializing ID from bytes: %v", err)
		return
	}
	if id.UUID().String() != "7d444840-9dc0-11d1-b245-5ffdce74fad2" {
		t.Errorf("Long UUID was invalid: %v", id.UUID().String())
	}
}

func TestIDScan(t *testing.T) {
	var id *ID
	var err error

	id = new(ID)
	err = id.Scan("7d444840-9dc0-11d1-b245-5ffdce74fad2")
	if err != nil {
		t.Errorf("Scan from string failed: %v", err)
	}

	id = new(ID)
	err = id.Scan(testUUIDBytes)
	if err != nil {
		t.Errorf("Scan from bytes failed: %v", err)
	}

	id = new(ID)
	err = id.Scan(nil)
	if err != nil {
		t.Errorf("Scan from nil failed: %v", err)
	}

	id = new(ID)
	err = id.Scan(123)
	if err == nil {
		t.Errorf("Scan from invalid type should have failed!")
	}
}

func TestIDValue(t *testing.T) {
	id, err := FromBytes(testUUIDBytes)
	if err != nil {
		t.Errorf("Error while initializing ID from bytes: %v", err)
		return
	}

	value, err := id.Value()

	if err != nil {
		t.Errorf("Value() encountered error: %v", err)
	}
	if value != "7d444840-9dc0-11d1-b245-5ffdce74fad2" {
		t.Errorf("Value had invalid result: %v", value)
	}
}

func TestDecode(t *testing.T) {
	id, err := Decode("QJQQv92onGb5t9gCsNLDgT")
	if err != nil {
		t.Fatalf("Decode failed to decode shortuuid: %v", err)
	}
	for i, b := range testUUIDBytes {
		if id[i] != b {
			t.Fatal("DecodedID does not match expected bytes!")
		}
	}
}

func TestMustDecodeValid(t *testing.T) {
	id := MustDecode("QJQQv92onGb5t9gCsNLDgT")
	for i, b := range testUUIDBytes {
		if id[i] != b {
			t.Fatal("Decoded ID does not match expected bytes!")
		}
	}
}

func TestMustDecodeInvalid(t *testing.T) {
	defer func() {
		recover()
	}()
	MustDecode("1234123441234")
	t.Error("MustDecode should have panicked with invalid shortuuid but did not!")
}

func TestEncode(t *testing.T) {
	id, err := FromBytes(testUUIDBytes)
	if err != nil {
		t.Errorf("Error while initializing ID from bytes: %v", err)
		return
	}

	result := Encode(id)

	if result != "QJQQv92onGb5t9gCsNLDgT" {
		t.Errorf("Encode() returned unexpected value: %v", result)
	}
}

func TestEmpty(t *testing.T) {
	id := Empty()
	if id != [16]byte{} {
		t.Errorf("Empty ID was not the right kind of empty! Got: %v", id)
	}
}

func TestNew(t *testing.T) {
	id := New()

	if len(id) != 16 {
		t.Errorf("New() gave weird output: %v", id)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	type model struct {
		ID ID
	}
	src := `{"ID": "gXeZJmzG3xwqWdJeumvbFy"}`

	var m *model
	if err := json.Unmarshal([]byte(src), &m); err != nil {
		t.Fatalf("Failed to unmarshal ID from json: %v", err)
	}

	if m.ID.String() != "gXeZJmzG3xwqWdJeumvbFy" {
		t.Fatalf("Got unexpected ID value: %s", m.ID)
	}
}

func TestMarshalJSON(t *testing.T) {
	type model struct {
		ID ID
	}

	m := &model{ID: MustDecode("gXeZJmzG3xwqWdJeumvbFy")}

	b, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("Failed to unmarshal ID from json: %v", err)
	}

	if string(b) != `{"ID":"gXeZJmzG3xwqWdJeumvbFy"}` {
		t.Fatalf("Got unexpected JSON marshaled: %s", string(b))
	}
}
