package builder

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/ikawaha/kagome-dict/dict"
	"golang.org/x/text/encoding"
)

// MaxInt16 represents the int16 limit value.
const MaxInt16 = 1<<15 - 1

// Config represents the configuration of dictionary builder.
type Config struct {
	paths      []string
	recordInfo *MorphRecordInfo
	unkInfo    *UnkRecordInfo
	enc        encoding.Encoding

	MatrixDefFileName string
	CharDefFileName   string
	UnkDefFileName    string
}

// NewConfig creates a configuration for dictionary builder.
func NewConfig(path string, other []string, enc encoding.Encoding, info *MorphRecordInfo, unk *UnkRecordInfo) *Config {
	paths := append([]string{path}, other...)
	return &Config{
		paths:      paths,
		recordInfo: info,
		unkInfo:    unk,
		enc:        enc,
		// default def file names
		MatrixDefFileName: "matrix.def",
		CharDefFileName:   "char.def",
		UnkDefFileName:    "unk.def",
	}
}

// Build builds a dictionary.
func Build(c *Config) (*dict.Dict, error) {
	if c == nil {
		return nil, fmt.Errorf("empty config")
	}
	if len(c.paths) == 0 {
		return nil, fmt.Errorf("empty path")
	}

	// Morph CSV
	var records Records
	for i, v := range c.paths {
		var enc encoding.Encoding
		if i == 0 {
			enc = c.enc
		}
		rec, err := parseCSVFiles(v, enc, c.recordInfo.ColSize)
		if err != nil {
			return nil, err
		}
		records = append(records, rec...)
	}
	sort.Sort(records)

	ret := dict.Dict{
		Morphs: make([]dict.Morph, 0, len(records)),
		POSTable: dict.POSTable{
			POSs: make([]dict.POS, 0, len(records)),
		},
		ContentsMeta: c.recordInfo.Meta,
		Contents:     make([][]string, 0, len(records)),
	}

	// ConnectionTable
	matrix, err := parseMatrixDefFile(c.paths[0] + "/" + c.MatrixDefFileName)
	if err != nil {
		return nil, err
	}
	ret.Connection.Row = matrix.rowSize
	ret.Connection.Col = matrix.colSize
	ret.Connection.Vec = matrix.vec

	// Words
	var keywords []string
	posMap := make(dict.POSMap)
	for _, rec := range records {
		keywords = append(keywords, rec[c.recordInfo.SurfaceIndex])
		l, err := strconv.Atoi(rec[c.recordInfo.LeftIDIndex])
		if err != nil {
			return nil, err
		}
		if l >= int(matrix.colSize) || l > MaxInt16 {
			return nil, fmt.Errorf("morph left ID %d > %d, record: %v", l, MaxInt16, rec)
		}
		r, err := strconv.Atoi(rec[c.recordInfo.RightIDIndex])
		if err != nil {
			return nil, err
		}
		if r >= int(matrix.rowSize) || r > MaxInt16 {
			return nil, fmt.Errorf("morph right ID %d > %d, record: %v", r, MaxInt16, rec)
		}
		w, err := strconv.Atoi(rec[c.recordInfo.WeightIndex])
		if err != nil {
			return nil, err
		}
		if w > MaxInt16 {
			return nil, fmt.Errorf("morph weight %d > %d, record: %v", r, MaxInt16, rec)
		}

		m := dict.Morph{LeftID: int16(l), RightID: int16(r), Weight: int16(w)}
		ret.Morphs = append(ret.Morphs, m)
		ret.POSTable.POSs = append(ret.POSTable.POSs, posMap.Add(
			rec[c.recordInfo.POSStartIndex:c.recordInfo.OtherContentsStartIndex]),
		)
		ret.Contents = append(ret.Contents, rec[c.recordInfo.OtherContentsStartIndex:])
	}
	ret.POSTable.NameList = posMap.List()

	// Index
	index, err := dict.BuildIndexTable(keywords)
	if err != nil {
		return nil, err
	}
	ret.Index = index

	// CharDef
	def, err := parseCharClassDefFile(c.paths[0] + "/" + c.CharDefFileName)
	if err != nil {
		return nil, err
	}
	ret.CharClass = def.charClass
	ret.CharCategory = def.charCategory
	ret.InvokeList = def.invokeMap
	ret.GroupList = def.groupMap

	// Unk
	unk, err := parseUnkDefFile(filepath.Join(c.paths[0], c.UnkDefFileName), c.enc, c.unkInfo, ret.CharClass)
	if err != nil {
		return nil, fmt.Errorf("unk file parse error, %v", err)
	}
	ret.UnkDict = *unk

	return &ret, err
}
