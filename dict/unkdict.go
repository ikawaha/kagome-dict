package dict

import (
	"encoding/binary"
	"fmt"
	"io"
	"sort"
)

// UnkDict represents an unknown word dictionary part.
type UnkDict struct {
	Morphs       Morphs
	Index        map[int32]int32
	IndexDup     map[int32]int32
	ContentsMeta ContentsMeta
	Contents     Contents
}

func writeMapInt32Int32(w io.Writer, m map[int32]int32) (n int64, err error) { //nolint:nonamedreturns
	keys := make([]int32, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	sz := int64(len(keys))
	if err := binary.Write(w, binary.LittleEndian, sz); err != nil {
		return n, err
	}
	n += int64(binary.Size(sz))
	for _, k := range keys {
		if err := binary.Write(w, binary.LittleEndian, k); err != nil {
			return n, err
		}
		n += int64(binary.Size(k))
		v := m[k]
		if err := binary.Write(w, binary.LittleEndian, v); err != nil {
			return n, err
		}
		n += int64(binary.Size(v))
	}
	return n, err
}

// WriteTo implements the io.WriterTo interface.
func (u UnkDict) WriteTo(w io.Writer) (n int64, err error) { //nolint:nonamedreturns
	size, err := writeMapInt32Int32(w, u.Index)
	if err != nil {
		return n, fmt.Errorf("write index error, %w", err)
	}
	n += size
	size, err = writeMapInt32Int32(w, u.IndexDup)
	if err != nil {
		return n, fmt.Errorf("write index dup error, %w", err)
	}
	n += size
	size, err = u.Morphs.WriteTo(w)
	if err != nil {
		return n, fmt.Errorf("write morph error, %w", err)
	}
	n += size

	size, err = u.ContentsMeta.WriteTo(w)
	if err != nil {
		return n, fmt.Errorf("write contents meta, %w", err)
	}
	n += size

	size, err = u.Contents.WriteTo(w)
	if err != nil {
		return n, fmt.Errorf("write contents error, %w", err)
	}
	n += size

	return n, nil
}

func readMapInt32Int32(r io.Reader) (map[int32]int32, error) {
	var sz int64
	if err := binary.Read(r, binary.LittleEndian, &sz); err != nil {
		return nil, err
	}
	m := make(map[int32]int32, sz)
	for range sz {
		var k int32
		if err := binary.Read(r, binary.LittleEndian, &k); err != nil {
			return nil, err
		}
		var v int32
		if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
			return nil, err
		}
		m[k] = v
	}
	return m, nil
}

// ReadUnkDic loads an unknown word dictionary.
func ReadUnkDic(r io.Reader) (UnkDict, error) {
	d := UnkDict{}
	ui, err := readMapInt32Int32(r)
	if err != nil {
		return d, fmt.Errorf("index: %w", err)
	}
	d.Index = ui
	ud, err := readMapInt32Int32(r)
	if err != nil {
		return d, fmt.Errorf("index_dup: %w", err)
	}
	d.IndexDup = ud

	ms, err := ReadMorphs(r)
	if err != nil {
		return d, err
	}
	d.Morphs = ms

	me, err := ReadContentsMeta(r)
	if err != nil {
		return d, err
	}
	d.ContentsMeta = me

	b, err := io.ReadAll(r)
	if err != nil {
		return d, err
	}
	d.Contents = NewContents(b)

	return d, nil
}
