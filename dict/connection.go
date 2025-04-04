package dict

import (
	"encoding/binary"
	"io"
)

// ConnectionTable represents a connection matrix of morphs.
//
//nolint:recvcheck
type ConnectionTable struct {
	Row, Col int64
	Vec      []int16
}

// At returns the connection cost of matrix[row, col].
func (t *ConnectionTable) At(row, col int) int16 {
	return t.Vec[t.Row*int64(col)+int64(row)] // connection matrix is transposed
}

// WriteTo implements the io.WriterTo interface
func (t ConnectionTable) WriteTo(w io.Writer) (n int64, err error) { //nolint:nonamedreturns
	if err := binary.Write(w, binary.LittleEndian, t.Row); err != nil {
		return n, err
	}
	n += int64(binary.Size(t.Row))
	if err := binary.Write(w, binary.LittleEndian, t.Col); err != nil {
		return n, err
	}
	n += int64(binary.Size(t.Col))
	for i := range t.Vec {
		if err := binary.Write(w, binary.LittleEndian, t.Vec[i]); err != nil {
			return n, err
		}
		n += int64(binary.Size(t.Vec[i]))
	}
	return n, nil
}

// ReadConnectionTable loads ConnectionTable from io.Reader.
func ReadConnectionTable(r io.Reader) (ConnectionTable, error) {
	var ret ConnectionTable
	if err := binary.Read(r, binary.LittleEndian, &ret.Row); err != nil {
		return ret, err
	}
	if err := binary.Read(r, binary.LittleEndian, &ret.Col); err != nil {
		return ret, err
	}
	ret.Vec = make([]int16, ret.Row*ret.Col)
	if err := binary.Read(r, binary.LittleEndian, &ret.Vec); err != nil {
		return ret, err
	}
	return ret, nil
}
