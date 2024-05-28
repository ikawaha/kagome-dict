package dict

import (
	"bytes"
	"testing"
)

func TestDictName_golden(t *testing.T) {
	rawName := DictName("test_dict")

	// Get gob encoded dictionary name.
	var gobName bytes.Buffer
	if _, err := rawName.WriteTo(&gobName); err != nil {
		t.Errorf("failed to get encoded name data: %v", err)
	}

	// Decode gob encoded dictionary name.
	decName := ReadDictName(&gobName)

	// Assert be equal.
	if string(rawName) != string(decName) {
		t.Errorf("want %v, got %v", rawName, decName)
	}
}

func TestDictName_bad_input(t *testing.T) {
	t.Run("empty name", func(t *testing.T) {
		name := DictName("")

		// Get gob encoded dictionary name.
		var gobName bytes.Buffer
		if _, err := name.WriteTo(&gobName); err != nil {
			t.Errorf("failed to encode dict name: %v", err)
		}

		// Decode gob encoded dictionary name.
		got := string(ReadDictName(&gobName))

		// Assert be equal to default name.
		want := UndefinedDictName
		if want != got {
			t.Errorf("empty name should return default name. want %v, got %v", want, got)
		}
	})

	t.Run("nil input", func(t *testing.T) {
		// Nil input shuold return default name.
		got := string(ReadDictName(nil))

		// Assert be equal to default name.
		want := UndefinedDictName
		if want != got {
			t.Errorf("nil input should return default name. want %v, got %v", want, got)
		}
	})

	t.Run("bad gob data", func(t *testing.T) {
		// Bad gob data should return default name.
		got := string(ReadDictName(bytes.NewReader([]byte{0x00})))

		// Assert be equal to default name.
		want := UndefinedDictName
		if want != got {
			t.Errorf("bad encoded data should return default name. want %v, got %v", want, got)
		}
	})
}

func TestDictName_WriteTo(t *testing.T) {
	name := DictName("test_dict")

	// Nil writer should return error.
	_, err := name.WriteTo(nil)

	// Assert error.
	if err == nil {
		t.Error("nil writer should return error")
	}
	// Assert error message.
	want := "given writer is nil"
	if want != err.Error() {
		t.Errorf("want %v, got %v", want, err.Error())
	}
}
