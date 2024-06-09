package dict

import (
	"bytes"
	"testing"
)

func TestDictName_golden(t *testing.T) {
	in := Info{Name: "test_dict"}

	// Get gob encoded dictionary name.
	var gobName bytes.Buffer
	if _, err := in.WriteTo(&gobName); err != nil {
		t.Errorf("failed to get encoded name data: %v", err)
	}

	// Decode gob encoded dictionary name.
	out := ReadDictInfo(&gobName)

	// Assert be equal.
	if in.Name != out.Name {
		t.Errorf("want %v, got %v", in, out)
	}
}

func TestDictName_bad_input(t *testing.T) {
	t.Run("empty name", func(t *testing.T) {
		in := Info{Name: ""}

		// Get gob encoded dictionary name.
		var gobName bytes.Buffer
		if _, err := in.WriteTo(&gobName); err != nil {
			t.Errorf("failed to encode dict name: %v", err)
		}

		// Decode gob encoded dictionary name.
		got := ReadDictInfo(&gobName)
		if got.Name != "" {
			t.Errorf("empty name should return empty name. got %v", got)
		}
	})

	t.Run("nil input", func(t *testing.T) {
		// Nil input shuold return default name.
		got := ReadDictInfo(nil)
		if got != nil {
			t.Errorf("nil input should return nil. got %v", got)
		}
	})

	t.Run("bad gob data", func(t *testing.T) {
		// Bad gob data should return default name.
		got := ReadDictInfo(bytes.NewReader([]byte{0x00}))
		if got != nil {
			t.Errorf("bad gob data should return nil. got %v", got)
		}
	})
}

func TestDictName_WriteTo(t *testing.T) {
	in := Info{Name: "test_dict"}

	// Nil writer should return error.
	_, err := in.WriteTo(nil)

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
