package dict

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

func TestNewMultiSizeReaderAt(t *testing.T) {
	b0 := bytes.NewReader([]byte("hello"))
	b1 := bytes.NewReader([]byte("goodbye"))
	b2 := bytes.NewReader([]byte("aloha"))

	m := MultiSizeReaderAt(b0, b1, b2)

	if want, got := b0.Size()+b1.Size()+b2.Size(), m.Size(); want != got {
		t.Errorf("size: want %d, got %d", want, got)
	}
	p := make([]byte, 5)
	k, err := m.ReadAt(p, b0.Size()+int64(len("good")))
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if k != len(p) {
		t.Errorf("expected buffer full read size %d, got %d", len(p), k)
	}
	if want, got := "byeal", string(p); want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestMultiSizeReaderAt_ReadAt(t *testing.T) {
	t.Run("offset=0, buffer size > reader size", func(t *testing.T) {
		m := MultiSizeReaderAt(
			bytes.NewReader([]byte("hello")),
			bytes.NewReader([]byte("goodbye")),
			bytes.NewReader([]byte("aloha")),
		)
		p := make([]byte, m.Size()+1)
		k, err := m.ReadAt(p, 0)
		if want, got := io.EOF, err; !errors.Is(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
		if want, got := m.Size(), int64(k); want != got {
			t.Errorf("want %d, got %d", want, got)
		}
	})
	t.Run("offset=0, buffer size == reader size", func(t *testing.T) {
		m := MultiSizeReaderAt(
			bytes.NewReader([]byte("hello")),
			bytes.NewReader([]byte("goodbye")),
			bytes.NewReader([]byte("aloha")),
		)
		p := make([]byte, m.Size())
		k, err := m.ReadAt(p, 0)
		if want, got := error(nil), err; !errors.Is(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
		if want, got := m.Size(), int64(k); want != got {
			t.Errorf("want %d, got %d", want, got)
		}
	})
	t.Run("offset=0, buffer size < reader size", func(t *testing.T) {
		m := MultiSizeReaderAt(
			bytes.NewReader([]byte("hello")),
			bytes.NewReader([]byte("goodbye")),
			bytes.NewReader([]byte("aloha")),
		)
		p := make([]byte, m.Size()-1)
		k, err := m.ReadAt(p, 0)
		if want, got := error(nil), err; !errors.Is(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
		if want, got := m.Size()-1, int64(k); want != got {
			t.Errorf("want %d, got %d", want, got)
		}
	})
	t.Run("offset=4, buffer size < reader size", func(t *testing.T) {
		m := MultiSizeReaderAt(
			bytes.NewReader([]byte("hello")),
			bytes.NewReader([]byte("goodbye")),
			bytes.NewReader([]byte("aloha")),
		)
		p := make([]byte, len("goodbye")+2)
		k, err := m.ReadAt(p, 4)
		if want, got := error(nil), err; !errors.Is(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
		if want, got := len(p), k; want != got {
			t.Errorf("want %d, got %d", want, got)
		}
		if want, got := "ogoodbyea", string(p); want != got {
			t.Errorf("want %s, got %s", want, got)
		}
	})
}
