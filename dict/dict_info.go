package dict

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io"
)

// Info represents the dictionary info.
type Info struct {
	Name string
	Src  string
}

// ReadDictInfo reads gob encoded dictionary info and returns it.
//
// For backward compatibility, if a dictionary name is not defined or empty, it
// returns UndefinedDictName.
func ReadDictInfo(r io.Reader) *Info {
	if r == nil {
		return nil
	}
	var name string
	dec := gob.NewDecoder(r)
	if err := dec.Decode(&name); err != nil {
		return nil
	}
	var src string
	if err := dec.Decode(&src); err != nil {
		return nil
	}
	return &Info{Name: name, Src: src}
}

// WriteTo implements the io.WriteTo interface.
func (d Info) WriteTo(w io.Writer) (n int64, err error) {
	if w == nil {
		return 0, errors.New("given writer is nil")
	}
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	if err := enc.Encode(d.Name); err != nil {
		return 0, err
	}
	if err := enc.Encode(d.Src); err != nil {
		return 0, err
	}
	return b.WriteTo(w)
}
