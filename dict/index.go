package dict

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"sort"

	"github.com/ikawaha/kagome-dict/dict/trie"
)

// IndexTable represents a dictionary index.
type IndexTable struct {
	Da  trie.DoubleArray
	Dup map[int32]int32

	// dup is Dup as a dense slice indexed by id, built by BuildIndexTable and
	// ReadIndexTable for fast lookups. It is nil when Dup has a key or value
	// that does not fit, and then lookups read Dup. Call RebuildDupCache after
	// changing Dup directly.
	dup []uint8
}

// RebuildDupCache rebuilds the lookup table of Dup. Call it after changing
// Dup directly.
func (idx *IndexTable) RebuildDupCache() {
	idx.dup = nil
	if len(idx.Dup) == 0 {
		return
	}
	var maxKey int32
	for k, v := range idx.Dup {
		if k < 0 || v < 0 || v > math.MaxUint8 {
			return
		}
		maxKey = max(maxKey, k)
	}
	dup := make([]uint8, int(maxKey)+1)
	for k, v := range idx.Dup {
		dup[k] = uint8(v) //nolint:gosec // G115: checked above
	}
	idx.dup = dup
}

// dupOf returns the number of the entries after id that share its keyword.
func (idx *IndexTable) dupOf(id int) int {
	if idx.dup != nil {
		if uint(id) < uint(len(idx.dup)) {
			return int(idx.dup[id])
		}
		return 0
	}
	return int(idx.Dup[int32(id)]) //nolint:gosec //G115: integer overflow conversion int -> int32
}

// BuildIndexTable constructs a index table from keywords.
func BuildIndexTable(sortedKeywords []string) (IndexTable, error) {
	idx := IndexTable{Dup: map[int32]int32{}}
	if !sort.StringsAreSorted(sortedKeywords) {
		return idx, errors.New("unsorted keywords")
	}
	keys := make([]string, 0, len(sortedKeywords))
	ids := make([]int, 0, len(sortedKeywords))
	prev := struct {
		no   int
		word string
	}{}
	for i, key := range sortedKeywords {
		if key == prev.word {
			idx.Dup[int32(prev.no)]++ //nolint:gosec //G115: integer overflow conversion int -> int32
			continue
		}
		prev.no = i
		prev.word = key
		keys = append(keys, key)
		ids = append(ids, i)
	}
	d, err := trie.BuildWithIDs(keys, ids)
	if err != nil {
		return idx, fmt.Errorf("build error, %w", err)
	}
	idx.Da = d
	idx.RebuildDupCache()
	return idx, nil
}

// CommonPrefixSearch finds keywords sharing common prefix in an input
// and returns the ids and it's lengths if found.
func (idx IndexTable) CommonPrefixSearch(input string) (lens []int, ids [][]int) { //nolint:nonamedreturns
	seeds, lens := idx.Da.CommonPrefixSearch(input)
	for _, id := range seeds {
		dup := idx.dupOf(id)
		list := make([]int, 1+dup)
		for i := 0; i < len(list); i++ {
			list[i] = id + i
		}
		ids = append(ids, list)
	}
	return
}

// CommonPrefixSearchCallback finds keywords sharing common prefix in an input
// and callback with id and length.
func (idx IndexTable) CommonPrefixSearchCallback(input string, callback func(id, l int)) {
	idx.Da.CommonPrefixSearchCallback(input, func(x, y int) {
		dup := idx.dupOf(x)
		for i := x; i < x+dup+1; i++ {
			callback(i, y)
		}
	})
}

// Search finds the given keyword and returns the id if found.
func (idx IndexTable) Search(input string) []int {
	id, ok := idx.Da.Find(input)
	if !ok {
		return nil
	}
	dup := idx.dupOf(id)
	list := make([]int, 1+dup)
	for i := 0; i < len(list); i++ {
		list[i] = id + i
	}
	return list
}

// WriteTo implements the io.WriterTo interface.
func (idx IndexTable) WriteTo(w io.Writer) (n int64, err error) { //nolint:nonamedreturns
	n, err = idx.Da.WriteTo(w)
	if err != nil {
		return n, err
	}
	keys := make([]int32, 0, len(idx.Dup))
	for k := range idx.Dup {
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
		v := idx.Dup[k]
		if err := binary.Write(w, binary.LittleEndian, v); err != nil {
			return n, err
		}
		n += int64(binary.Size(v))
	}
	return n, nil
}

// ReadIndexTable loads a index table.
func ReadIndexTable(r io.Reader) (IndexTable, error) {
	idx := IndexTable{}
	d, err := trie.Read(r)
	if err != nil {
		return idx, fmt.Errorf("read index error, %w", err)
	}
	idx.Da = d

	var sz int64
	if err := binary.Read(r, binary.LittleEndian, &sz); err != nil {
		return idx, err
	}
	idx.Dup = make(map[int32]int32, sz)
	for range sz {
		var k int32
		if err := binary.Read(r, binary.LittleEndian, &k); err != nil {
			return idx, err
		}
		var v int32
		if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
			return idx, err
		}
		idx.Dup[k] = v
	}
	idx.RebuildDupCache()

	return idx, nil
}
