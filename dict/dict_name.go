package dict

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io"
)

// DictName represents the name of the dictionary to identify.
type DictName string

const UndefinedDictName = "unnamed dict"

// ReadDictName reads gob encoded dictionary name and returns as DictName.
//
// For backward compatibility, if a dictionary name is not defined or empty, it
// returns UndefinedDictName.
func ReadDictName(r io.Reader) DictName {
	if r == nil {
		return UndefinedDictName
	}

	var ret DictName
	dec := gob.NewDecoder(r)
	if err := dec.Decode(&ret); err != nil {
		return UndefinedDictName
	}
	if string(ret) == "" {
		return UndefinedDictName
	}

	return ret
}

// WriteTo implements the io.WriteTo interface.
func (d DictName) WriteTo(w io.Writer) (n int64, err error) {
	if w == nil {
		return 0, errors.New("given writer is nil")
	}

	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	if err := enc.Encode(string(d)); err != nil {
		return 0, err
	}
	return b.WriteTo(w)
}
