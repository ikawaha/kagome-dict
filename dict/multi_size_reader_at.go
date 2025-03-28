package dict

import (
	"io"
)

// SizeReaderAt is the interface that wraps the Size and ReadAt method.
type SizeReaderAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
	Size() int64
}

type multiSizeReaderAt struct {
	readers []SizeReaderAt
	size  int64
}

// MultiSizeReaderAt returns a SizeReaderAt that is the logical concatenation of the provided input readers.
func MultiSizeReaderAt(rs ...SizeReaderAt) SizeReaderAt {
	var size int64
	for _, v := range rs {
		size += v.Size()
	}
	return &multiSizeReaderAt{
		readers: rs,
		size:    size,
	}
}

// Size returns the size of the reader.
func (m multiSizeReaderAt) Size() int64{
	return m.size
}

// ReadAt reads len(p) bytes into p starting at offset off in the underlying input source.
// It returns the number of bytes read (0 <= n <= len(p)) and any error encountered.
func (m *multiSizeReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	full := len(p)
	if off >= m.size {
		return 0, io.EOF
	}
	for i:=0; i < len(m.readers) && len(p) !=0; i++{
		if int(m.readers[i].Size()) < int(off){
			off -= m.readers[i].Size()
			continue
		}
		k, err :=m.readers[i].ReadAt(p, off)
		n +=  k
		if err != nil {
			if err.Error() != "EOF" {
				return n, err
			}
			if i == len(m.readers)-1 && err.Error() == "EOF" {
				return n, io.EOF
			}
		}
		p = p[k:]
		off = 0
	}
	if n != full {
		return n, io.EOF
	}
	return n, nil
}